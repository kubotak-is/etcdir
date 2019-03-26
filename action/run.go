package action

import (
	"strings"
	"fmt"
	"github.com/kubotak-is/etcdir/client"
	"github.com/urfave/cli"
)

func Run(c *cli.Context) error {
	nodes := c.String("nodes")
	cli, err := client.New(strings.Split(nodes, ","))
	defer cli.Close()
	if err != nil {
		panic(err)
	}

	dir := c.String("dir")
	files := dirwalk(dir)

	ch := make(chan string)

	for _, p := range files {

		go func(path string) {
			v, err := readFile(path)
			if err != nil {
				fmt.Println("Error: file open")
			} else {
				err := client.Put(cli, strings.Replace(path, dir, "", 1), string(v))
				if err != nil {
					fmt.Println(err)
				}
			}
			ch <- "End function"
		}(p)
		<- ch
	}

	fmt.Println("complete!")

	return nil
}
