package main

import (
	"fmt"
	"net/url"
	"strconv"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	captureCMD    = kingpin.Command("capture", "Capture an authentication")
	captureAmount = captureCMD.Flag("amount", "Amount to capture").Short('a').Int()
	captureID     = captureCMD.Arg("id", "authorization ID").Required().String()
)

func doCapture() (e error) {
	body := make(url.Values)

	if captureAmount != nil && *captureAmount > 0 {
		body.Set("amount", strconv.Itoa(*captureAmount))
	}

	path := fmt.Sprintf("/authorizations/%s/captures", *captureID)

	trx, e := sendRequest(path, body.Encode())
	if e != nil {
		return
	}

	fmt.Println(trx.Pretty())

	return
}
