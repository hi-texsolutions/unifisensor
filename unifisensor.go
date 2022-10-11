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
	controller := flag.String("Controller", "", "FQDN of controller")
	port := flag.Int("port", 8443, "Port of the controller")
	site := flag.String("site", "default", "Site name from controller")
	username := flag.String("username", "", "Username to log in with")
	password := flag.String("password", "", "Password used to log in")
	debug := flag.Bool("debug", false, "Add debug logging")
	lib.Sensor(controller, port, site, username, password, debug)
}
