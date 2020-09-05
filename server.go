package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	port = flag.Int("port", 3000, "port to listen on")
	host = flag.String("host", "localhost", "host/ip to listen on")
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func blue(s string) string {
	return "\033[34m" + s + "\033[0m"
}

func green(s string) string {
	return "\033[33m" + s + "\033[0m"
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println(blue(time.Now().Format("15:04:05") + " -----------------------------------------------------------------------"))

	fmt.Print(green(req.Method), " ", req.URL, "\n")

	fmt.Println(green("\nHeaders"))
	for k, v := range req.Header {
		fmt.Println(k, ":", v[0])
	}

	fmt.Println(green("\nBody"))
	body, err := ioutil.ReadAll(req.Body)
	check(err)
	fmt.Println(string(body))
	fmt.Println()

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type") // "*" isn't respected by chrome at least
	w.Header().Set("Access-Control-Allow-Credentials", "*")
	w.WriteHeader(200)
	io.WriteString(w, "recorded")
}

func main() {
	flag.Parse()

	addr := fmt.Sprintf("%s:%d", *host, *port)

	fmt.Printf("Listening on %s\n", addr)

	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(addr, nil))
}
