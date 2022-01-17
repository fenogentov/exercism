package ledger

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

type Msg struct {
	i   int
	str string
	err error
}

var locales = map[string]struct {
	header       string
	currency     map[string]string
	formatDate   string
	decimalPoint string
	thousandSep  string
	negative     func(string) string
}{
	"nl-NL": {
		header:       "Datum      | Omschrijving              | Verandering\n",
		currency:     map[string]string{"EUR": "€ ", "USD": "$ "},
		formatDate:   "",
		decimalPoint: ",",
		thousandSep:  ".",
		negative: func(n string) string {
			return n + "-"
		},
	},
	"en-US": {
		header:       "Date       | Description               | Change\n",
		currency:     map[string]string{"EUR": "€", "USD": "$"},
		formatDate:   "",
		decimalPoint: ".",
		thousandSep:  ",",
		negative: func(n string) string {
			return "(" + n + ")"
		},
	},
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	if currency != "USD" && currency != "EUR" {
		return "", errors.New("invalid currency")
	}
	if locale != "nl-NL" && locale != "en-US" {
		return "", errors.New("invalid locale")
	}

	entriesCopy := append([]Entry{}, entries...)

	if len(entries) == 0 {
		if _, err := FormatLedger(currency, "en-US", []Entry{{Date: "2014-01-01", Description: "", Change: 0}}); err != nil {
			return "", err
		}
	}
	m1 := map[bool]int{true: 0, false: 1}
	m2 := map[bool]int{true: -1, false: 1}
	es := entriesCopy

	for len(es) > 1 {
		first, rest := es[0], es[1:]
		success := false
		for !success {
			success = true
			for i, e := range rest {
				if (m1[e.Date == first.Date]*m2[e.Date < first.Date]*4 +
					m1[e.Description == first.Description]*m2[e.Description < first.Description]*2 +
					m1[e.Change == first.Change]*m2[e.Change < first.Change]*1) < 0 {
					es[0], es[i+1] = es[i+1], es[0]
					success = false
				}
			}
		}
		es = es[1:]
	}

	// Parallelism, always a great idea
	co := make(chan Msg)
	for i, et := range entriesCopy {
		go func(i int, entry Entry) {
			if len(entry.Date) != 10 {
				co <- Msg{err: errors.New("")}
			}

			d1, d2, d3, d4, d5 := entry.Date[0:4], entry.Date[4], entry.Date[5:7], entry.Date[7], entry.Date[8:10]
			if d2 != '-' || d4 != '-' {
				co <- Msg{err: errors.New("")}
			}
			de := entry.Description
			if len(de) > 25 {
				de = de[:22] + "..."
			}
			var d string
			if locale == "nl-NL" {
				d = d5 + "-" + d3 + "-" + d1
			} else {
				d = d3 + "/" + d5 + "/" + d1
			}

			cents := entry.Change

			co <- Msg{i: i, str: formatRow(d, de, money(locale, currency, cents))}

		}(i, et)
	}
	ss := make([]string, len(entriesCopy))
	for range entriesCopy {
		v := <-co
		if v.err != nil {
			return "", v.err
		}
		ss[v.i] = v.str
	}
	return locales[locale].header + strings.Join(ss, ""), nil
}
func formatRow(d, de, a string) string {
	return fmt.Sprintf("%-10s | %-25s | %13s\n", d, de, a)
}

func money(locale, currency string, cents int) string {
	negative := false
	if cents < 0 {
		cents = -cents
		negative = true
	}

	centsStr := fmt.Sprintf("%03d", cents)

	rest := centsStr[:len(centsStr)-2]

	var parts []string
	for len(rest) > 3 {
		parts = append(parts, rest[len(rest)-3:])
		rest = rest[:len(rest)-3]
	}

	if len(rest) > 0 {
		parts = append(parts, rest)
	}

	var a string
	a += locales[locale].currency[currency]
	for i := len(parts) - 1; i >= 0; i-- {
		a += parts[i] + locales[locale].thousandSep
	}
	a = a[:len(a)-1] + locales[locale].decimalPoint + centsStr[len(centsStr)-2:]

	if negative {
		a = locales[locale].negative(a)
	} else {
		a += " "
	}

	return a
}

