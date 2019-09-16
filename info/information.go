package info

import (
	"fmt"
	gttimeentry "github.com/dougEfresh/gtoggl-api/gttimentry"
	"github.com/fatih/color"
	"time"
)

const (
	DateTimePattern = "02.01.2006 15:04:05"
)

func ShowInformation(timeList gttimeentry.TimeEntries, rate float64, details bool) {
	var duration time.Duration
	for _, v := range timeList {
		if details {
			showTimeEntryInfo(v)
		}
		duration += time.Duration(v.Duration) * time.Second
	}
	fmt.Printf("Сумма: %.2f руб.\n", duration.Hours()*rate)
	fmt.Println("Ставка в час:", rate, "руб./час")
	fmt.Println("Общее время:", durationToTime(duration))

}

func durationToTime(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

func showTimeEntryInfo(entry gttimeentry.TimeEntry) {
	color.Green("= = = = = = = = = = I N F O = = = = = = = = = =")

	if entry.Project != nil {
		fmt.Println("Проект:", entry.Project)
	}
	if entry.Tags != nil {
		fmt.Println("Тэги:", entry.Tags)
	}
	fmt.Println("Описание:", entry.Description)
	fmt.Println("Начало работы:", entry.Start.Format(DateTimePattern))
	fmt.Println("Конец работы:", entry.Stop.Format(DateTimePattern))
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - -")
	fmt.Println()
}
