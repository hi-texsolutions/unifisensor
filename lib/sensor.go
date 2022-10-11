package lib

import (
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

func Sensor(controller *string, port *int, site *string, username *string, pass *string, debug *bool) {
	log.Printf("%v %v %v %v %v %v ", controller, port, site, username, pass, debug)

	client := resty.New()
	resp, err := client.R().EnableTrace().Post("https://")
}