func oldFormatLedger(currency string, locale string, entries []Entry) (string, error) {
	var entriesCopy []Entry
	for _, e := range entries {
		entriesCopy = append(entriesCopy, e)
	}
	if len(entries) == 0 {
		if _, err := FormatLedger(currency, "en-US", []Entry{{Date: "2014-01-01", Description: "", Change: 0}}); err != nil {
			return "", err
		}
	}
	m1 := map[bool]int{true: 0, false: 1}
	m2 := map[bool]int{true: -1, false: 1}
	es := entriesCopy
	for len(es) > 1 {
		first, rest := es[0], es[1:]
		success := false
		for !success {
			success = true
			for i, e := range rest {
				if (m1[e.Date == first.Date]*m2[e.Date < first.Date]*4 +
					m1[e.Description == first.Description]*m2[e.Description < first.Description]*2 +
					m1[e.Change == first.Change]*m2[e.Change < first.Change]*1) < 0 {
					es[0], es[i+1] = es[i+1], es[0]
					success = false
				}
			}
		}
		es = es[1:]
	}

	var s string
	if locale == "nl-NL" {
		s = "Datum" +
			strings.Repeat(" ", 10-len("Datum")) +
			" | " +
			"Omschrijving" +
			strings.Repeat(" ", 25-len("Omschrijving")) +
			" | " + "Verandering" + "\n"
	} else if locale == "en-US" {
		s = "Date" +
			strings.Repeat(" ", 10-len("Date")) +
			" | " +
			"Description" +
			strings.Repeat(" ", 25-len("Description")) +
			" | " + "Change" + "\n"
	} else {
		return "", errors.New("")
	}
	// Parallelism, always a great idea
	co := make(chan struct {
		i int
		s string
		e error
	})
	for i, et := range entriesCopy {
		go func(i int, entry Entry) {
			if len(entry.Date) != 10 {
				co <- struct {
					i int
					s string
					e error
				}{e: errors.New("")}
			}
			d1, d2, d3, d4, d5 := entry.Date[0:4], entry.Date[4], entry.Date[5:7], entry.Date[7], entry.Date[8:10]
			if d2 != '-' {
				co <- struct {
					i int
					s string
					e error
				}{e: errors.New("")}
			}
			if d4 != '-' {
				co <- struct {
					i int
					s string
					e error
				}{e: errors.New("")}
			}
			de := entry.Description
			if len(de) > 25 {
				de = de[:22] + "..."
			} else {
				de = de + strings.Repeat(" ", 25-len(de))
			}
			var d string
			if locale == "nl-NL" {
				d = d5 + "-" + d3 + "-" + d1
			} else if locale == "en-US" {
				d = d3 + "/" + d5 + "/" + d1
			}
			negative := false
			cents := entry.Change
			if cents < 0 {
				cents = cents * -1
				negative = true
			}
			var a string
			if locale == "nl-NL" {
				if currency == "EUR" {
					a += "€"
				} else if currency == "USD" {
					a += "$"
				} else {
					co <- struct {
						i int
						s string
						e error
					}{e: errors.New("")}
				}
				a += " "
				centsStr := strconv.Itoa(cents)
				switch len(centsStr) {
				case 1:
					centsStr = "00" + centsStr
				case 2:
					centsStr = "0" + centsStr
				}
				rest := centsStr[:len(centsStr)-2]
				var parts []string
				for len(rest) > 3 {
					parts = append(parts, rest[len(rest)-3:])
					rest = rest[:len(rest)-3]
				}
				if len(rest) > 0 {
					parts = append(parts, rest)
				}
				for i := len(parts) - 1; i >= 0; i-- {
					a += parts[i] + "."
				}
				a = a[:len(a)-1]
				a += ","
				a += centsStr[len(centsStr)-2:]
				if negative {
					a += "-"
				} else {
					a += " "
				}
			} else if locale == "en-US" {
				if negative {
					a += "("
				}
				if currency == "EUR" {
					a += "€"
				} else if currency == "USD" {
					a += "$"
				} else {
					co <- struct {
						i int
						s string
						e error
					}{e: errors.New("")}
				}
				centsStr := strconv.Itoa(cents)
				switch len(centsStr) {
				case 1:
					centsStr = "00" + centsStr
				case 2:
					centsStr = "0" + centsStr
				}
				rest := centsStr[:len(centsStr)-2]
				var parts []string
				for len(rest) > 3 {
					parts = append(parts, rest[len(rest)-3:])
					rest = rest[:len(rest)-3]
				}
				if len(rest) > 0 {
					parts = append(parts, rest)
				}
				for i := len(parts) - 1; i >= 0; i-- {
					a += parts[i] + ","
				}
				a = a[:len(a)-1]
				a += "."
				a += centsStr[len(centsStr)-2:]
				if negative {
					a += ")"
				} else {
					a += " "
				}
			} else {
				co <- struct {
					i int
					s string
					e error
				}{e: errors.New("")}
			}
			var al int
			for range a {
				al++
			}
			co <- struct {
				i int
				s string
				e error
			}{i: i, s: d + strings.Repeat(" ", 10-len(d)) + " | " + de + " | " +
				strings.Repeat(" ", 13-al) + a + "\n"}
		}(i, et)
	}
	ss := make([]string, len(entriesCopy))
	for range entriesCopy {
		v := <-co
		if v.e != nil {
			return "", v.e
		}
		ss[v.i] = v.s
	}
	for i := 0; i < len(entriesCopy); i++ {
		s += ss[i]
	}
	return s, nil
}
