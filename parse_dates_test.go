package datesfinder

import (
	"io"
	"os"
	"testing"
	"time"
)

func TestParseDatesInDocument(t *testing.T) {
	cases := []struct {
		name     string
		text     string
		expected []time.Time
	}{
		{
			name: "Italian",
			text: getTestText("1.txt"),
			expected: []time.Time{
				time.Date(2023, 5, 11, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "French",
			text: getTestText("2.txt"),
			expected: []time.Time{
				time.Date(2023, 7, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, 6, 30, 0, 0, 0, 0, time.UTC),
				time.Date(2023, 7, 2, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "French2",
			text: getTestText("3.txt"),
			expected: []time.Time{
				time.Date(2023, 4, 13, 0, 0, 0, 0, time.UTC),
				time.Date(2023, 5, 31, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			dates, errors := FindDates(tc.text)
			if len(errors) > 0 {
				t.Fatalf("expected no errors, got %v", errors)
			}
			if len(dates) != len(tc.expected) {
				t.Fatalf("expected %d dates, got %d: %v", len(tc.expected), len(dates), dates)
			}
			for i, date := range dates {
				if date != tc.expected[i] {
					t.Errorf("expected date %d to be %s, got %s", i, tc.expected[i], date)
				}
			}
		})
	}
}

func getTestText(s string) string {
	f, err := os.Open("./resources/test/" + s)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	text, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return string(text)
}
