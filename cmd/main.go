package main

import (
	"flag"
	"fmt"
	"github.com/deepin-cve/tracker/internal/config"
	"github.com/deepin-cve/tracker/pkg/db"
	v0 "github.com/deepin-cve/tracker/pkg/rest/v0"
	"time"
)

var (
	conf  = flag.String("c", "./configs/config.yaml", "the configuration filepath")
	debug = flag.Bool("d", true, "enable debug mode")
	host = flag.String("h","10.20.32.240","host")
	pwd = flag.String("p","deepin20200202@..","the password of mysql")
)

func main() {
	flag.Parse()

	var c = config.GetConfig(*conf)
	db.Init(*host,*pwd)

	go func() {
		for {
			time.Sleep(time.Hour * 10)
			err := db.SessionClean()
			if err != nil {
				fmt.Println("Failed to clean session:", err)
			}
		}
	}()

	err := v0.Route(fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port),
		*debug)
	if err != nil {
		fmt.Println("Failed to route:", err)
	}
}
