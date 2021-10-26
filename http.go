package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/NYTimes/gziphandler"
)

func ServeDir(prefix, path string) {
	handler := http.StripPrefix(prefix, http.FileServer(http.Dir(path)))
	_handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Server", "https://github.com/weaming/static-file-server")
		handler.ServeHTTP(w, r)
	})
	http.Handle(prefix, gziphandler.GzipHandler(_handler))
}

func GetIntranetIP() (rv []string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				rv = append(rv, ipnet.IP.String())
			}
		}
	}
	return
}
