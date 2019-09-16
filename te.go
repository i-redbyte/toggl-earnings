package main

import (
	"encoding/xml"
	"github.com/dougEfresh/gtoggl"
	"io/ioutil"
	"log"
	"timeCalc/cmd"
)

type Config struct {
	TogglToken string
}

var config Config

func init() {
	xmlFile, err := ioutil.ReadFile("config.xml")
	if err != nil {
		log.Fatal(err)
	}
	err = xml.Unmarshal(xmlFile, &config)
	if err != nil {
		panic(err)
	}
	cmd.InitCLI()
}

func main() {
	toggl, err := gtoggl.NewClient(config.TogglToken)
	if err != nil {
		panic(err)
	}
	cmd.SetupActionsAndRun(toggl)
}
