// package main
package httpserver

import (
	"fmt"
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html>
			<html>
				<head>
					<title>Hello World</title>
				</head>
				<body>
					Hello World!
				</body>
			</html>`,
	)
}

func main2() {
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
