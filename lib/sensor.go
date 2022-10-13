package lib

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
	log "github.com/sirupsen/logrus"
	"io"
	"time"
)

var Token string

func JsonDecoder(reader io.Reader) (map[string]interface{}, error) {
	var m map[string]interface{}
	return m, json.NewDecoder(reader).Decode(&m)
}

func Sensor(cont string, port int, site string, user string, pass string, debug bool) {
	client := req.C().
		SetUserAgent("Unifi PRTG Sensor").
		SetTimeout(1000 * time.Millisecond).
		EnableInsecureSkipVerify().
		DevMode()

	resp, err := client.R().
		SetBodyJsonString(fmt.Sprintf(`{"username":"%v", "password":"%v"}`, user, pass)).
		EnableDump().
		Post(fmt.Sprintf("https://%v:%v/api/auth/login", cont, port))
	if err != nil {
		log.Errorln(err)
		return
	}
	if resp.IsSuccess() {
		// fmt.Println(resp.Body)
		Token = resp.Header.Get("set-cookie")
		return
	}
	if resp.IsError() {
		fmt.Printf("<prtg>\n<error>1</error>\n<text>Authentication Failed: %v</text>\n</prtg>", resp.Status)
		// log.Errorln(resp.StatusCode)
		return
	}

	switchInfo, err := client.R().
		SetHeader("Cookie", Token).
		EnableDump().
		Get(fmt.Sprintf("https://%v:%v/proxy/network/api/s/%v/stat/device", cont, port, site))
	if err != nil {
		log.Errorln(err)
		return
	}
	if switchInfo.IsSuccess() {
		body, err := JsonDecoder(switchInfo.Body)
		if err != nil {
			log.Errorln(err)
		}
		log.Printf("Body : %v", body)

	}
}
