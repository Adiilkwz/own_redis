package protocol

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Port int
}

func ParseFlags() Config {
	helpFlag := flag.Bool("help", false, "Show this screen.")
	portFlag := flag.Int("port", 8080, "Port number.")

	flag.Usage = func() {
		fmt.Print(`Own Redis

Usage:
  own-redis [--port <N>]
  own-redis --help

Options:
  --help       Show this screen.
  --port N     Port number.
`)
	}

	flag.Parse()

	if *helpFlag {
		flag.Usage()
		os.Exit(0)
	}

	return Config{
		Port: *portFlag,
	}
}
