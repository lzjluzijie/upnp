package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/prestonTao/upnp"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:"upnp",
		Usage:"A simple upnp tool",
		Version:"0.0.1",
		Author:"Halulu",
		Email:"lzjluzijie@gmail.com",
		Action:func(c *cli.Context) (err error) {
			internalPort, err := strconv.Atoi(c.Args().Get(0))
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			externalPort, err := strconv.Atoi(c.Args().Get(1))
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Println(internalPort)

			u := new(upnp.Upnp)
			err = u.SearchGateway()
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			err = u.ExternalIPAddr()
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			fmt.Println("internet ip address: ", u.GatewayOutsideIP)

			err = u.AddPortMapping(internalPort, externalPort, "TCP")
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			fmt.Println("success")

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil{
		fmt.Println(err.Error())
	}
}
