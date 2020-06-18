package main

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	voidCMD = kingpin.Command("void", "Void an uncaptured authorization")
	voidID  = voidCMD.Arg("id", "authorization ID").Required().String()
)

func doVoid() (e error) {
	path := fmt.Sprintf("/authorizations/%s/voids", *voidID)

	trx, e := sendRequest(path, "")
	if e != nil {
		return
	}

	fmt.Println(trx.Pretty())

	return
}
