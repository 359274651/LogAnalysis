package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/hpcloud/tail"
	"github.com/valyala/fasthttp"
	"log"
	"strings"
	"time"
)

var (
	httpslogpath *string //https 访问日志
	sourceAddr   *string //访问的测试的环境的地址
	keepalive    *bool
)

//$remote_addr|$remote_user|[$time_local]|$host|$request|$status|$body_bytes_sent|$http_referer|'
//                     '$http_user_agent|$http_x_forwarded_for|$request_time|$upstream_response_time|'
//                     '$upstream_connect_time|$upstream_header_time|$upstream_http_via|$upstream_addr|'
//                     '$upstream_http_x_e_reqid|$upstream_http_x_m_reqid'

//103.20.32.163|-|[03/May/2017:17:33:45 +0800]|s1.cdn.xiangha.com|GET /favicon.ico HTTP/2.0|401|399|https://s1.cdn.xiangha.com/caipu/201507/0417/041728029421.jpg/MjgweDIyMA|Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36|-|0.196|0.196|0.000|0.196|http/1.1 cdn-cnc-gddg-dg-8 (ApacheTrafficServer/6.2.0 [cMsSf ])|127.0.0.1:8080|1493804025673921299:4639915|5lgAAEs4eMDYDrsU

//
func init() {
	httpslogpath = flag.String("logpath", "", "nginx 对应的https 的日志路径")
	sourceAddr = flag.String("addr", "", "需要访问的测试环境的端口和ip eg:127.0.0.1:443")
	keepalive = flag.Bool("ka", false, "keep alive default is false to close")
}

func main() {
	flag.Parse()
	lines := make(chan []string, 100)
	if *httpslogpath == "" || *sourceAddr == "" {
		flag.PrintDefaults()

		return
	}
	t, err := tail.TailFile(*httpslogpath, tail.Config{Poll: true, Follow: true})
	if err != nil {
		panic(err)
	}
	for w := 1; w <= 5; w++ {
		go generateUrl(w, lines, *sourceAddr, *keepalive)
	}
	for line := range t.Lines {
		data := line.Text
		if data == "" {
			continue
		} else {
			lineslice := strings.Split(data, "|")
			if len(lineslice) < 18 {
				log.Println("长度不够", data)
				continue
			}
			if lineslice[3] == "~.*" {
				log.Println("请求格式不正确直接跳过 ：", data)
				continue
			}
			lines <- lineslice
		}

		//121.31.64.106|-|[03/May/2017:18:28:49 +0800]|media.finger66.com|GET /posts/28008990000/MTQ5MTQ1OTQ2ODc0Nw==.jpg?imageView2/1/w/320/q/100 HTTP/1.1|

		//fmt.Println(line.Text)
	}

}

func generateUrl(id int, linedetail <-chan []string, source string, ka bool) {
	var infoline []string
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			log.Println("stop work id: ", id, " -zyc recover error: ", err, "-----", strings.Join(infoline, "|"))
		}
	}()
	for lines := range linedetail {
		infoline = lines
		host := lines[3]
		request := (strings.Split(lines[4], " "))[1]
		referer := lines[7]
		urls := fmt.Sprintf("https://%s%s", host, request)

		c := &fasthttp.HostClient{
			Addr:                source,
			TLSConfig:           &tls.Config{InsecureSkipVerify: true},
			IsTLS:               true,
			Name:                "zyc https replay",
			MaxIdleConnDuration: 25 * time.Second,
		}
		req := fasthttp.AcquireRequest()
		if ka {
			req.Header.Add("Connection", "keep-alive")
		}

		req.Header.Add("Cache-Control", "no-cache")
		//req.Header.Add("Host", host)
		if referer != "" && referer != "-" {
			req.Header.Add("Referer", referer)
		}

		resp := fasthttp.AcquireResponse()
		//client := &fasthttp.Client{}

		req.SetRequestURI(urls)
		err := c.Do(req, resp)
		if err != nil {
			log.Println("zyc error: ", err.Error())
			log.Println("waiting work id: ", id, "--", " urls:", urls, " host:", host, " referer:", referer, " request: ", request)
			continue
		}
		if resp.StatusCode() != 200 {
			log.Println("exec work id: ", id, "--", resp.StatusCode(), " urls:", urls, " host:", host, " referer:", referer, " request: ", request)
		}
	}

}

//func newTransport(proxy string) http.Transport {
//	proxyurl := func(*http.Request) (*url.URL, error) {
//		return url.Parse("http://" + proxy)
//	}
//	return http.Transport{
//		Proxy: proxyurl,
//		DialContext: (&net.Dialer{
//			Timeout:   30 * time.Second,
//			KeepAlive: 30 * time.Second,
//		}).DialContext,
//		MaxIdleConns:          100,
//		IdleConnTimeout:       90 * time.Second,
//		TLSHandshakeTimeout:   10 * time.Second,
//		ExpectContinueTimeout: 1 * time.Second,
//		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
//	}
//}
