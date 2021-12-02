package main

import (
	"fmt"
	"os"

	"brandonplank.org/parcelilib"
	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("parceli", "Universal package tracker written in Go") // Bin description

	// Args, make sure to deref when using
	var verbose *bool = parser.Flag("v", "verbose", &argparse.Options{Required: false, Help: "Prints more info"})
	var service *string = parser.Selector("s", "service", []string{"ups", "usps", "fedex"}, &argparse.Options{Required: true, Help: "Your package service"})
	var tracking *int = parser.Int("t", "tracking", &argparse.Options{Required: true, Help: "Your tracking number"})

	error := parser.Parse(os.Args)
	if error != nil {
		fmt.Print(parser.Usage(error))
		return
	}

	parcelilib.Parceli(*service, *tracking, *verbose)
}
