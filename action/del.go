package action

import (
	"strings"
	"fmt"
	"github.com/kubotak-is/etcdir/client"
	"github.com/urfave/cli"
)

func Del(c *cli.Context) error {
	nodes := c.String("nodes")
	cli, err := client.New(strings.Split(nodes, ","))
	defer cli.Close()
	if err != nil {
		panic(err)
	}

	k := c.String("key")
	if k == "" {
		k = "/"
	}

	err = client.Delete(cli, k)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("complete!")

	return nil
}
