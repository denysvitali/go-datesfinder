package datesfinder

import (
	"github.com/goodsign/monday"
	"regexp"
	"strings"
	"time"
)

type dateFormat struct {
	expr   *regexp.Regexp
	format string
	locale monday.Locale
}

var dateFormats = []dateFormat{
	{
		expr:   regexp.MustCompile(`(\d{1,2}\.\s*(?:Januar|Februar|März|April|Mai|Juni|Juli|August|September|Oktober|November|Dezember)\s*\d{4})`),
		format: "02. January 2006",
		locale: monday.LocaleDeDE,
	},
	{
		expr:   regexp.MustCompile(`(?:Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec)\w*\b\s*\d{1,2},?\s*\d{4}|(?:\d{1,2}\s+(?:Jan(?:uary)?|Feb(?:ruary)?|Mar(?:ch)?|Apr(?:il)?|May|Jun(?:e)?|Jul(?:y)?|Aug(?:ust)?|Sep(?:tember)?|Oct(?:ober)?|Nov(?:ember)?|Dec(?:ember)?)\w*\s*\d{4})\b`),
		format: "January 2, 2006",
		locale: monday.LocaleEnUS,
	},
	{
		expr:   regexp.MustCompile(`(\d{1,2}\s*(?:gennaio|febbraio|marzo|aprile|maggio|giugno|luglio|agosto|settembre|ottobre|novembre|dicembre)\s*\d{4})`),
		format: "2 January 2006",
		locale: monday.LocaleItIT,
	},
	{
		expr:   regexp.MustCompile(`\d{4}-\d{2}-\d{2}`),
		format: "2006-01-02",
		locale: monday.LocaleEnUS,
	},
	{
		expr:   regexp.MustCompile(`\d{2}/\d{2}/\d{4}`),
		format: "02/01/2006",
		locale: monday.LocaleEnUS,
	},
	{
		expr:   regexp.MustCompile(`\d{2}\.\d{2}\.\d{4}`),
		format: "02.01.2006",
		locale: monday.LocaleEnUS,
	},
	{
		expr:   regexp.MustCompile(`(\d{1,2}\s*(?:janvier|février|mars|avril|mai|juin|juillet|août|septembre|octobre|novembre|décembre)\s*\d{4})`),
		format: "2 January 2006",
		locale: monday.LocaleFrFR,
	},
}

type dateWithFormat struct {
	format dateFormat
	date   string
}

func FindDates(text string) ([]time.Time, []error) {
	// We run the regular expressions on each line, so that they're sorted by line number.
	var dates []dateWithFormat
	for _, line := range strings.Split(text, "\n") {
		for _, f := range dateFormats {
			if match := f.expr.FindString(line); match != "" {
				dates = append(dates, dateWithFormat{format: f, date: match})
			}
		}
	}

	var finalDates []time.Time
	var errors []error
	// Now we parse the dates as time.Time and add them to the final array
	for _, d := range dates {
		d, err := monday.Parse(d.format.format, d.date, d.format.locale)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		finalDates = append(finalDates, d)
	}

	return finalDates, errors
}
