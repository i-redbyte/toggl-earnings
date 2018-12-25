package main

import (
	"encoding/xml"
	"fmt"
	"github.com/dougEfresh/gtoggl"
	"github.com/dougEfresh/gtoggl-api/gttimentry"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"time"
)

type Config struct {
	TogglToken string
}

func main() {
	var config Config
	xmlFile, err := ioutil.ReadFile("config.xml")
	if err != nil {
		log.Fatal(err)
	}
	err = xml.Unmarshal(xmlFile, &config)
	if err != nil {
		panic(err)
	}
	toggl, err := gtoggl.NewClient(config.TogglToken)
	if err != nil {
		panic(err)
	}

	pattern := "02.01.2006"

	strStart := "06.12.2018"
	strStop := "21.12.2018"

	start, err := time.Parse(pattern, strStart)
	stop, err := time.Parse(pattern, strStop)

	if err != nil {
		panic(err)
	}
	timeList, err := toggl.TimeentryClient.GetRange(start, stop)
	if err != nil {
		panic(err)
	}

	showInformation(timeList)
}

func showInformation(timeList gttimeentry.TimeEntries) {
	var duration time.Duration
	for _, v := range timeList {
		showTimeEntryInfo(v)
		duration += time.Duration(v.Duration) * time.Second
	}
	fmt.Println("Сумма:", duration.Hours()*1000, "руб.")

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
	fmt.Println("Начало работы:", entry.Start.Format("02.01.2006 15:04:05"))
	fmt.Println("Конец работы:", entry.Stop.Format("02.01.2006 15:04:05"))
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - -")
	fmt.Println()
}
