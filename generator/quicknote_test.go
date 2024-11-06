package generator

import (
	"testing"
	"time"
)

func TestQuicknoteName(t *testing.T) {
	tests := []struct {
		year, day, hour, minute, second int
		month                           time.Month
		expected                        string
	}{
		{
			year:     2023,
			day:      1,
			hour:     9,
			minute:   42,
			second:   43,
			month:    time.October,
			expected: "20231001094243",
		},
		{
			year:     2024,
			day:      12,
			hour:     15,
			minute:   56,
			second:   10,
			month:    time.May,
			expected: "20240512155610",
		},
	}

	for _, test := range tests {
		time := time.Date(
			test.year,
			test.month,
			test.day,
			test.hour,
			test.minute,
			test.second,
			0,
			time.UTC,
		)
		result := GenerateQuickNoteTitle(time)
		if result != test.expected {
			t.Errorf("Build(%v) = \"%s\"; want \"%s\"", time, result, test.expected)
		}
	}
}
