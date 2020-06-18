package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type transaction map[string]interface{}

func (r transaction) ID() string {
	var id string
	val, _ := r["id"]
	if val != nil {
		id, _ = val.(string)
	}

	return id
}

func (r transaction) Pretty() string {
	b, _ := json.MarshalIndent(r, "", "  ")

	return string(b)
}

func addSignature(body string, req *http.Request) (e error) {
	der, _ := pem.Decode([]byte(*privkey))

	keyItf, e := x509.ParsePKCS8PrivateKey(der.Bytes)
	if e != nil {
		return
	}

	pkey := keyItf.(*rsa.PrivateKey)

	hashF := sha256.New()

	hashF.Write([]byte(body))
	hash := hashF.Sum(nil)

	sig, e := pkey.Sign(rand.Reader, hash, crypto.SHA256)
	if e != nil {
		return
	}

	hexSig := make([]byte, hex.EncodedLen(len(sig)))
	hex.Encode(hexSig, sig)

	req.Header.Add("Signature", fmt.Sprintf("%s RS256-hex %s", *signingkey, hexSig))

	return
}

func sendRequest(path, body string) (trx transaction, e error) {
	uri, e := url.Parse(*host)
	if e != nil {
		return
	}
	uri.Path = path

	reader := strings.NewReader(body)

	req, e := http.NewRequest("post", uri.String(), reader)
	if e != nil {
		return
	}

	if *httpHost != "" {
		req.Host = *httpHost
	}

	req.SetBasicAuth(*apikey, "")
	e = addSignature(body, req)
	if e != nil {
		return
	}

	if *dryrun {
		req.Write(os.Stdout)
		return
	}

	resp, e := http.DefaultClient.Do(req)
	if e != nil {
		return
	}
	defer resp.Body.Close()

	respBody, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return
	}

	if resp.StatusCode >= 300 {
		e = fmt.Errorf("Unexpected status code %d: %s", resp.StatusCode, respBody)
		return
	}

	trx = make(transaction)
	e = json.Unmarshal(respBody, &trx)

	return
}
