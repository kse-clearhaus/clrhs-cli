package main

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	dryrun = kingpin.Flag("dryrun", "Do a dryrun and don't send a request").Short('n').Bool()

	apikey = kingpin.Flag("apikey", "Gateway API key").Short('k').
		Envar("APIKEY").Required().String()
	signingkey = kingpin.Flag("signingkey", "Signing key used for signing requests").Short('s').
			Envar("SIGNINGKEY").Required().String()
	privkey = kingpin.Flag("privkey", "RSA private key for signing").Envar("PRIVKEY").Required().
		String()

	host     = kingpin.Flag("host", "Gateway host").Default("https://gateway.clearhaus.com").String()
	httpHost = kingpin.Flag("httphost", "HTTP HOST header value").String()
)

func main() {
	var e error

	kingpin.HelpFlag.Short('h')

	switch kingpin.Parse() {
	case authCommand.FullCommand():
		e = doAuthorization()
	case captureCMD.FullCommand():
		e = doCapture()
	case refundCMD.FullCommand():
		e = doRefund()
	case voidCMD.FullCommand():
		e = doVoid()
	}

	if e != nil {
		fmt.Println(e.Error())
	}
}
