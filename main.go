package main

import (
	"log"
	"net/http"

	"github.com/guaiu/myGee/web"
)

func main() {
	engine := new(web.Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
