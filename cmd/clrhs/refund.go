package main

import (
	"fmt"
	"net/url"
	"strconv"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	refundCMD    = kingpin.Command("refund", "Refund a captured amount")
	refundAmount = refundCMD.Flag("amount", "Amount to refund").Short('a').Int()
	refundID     = refundCMD.Arg("id", "authorization ID").Required().String()
)

func doRefund() (e error) {
	body := make(url.Values)

	if refundAmount != nil && *refundAmount > 0 {
		body.Set("amount", strconv.Itoa(*refundAmount))
	}

	path := fmt.Sprintf("/authorizations/%s/refunds", *refundID)

	trx, e := sendRequest(path, body.Encode())
	if e != nil {
		return
	}

	fmt.Println(trx.Pretty())

	return
}
