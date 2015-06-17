package main

import (
	"io"
	"os"

	"github.com/BurntSushi/toml"
)

type config struct {
	HTTP          string
	WordEmbedding []wordEmbedding
}

type wordEmbedding struct {
	Name string
	Path string
}

func defaultConfiguration() *config {
	return &config{
		HTTP: ":8080",
	}
}

func parseConfig(reader io.Reader) (*config, error) {
	config := defaultConfiguration()
	_, err := toml.DecodeReader(reader, config)
	return config, err
}

func readConfigOrExit(filename string) *config {
	f, err := os.Open(filename)
	fatalIfErr("Error opening configuration file", err)
	defer f.Close()

	config, err := parseConfig(f)
	fatalIfErr("Error parsing configuration file", err)

	return config
}
