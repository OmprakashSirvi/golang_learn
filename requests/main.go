package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error : ", err)
	}
	log.Panic("panic")
	io.Copy(os.Stdout, resp.Body)
}
