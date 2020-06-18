package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	authCommand   = kingpin.Command("auth", "Perform an authorization")
	authorization = authCommand.Flag("authorization", "File containing authorization in JSON format").Short('a').Required().String()
	rreq          = authCommand.Flag("rreq", "3DSv2 RReq file").Short('r').String()
)

func parseAuthorization() (body url.Values, e error) {
	data := make(map[string]string)
	b, e := ioutil.ReadFile(*authorization)
	if e != nil {
		return
	}

	e = json.Unmarshal(b, &data)
	if e != nil {
		return
	}

	body = make(url.Values)

	for k, v := range data {
		body.Add(k, v)
	}

	if *rreq == "" {
		return
	}

	b, e = ioutil.ReadFile(*rreq)
	if e != nil {
		return
	}

	body.Add("card[3dsecure][v2][rreq]", string(b))

	return
}

func doAuthorization() (e error) {
	bodyValues, e := parseAuthorization()
	if e != nil {
		return
	}

	body := bodyValues.Encode()
	path := "/authorizations"

	trx, e := sendRequest(path, body)
	if e != nil {
		return
	}

	fmt.Println(trx.Pretty())

	return
}
