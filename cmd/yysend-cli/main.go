package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yangtools/yysend"
)

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "version", false, "print version information")
	flag.BoolVar(&showVersion, "v", false, "print version information")
	flag.Parse()

	if showVersion {
		fmt.Println(yysend.Version)
		os.Exit(0)
	}
}
