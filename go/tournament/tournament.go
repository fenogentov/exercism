// Package tournament is a solution to lerning #14 (exercism.io)
package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

const (
	tableHeader = "Team                           | MP |  W |  D |  L |  P\n"
	tableRow    = "%-30s | %2d | %2d | %2d | %2d | %2d\n"
)

type results struct {
	matches, wins, losses, draws, points int
}

// Tally receives tournament results from io.Reader and creates a table of results in io.Writer
// Results sorted by points and alphabetically
func Tally(r io.Reader, w io.Writer) error {
	table := map[string]*results{}
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || line[0] == '#' {
			continue
		}
		splitString := strings.Split(line, ";")
		if len(splitString) != 3 {
			return errors.New("err")
		}
		if _, ok := table[splitString[0]]; !ok {
			table[splitString[0]] = &results{}
		}
		if _, ok := table[splitString[1]]; !ok {
			table[splitString[1]] = &results{}
		}

		switch splitString[2] {
		case "win":
			win(table, splitString)
		case "draw":
			draw(table, splitString)
		case "loss":
			loss(table, splitString)
		default:
			return errors.New("err")
		}
	}

	_, err := w.Write([]byte(printTable(table)))
	if err != nil {
		return err
	}

	return nil
}

func win(table map[string]*results, line []string) {
	table[line[0]].matches++
	table[line[0]].wins++
	table[line[0]].points += 3
	table[line[1]].matches++
	table[line[1]].losses++
}

func draw(table map[string]*results, line []string) {
	table[line[0]].matches++
	table[line[0]].draws++
	table[line[0]].points++
	table[line[1]].matches++
	table[line[1]].draws++
	table[line[1]].points++
}
func loss(table map[string]*results, line []string) {
	table[line[0]].matches++
	table[line[0]].losses++
	table[line[1]].matches++
	table[line[1]].wins++
	table[line[1]].points += 3
}

func printTable(table map[string]*results) (str string) {
	str = tableHeader

	keys, listNames := getKeyAndListNames(table)

	for _, key := range keys {
		names := listNames[key]
		for _, name := range names {
			str += fmt.Sprintf(tableRow, name, table[name].matches, table[name].wins, table[name].draws, table[name].losses, table[name].points)
		}
	}
	return str
}

func getKeyAndListNames(table map[string]*results) ([]int, map[int][]string) {
	listNames := map[int][]string{}

	k := map[int]bool{}
	for name, strct := range table {
		k[strct.points] = true
		listNames[strct.points] = append(listNames[strct.points], name)
	}

	keys := []int{}
	for key := range k {
		keys = append(keys, key)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	for i, names := range listNames {
		sort.Strings(names)
		listNames[i] = names
	}
	return keys, listNames
}
