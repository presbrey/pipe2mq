package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var (
	escapeBody = flag.Bool("escapeBody", false, "request body will be Go-escaped")
)

func escape(s string) (r string) {
	if len(s) > 0 {
		r = strconv.Quote(s)
		r = strings.Replace(r[1:len(r)-1], `\"`, `"`, -1)
	}
	return
}

func init() {
	flag.Parse()
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := scanner.Text()

		if *escapeBody {
			data = escape(data)
		}

		elt := newRequest(data)
		incoming <- elt
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}
