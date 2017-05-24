package main

import (
	//"crypto/tls"
	"crypto/tls"
	"fmt"
	"log"
	//"net"
	//"net/http"
	//"net/url"
	//"io/ioutil"
	//"net/http/httputil"
	//"os"
	"github.com/valyala/fasthttp"
	//"time"
)

func main() {

	host := "qnwww2.autoimg.cn"
	request := "/youchuang/g4/M14/8C/46/autohomecar__wKgHy1kJ8hiAOiaUAAE_GaBvKxo295.jpg?imageView2/1/w/400/h/225"
	referer := "http://www.autohome.com.cn"
	urls := fmt.Sprintf("https://%s%s", host, request)
	//req, err := http.NewRequest(http.MethodGet, urls, nil)
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(urls)
	if host != "-" {
		req.Header.Add("Host", host)
	}

	if referer != "" && referer != "-" {
		req.Header.Add("referer", referer)
	}
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cache-Control", "no-cache")

	c := &fasthttp.HostClient{
		Addr:      "157.255.158.82:443",
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
		IsTLS:     true,
		Name:      "zyc https replay",
	}

	resp := fasthttp.AcquireResponse()
	//client := &fasthttp.Client{}

	err := c.Do(req, resp)
	if err != nil {
		panic(err)
	}
	log.Println(resp.StatusCode())

	//statusCode, body, err := c.Get(nil, "http://google.com/foo/bar")
	//if err != nil {
	//	log.Fatalf("Error when loading google page through local proxy: %s", err)
	//}
	//if statusCode != fasthttp.StatusOK {
	//	log.Fatalf("Unexpected status code: %d. Expecting %d", statusCode, fasthttp.StatusOK)
	//}

	//proxyurl := func(*http.Request) (*url.URL, error) {
	//	return url.Parse("https://157.255.158.82:443")
	//}
	//client := http.DefaultClient
	////http.ProxyFromEnvironment()
	//tr := &http.Transport{
	//	Proxy: proxyurl,
	//	//Proxy: http.ProxyFromEnvironment,
	//	DialContext: (&net.Dialer{
	//		Timeout:   30 * time.Second,
	//		KeepAlive: 30 * time.Second,
	//	}).DialContext,
	//	MaxIdleConns:          100,
	//	IdleConnTimeout:       90 * time.Second,
	//	TLSHandshakeTimeout:   10 * time.Second,
	//	ExpectContinueTimeout: 1 * time.Second,
	//	TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
	//	//TLSClientConfig:       &tls.c
	//	DisableCompression: true,
	//}
	//client.Transport = tr
	//client := &http.Client{
	//	Transport: tr,
	//}

	//datareq, _ := httputil.DumpRequest(req, true)
	//log.Println(string(datareq), req.URL.String())

	//resp, err := client.Do(req)
	////resp, err := client.Get(urls)
	//if err != nil {
	//	panic(err)
	//
	//}
	//
	//datare, _ := httputil.DumpResponse(resp, true)
	//log.Println("---------------", string(datare), req.Host, req.RequestURI, req.URL.String())

}
