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
	app.Version = "0.2.0"

	app.Commands = []cli.Command{
		{
			Name: "put",
			Aliases: []string{"p"},
			Usage: "Update by difference",
			Action: func (c *cli.Context) error {
				nodes := c.String("nodes")
				dir := c.String("dir")
				return action.Run(nodes, dir)
			},
			Flags: []cli.Flag {
				cli.StringFlag{
					Name: "nodes, n",
					Usage: "etcd host name",
				},
				cli.StringFlag{
					Name: "dir, d",
					Value: "./",
					Usage: "root directory path",
				},
			},
		},
		{
			Name: "delete",
			Aliases: []string{"d"},
			Usage: "Delete all value",
			Action: func (c *cli.Context) error {
				nodes := c.String("nodes")
				key := c.String("key")
				return action.Del(nodes, key)
			},
			Flags: []cli.Flag {
				cli.StringFlag{
					Name: "nodes, n",
					Usage: "etcd host name",
				},
				cli.StringFlag{
					Name: "key, k",
					Value: "/",
					Usage: "delete key",
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

