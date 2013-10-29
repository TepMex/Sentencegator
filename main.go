package main

import (
	"os"
	"strings"
)

var WaniKaniApiKey string

func main() {
	for _, arg := range os.Args {
		pair := strings.Split(arg, "=")
		if len(pair) > 2 {
			continue
		}

		switch pair[0] {
		case "--apik":
			WaniKaniApiKey = pair[1]
		}
	}
}
