package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

// Gerar cria a aplicação
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e nomes de servidores na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}
	app.Flags = flags

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na internet",
			Flags:  flags,
			Action: buscarIPs,
		},
		{
			Name:   "servidores",
			Usage:  "Busca nomes de servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}
	return app
}

func buscarIPs(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
