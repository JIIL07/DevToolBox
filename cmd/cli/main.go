package main

import (
	"log"

	"github.com/JIIL07/devtoolbox/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}
