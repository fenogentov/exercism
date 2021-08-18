// Package twelve is a solution to lerning #15 (exercism.io)
package twelve

import "fmt"

var gifts = []string{
	"a Partridge in a Pear Tree.",
	"two Turtle Doves, and ",
	"three French Hens, ",
	"four Calling Birds, ",
	"five Gold Rings, ",
	"six Geese-a-Laying, ",
	"seven Swans-a-Swimming, ",
	"eight Maids-a-Milking, ",
	"nine Ladies Dancing, ",
	"ten Lords-a-Leaping, ",
	"eleven Pipers Piping, ",
	"twelve Drummers Drumming, ",
}

var nor = []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}

// Song composes a famous song about gifts in 12 days
func Song() string {
	song := ""

	for day := 1; day <= 12; day++ {
		line := Verse(day)
		song += line
		if day < 12 {
			song += fmt.Sprintln()
		}
	}

	return song
}

// Verse composes the line of the song
func Verse(day int) string {
	day--
	presents := ""
	for i := day; i >= 0; i-- {
		presents += gifts[i]
	}
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s", nor[day], presents)
}
