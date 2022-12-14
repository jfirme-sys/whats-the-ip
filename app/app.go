package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func GenApp() *cli.App {
	app := cli.NewApp()
	app.Name = "Whats the IP"
	app.Usage = "Search IPS and server names"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IPS",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "server",
			Usage:  "Search server name",
			Flags:  flags,
			Action: searchServer,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServer(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server)
	}
}
