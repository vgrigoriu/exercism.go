// Package twelve solves the Twelve Days problem from Exercism.
package twelve

import "fmt"

// Song returns the words of the Twelve Days of Christmas song.
func Song() string {
	result := ""
	for day := 1; day < len(days); day++ {
		result += Verse(day) + "\n"
	}
	return result
}

// Verse returns one verse from the Twelve Days of Christmas song.
func Verse(day int) string {
	if day < 1 || len(days) <= day {
		panic(fmt.Sprintf("day must be between 1 and %d", len(days)-1))
	}
	return "On the " + days[day].ordinal + " day of Christmas my true love gave to me: " + gifts(day) + "."
}

var memoGifts = make([]string, len(days))

func gifts(day int) string {
	if memoGifts[day] != "" {
		return memoGifts[day]
	}

	result := ""
	switch day {
	case 1:
		result = days[day].gift
	case 2:
		result = days[day].gift + ", and " + gifts(day-1)
	default:
		result = days[day].gift + ", " + gifts(day-1)
	}
	memoGifts[day] = result
	return result
}

var days = []struct {
	ordinal string
	gift    string
}{
	{},
	{"first", "a Partridge in a Pear Tree"},
	{"second", "two Turtle Doves"},
	{"third", "three French Hens"},
	{"fourth", "four Calling Birds"},
	{"fifth", "five Gold Rings"},
	{"sixth", "six Geese-a-Laying"},
	{"seventh", "seven Swans-a-Swimming"},
	{"eighth", "eight Maids-a-Milking"},
	{"ninth", "nine Ladies Dancing"},
	{"tenth", "ten Lords-a-Leaping"},
	{"eleventh", "eleven Pipers Piping"},
	{"twelfth", "twelve Drummers Drumming"},
}
