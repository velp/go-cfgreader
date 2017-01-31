package main
 
import (
	"os"
	"fmt"

	"github.com/urfave/cli"

	"github.com/velp/go-cfgreader/config"
)

func main() {
	var configfile string

	app := cli.NewApp()

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "config, c",
			Value: "",
			Usage: "path to YAML config file",
			Destination: &configfile,
		},
	}

	app.Action = func(c *cli.Context) error {
		if configfile == "" {
			return fmt.Errorf(`Bad argument "config"`)
		}
		cfg, err := config.Parse(configfile)
		if err != nil {
			return err
		}
		fmt.Printf("Config struct: %#v\n", cfg)
		return nil
	}

	app.Run(os.Args)
}