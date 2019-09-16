package cmd

import (
	"github.com/dougEfresh/gtoggl"
	"github.com/urfave/cli"
	"os"
	"time"
	"timeCalc/info"
)

const datePattern = "02.01.2006"

var app *cli.App

func InitCLI() {
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
			Value: timeDayAgo.Format(datePattern),
			Usage: "Начало работы в формате dd.mm.yyyy",
		},
		cli.StringFlag{
			Name:  "stop, sp",
			Value: timeNow.Format(datePattern),
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

func SetupActionsAndRun(toggl *gtoggl.TogglClient) {
	app.Action = func(c *cli.Context) error {
		strStart := c.GlobalString("start")
		strStop := c.GlobalString("stop")
		rate := c.GlobalFloat64("rate")
		details := c.GlobalBool("details")

		start, err := time.Parse(datePattern, strStart)
		stop, err := time.Parse(datePattern, strStop)

		if err != nil {
			panic(err)
		}
		timeList, err := toggl.TimeentryClient.GetRange(start, stop)
		if err != nil {
			panic(err)
		}
		info.ShowInformation(timeList, rate, details)
		return nil
	}
	_ = app.Run(os.Args)
}
