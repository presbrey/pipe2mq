package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/streadway/amqp"
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
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := scanner.Text()

		if *escapeBody {
			data = escape(data)
		}

		elt := &Request{
			Headers: amqp.Table{
				"Time": fmt.Sprintf("%d", time.Now().Unix()),
			},
			Body: data,
		}
		incoming <- elt
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}
