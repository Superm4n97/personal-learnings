package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-type",
		"text/html",
	)
	io.WriteString(
		res,
		`<DOCTYPE html>
<html>
  <head>
      <title>Hello World</title>
  </head>
  <body>
      Hello Rasel!
  </body>
</html>`,
	)
}

func main() {
	port := flag.Int("port", 9000, "server port")
	fmt.Println("Before parsing: ", *port)
	flag.Parse()
	fmt.Println("After parsing: ", *port)
	http.HandleFunc("/hello", hello)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
