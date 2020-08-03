package main

import (
	"github.com/Benbentwo-Sandbox/test-go-bin-generic-make-mine-6/app"
	"os"
)

func main() {
	if err := app.Run(nil); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
