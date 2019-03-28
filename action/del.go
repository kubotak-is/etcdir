package action

import (
	"fmt"
	"github.com/kubotak-is/etcdir/client"
	"strings"
)

func Del(n string, k string) error {
	cli, err := client.New(strings.Split(n, ","))
	defer cli.Close()
	if err != nil {
		panic(err)
	}

	err = client.Delete(cli, k)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("complete!")

	return nil
}
