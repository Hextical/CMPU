package main

import (
	"log"
	"strconv"

	"github.com/valyala/fasthttp"
)

var api = "https://addons-ecs.forgesvc.net/api/v2/"

func connectWithHash(jarFingerprint int) []byte {

	//log.Printf("\nConnecting with fingerprint: %d", jarFingerprint)

	requestbody := []byte("[" + strconv.Itoa(jarFingerprint) + "]")
	req := fasthttp.AcquireRequest()
	req.URI().Update(api + "fingerprint")
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")
	req.SetBody(requestbody)

	res := fasthttp.AcquireResponse()
	err := fasthttp.Do(req, res)

	if err != nil {
		log.Println(err)
	}

	body := res.Body()

	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res)

	return body

}

func connectWithProjectID(projectID string) []byte {
	//log.Printf("\nConnecting with projectID: %s", projectID)

	req := fasthttp.AcquireRequest()
	req.URI().Update(api + "addon/" + projectID + "/files")
	req.Header.SetMethod("GET")

	res := fasthttp.AcquireResponse()
	err := fasthttp.Do(req, res)

	if err != nil {
		log.Println(err)
	}

	body := res.Body()

	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res)

	return body

}
