package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"onlyGetPostProxy/conf"
	"fmt"
	"log"
)
var config = conf.NewConfig()

func NewHostsReverseProxy(target url.URL) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		// 获取_methods参数
		queryForm, err := url.ParseQuery(req.URL.RawQuery)
		if err == nil && len(queryForm[config.MethodKey]) > 0 {
			req.Method = strings.ToUpper(queryForm[config.MethodKey][0])
		}
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
	}
	return &httputil.ReverseProxy{Director: director}
}
func main() {
	config.ReadConfig()
	proxy := NewHostsReverseProxy(url.URL{
		Scheme: config.Scheme,
		Host:   config.ServerHost,
	})
	fmt.Println("runing at " + config.ListenPort)
	fmt.Println("proxy to " + config.Scheme + "://" + config.ServerHost)
	log.Fatal(http.ListenAndServe(":" + config.ListenPort, proxy))
}