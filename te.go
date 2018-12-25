package main

import (
	"encoding/xml"
	"fmt"
	"github.com/dougEfresh/gtoggl"
	"github.com/dougEfresh/gtoggl-api/gttimentry"
	"github.com/fatih/color"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Config struct {
	TogglToken string
}

var config Config
var app *cli.App
var pattern = "02.01.2006"

func init() {
	xmlFile, err := ioutil.ReadFile("config.xml")
	if err != nil {
		log.Fatal(err)
	}
	err = xml.Unmarshal(xmlFile, &config)
	if err != nil {
		panic(err)
	}
	app = cli.NewApp()
	app.Name = "Toggl Earnings"
	app.Usage = "Простая, консольная, утилитка для быстрого просмотра отчетности с сервиса https://toggl.com"
	app.UsageText = "te -start 06.08.2018 -stop 01.01.2019 -r 1250 -d \n \tВсе параметры не обязательные"
	app.Version = "1.0.0"
	app.Author = "Соколов Илья @red_byte"
	app.Email = "developer.sokolov@gmail.com"
	app.Copyright = "[CopyLeft]"

	timeNow := time.Now()
	timeDayAgo := timeNow.AddDate(0, -1, 0)

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "start, st",
			Value: timeDayAgo.Format(pattern),
			Usage: "Начало работы в формате dd.mm.yyyy",
		},
		cli.StringFlag{
			Name:  "stop, sp",
			Value: timeNow.Format(pattern),
			Usage: "Конец работы в формате dd.mm.yyyy",
		},
		cli.Float64Flag{
			Name:  "rate, r",
			Value: 1000.00,
			Usage: "Денежная ставка в час (в рублях)",
		},
		cli.BoolFlag{
			Name:  "details, d",
			Usage: "Детализация по каждому дню",
		},
	}

}

func main() {
	toggl, err := gtoggl.NewClient(config.TogglToken)
	if err != nil {
		panic(err)
	}

	app.Action = func(c *cli.Context) error {
		strStart := c.GlobalString("start")
		strStop := c.GlobalString("stop")
		rate := c.GlobalFloat64("rate")
		details := c.GlobalBool("details")

		start, err := time.Parse(pattern, strStart)
		stop, err := time.Parse(pattern, strStop)

		if err != nil {
			panic(err)
		}
		timeList, err := toggl.TimeentryClient.GetRange(start, stop)
		if err != nil {
			panic(err)
		}
		showInformation(timeList, rate, details)
		return nil
	}
	_ = app.Run(os.Args)
}

func showInformation(timeList gttimeentry.TimeEntries, rate float64, details bool) {
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
	fmt.Println("Начало работы:", entry.Start.Format("02.01.2006 15:04:05"))
	fmt.Println("Конец работы:", entry.Stop.Format("02.01.2006 15:04:05"))
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - -")
	fmt.Println()
}
