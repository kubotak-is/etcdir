package action

import (
	"strings"
	"fmt"
	"github.com/kubotak-is/etcdir/client"
)

func Run(n string, d string) error {
	cli, err := client.New(strings.Split(n, ","))
	defer cli.Close()
	if err != nil {
		panic(err)
	}

	files := dirwalk(d)

	ch := make(chan string)

	for _, p := range files {

		go func(path string) {
			v, err := readFile(path)
			if err != nil {
				fmt.Println("Error: file open")
			} else {
				p := strings.Replace(path, d, "", 1)
				diff, err := client.Diff(cli, p, string(v))
				if diff {
					// Update only when there is a difference
					err = client.Put(cli, p, string(v))
				}
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
