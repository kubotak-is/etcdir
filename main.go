package main

import (
	"github.com/kubotak-is/etcdir/action"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "etcdir"
	app.Usage = "Reflect directory structure to etcd"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "nodes, n",
			Usage: "etcd host name",
		},
		cli.StringFlag{
			Name: "dir, d",
			Value: "./",
			Usage: "root directory path",
		},
	}

	app.Action =  action.Run
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

