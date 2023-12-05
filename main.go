package main

import (
	"TCA-Plugins-test/check"
	"TCA-Plugins-test/scan"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	case "check":
		check.Check()
	case "scan":
		scan.Scan()
	default:
		fmt.Println("Unknown command")
	}

}
