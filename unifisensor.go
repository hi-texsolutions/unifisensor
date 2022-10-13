package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"unifisensor/lib"
)

func init() {
	log.SetLevel(log.InfoLevel)
}

func main() {
	controller := flag.String("c", "", "controller FQDN")
	port := flag.Int("p", 8443, "Controller Port")
	site := flag.String("s", "default", "Site name")
	username := flag.String("u", "", "Username")
	password := flag.String("pw", "", "Password")
	debug := flag.Bool("d", false, "Debug mode")
	flag.Parse()

	lib.Sensor(*controller, *port, *site, *username, *password, *debug)

}
