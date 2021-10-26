package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	fp "path/filepath"
	"strings"
)

const DEFAULT_PW = "admin"

var size int

func main() {
	var LISTEN = flag.String("listen", ":8000", "Listen [host]:port, default bind to 0.0.0.0")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] ROOT\nThe ROOT is the directory to be serve.\n\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	// check the directory path
	ROOT, _ := fp.Abs(flag.Arg(0))
	fi, err := os.Stat(ROOT)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if !fi.IsDir() {
		fmt.Fprintln(os.Stderr, "The path should be a directory!!")
		os.Exit(1)
	}

	log.Printf("target direcotry is %v", ROOT)
	ServeDir("/", ROOT)

	log.Printf("Open http://127.0.0.1:%v to start\n", strings.Split(*LISTEN, ":")[1])
	for _, ip := range GetIntranetIP() {
		log.Printf("Your intranet IP: %v ==> http://%v:%v\n", ip, ip, strings.Split(*LISTEN, ":")[1])
	}
	log.Fatal(http.ListenAndServe(*LISTEN, nil))
}
