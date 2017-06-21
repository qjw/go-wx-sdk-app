package main

import (
	"os"
	"github.com/ianschenck/envflag"
	"github.com/codegangsta/cli"
	"time"
	"math/rand"
	"github.com/qjw/go-wx-sdk-app/const_var"
)

func main() {
	rand.Seed(time.Now().Unix())
	envflag.Parse()
	app := cli.NewApp()
	app.Name = "go-wx-sdk-app"
	app.Version = const_var.Version
	app.Usage = "微信SDK命令行工具"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "d,debug",
			Usage:  "开启调试模式",
		},
	}
	app.Commands = []cli.Command{
		serverCmd,
	}
	app.Run(os.Args)
}