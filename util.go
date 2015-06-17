package main

import (
	"fmt"
	"log"
	"net/http"
)

func assetOrError(assetName string, w http.ResponseWriter) {
	if data, err := Asset(assetName); err == nil {
		w.Write(data)
	} else {
		http.Error(w, fmt.Sprintf("Cannot read asset '%s'", assetName), 500)
	}
}

func fatalIfErr(prefix string, err error) {
	if err != nil {
		log.Fatalf("%s: %s", prefix, err.Error())
	}
}
