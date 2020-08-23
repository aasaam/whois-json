package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Usage = "whois-json"
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		{
			Name:  "json",
			Usage: "json that containe the domain result for who is",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "domain", Required: true, Aliases: []string{"d"}, Usage: "domain want to check"},
			},
			Action: func(c *cli.Context) error {
				domain := c.String("domain")
				domainType, e := DomainValidation(domain)
				if e != nil {
					return cli.Exit(e, 128)
				}
				result, e := DomainParse(domainType)
				if e != nil {
					return cli.Exit(e, 1)
				}
				json, _ := json.Marshal(result)
				fmt.Println(string(json))
				return nil
			},
		},
		{
			Name:  "validate",
			Usage: "validate domain and return domain data",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "domain", Required: true, Aliases: []string{"d"}, Usage: "domain want to check"},
			},
			Action: func(c *cli.Context) error {
				domain := c.String("domain")
				domainType, e := DomainValidation(domain)
				if e != nil {
					return cli.Exit(e, 128)
				}
				json, _ := json.Marshal(domainType)
				fmt.Println(string(json))
				return nil
			},
		},
		{
			Name:  "webserver",
			Usage: "HTTP Server for REST API",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "port", Aliases: []string{"p"}, EnvVars: []string{"AASAAM_WEBSERVER_PORT"}, Value: "9000", DefaultText: "9000", Usage: "HTTP port want to listen"},
				&cli.StringFlag{Name: "base-url", Aliases: []string{"b"}, EnvVars: []string{"AASAAM_WEBSERVER_BASEURL"}, Value: "/", DefaultText: "/", Usage: "Base URL to serve HTTP endpoints"},
				&cli.StringFlag{Name: "username", Aliases: []string{"user"}, EnvVars: []string{"AASAAM_WEBSERVER_BASIC_USERNAME"}, Value: "username", DefaultText: "username", Usage: "Basic auth username"},
				&cli.StringFlag{Name: "password", Aliases: []string{"pass"}, EnvVars: []string{"AASAAM_WEBSERVER_BASIC_PASSWORD"}, Value: "password", DefaultText: "password", Usage: "Basic auth password"},
			},
			Action: func(c *cli.Context) error {
				app, _ := HTTPServer(
					c.String("base-url"),
					c.String("username"),
					c.String("password"),
					true,
				)

				app.Listen(c.Int("port"))
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
