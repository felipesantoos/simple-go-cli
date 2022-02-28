package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Returns a application CLI ready to be executed.
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Interface's Application"
	app.Usage = "Search server IPs and names from Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ips",
			Usage:  "Search IP addresses from Internet",
			Flags:  flags,
			Action: searchIPs,
		},
		{
			Name:   "servers",
			Usage:  "Search the server names from Internet",
			Flags:  flags,
			Action: searchServerNames,
		},
	}

	return app
}

func searchIPs(ctx *cli.Context) {
	host := ctx.String("host")
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServerNames(ctx *cli.Context) {
	host := ctx.String("host")
	serverNames, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, serverName := range serverNames {
		fmt.Println(serverName.Host)
	}
}
