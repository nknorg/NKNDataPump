package main

import (
	"github.com/nknorg/NKNDataPump/cli/commands"
	"github.com/nknorg/NKNDataPump/common"
	"github.com/nknorg/NKNDataPump/config"
	"github.com/nknorg/NKNDataPump/server"
	"github.com/nknorg/NKNDataPump/storage"
	"github.com/nknorg/NKNDataPump/task/chainDataPump"
	"github.com/nknorg/NKNDataPump/task/dataInspection"
	"github.com/urfave/cli"
	"os"
	"time"
	"github.com/nknorg/NKNDataPump/network/rpcRequest"
)

func getCliApp() (app *cli.App) {
	app = cli.NewApp()
	app.Name = "NKNDataPump"
	app.Version = "0.0.1"
	app.HelpName = "NKNDataPump"
	app.Usage = "DNA blockchain gateway"
	app.UsageText = "chainGateway [options] [args]"
	app.HideHelp = false
	app.HideVersion = false

	return
}

func main() {
	//get app
	gateway := getCliApp()

	//set some flags
	commands.SetFlags(gateway)
	commands.SetAction(gateway)

	//run
	err := gateway.Run(os.Args)
	if nil != err {
		os.Exit(-1)
	}

	//init logs
	common.InitLog(config.PumpConfig.Logfile, config.PumpConfig.LogLevel)

	//init db
	storage.Init()

	//set up api
	rpcRequest.Api.Build()

	//start data pump
	chainDataPump.Start()

	//start data inspection
	dataInspection.Start()

	//start api & web server
	server.Start()

	for {
		time.Sleep(time.Second)
	}
}
