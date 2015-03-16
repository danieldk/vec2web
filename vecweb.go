package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/danieldk/go2vec"
)

func assetOrError(assetName string, w http.ResponseWriter) {
	if data, err := Asset(assetName); err == nil {
		w.Write(data)
	} else {
		http.Error(w, fmt.Sprintf("Cannot read asset '%s'", assetName), 500)
	}
}

func createAnalogy(vecs *go2vec.Vectors) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		word1 := strings.TrimSpace(r.FormValue("word1"))
		word2 := strings.TrimSpace(r.FormValue("word2"))
		word3 := strings.TrimSpace(r.FormValue("word3"))

		result, err := vecs.Analogy(word1, word2, word3, 10)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		resultJSON, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write(resultJSON)
	}
}

func createSimilarity(vecs *go2vec.Vectors) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		word := strings.TrimSpace(r.FormValue("word"))

		result, err := vecs.Similarity(word, 10)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		resultJSON, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write(resultJSON)
	}
}

var httpBind = flag.String("http", ":8080", "Host/port to bind to")

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "Usage: vecweb [OPTION...] vectors.bin")
		flag.PrintDefaults()
		os.Exit(1)
	}

	f, err := os.Open(flag.Arg(0))
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	vecs, err := go2vec.ReadVectors(bufio.NewReader(io.Reader(f)))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded vectors from %s", flag.Arg(0))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		assetOrError("static/index.html", w)
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/about" {
			http.NotFound(w, r)
			return
		}

		assetOrError("static/about.html", w)
	})
	http.HandleFunc("/analogy", createAnalogy(vecs))
	http.HandleFunc("/distance", createSimilarity(vecs))

	log.Printf("Starting to serve from %s", *httpBind)
	http.ListenAndServe(*httpBind, nil)
}
