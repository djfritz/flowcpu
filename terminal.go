package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data bytes.Buffer
		io.Copy(&data, r.Body)
		r.Body.Close()
		v, _ := strconv.Atoi(data.String())
		fmt.Printf("%c", v)
	})

	log.Fatal(http.ListenAndServe(":4444", nil))
}
