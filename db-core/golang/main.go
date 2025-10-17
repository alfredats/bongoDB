package main

import (
	"flag"
	"log"
	"os"

	"github.com/alfredats/bongoDB/db-core/golang/src"
)

func main() {
	var portnum *int = flag.Int("port", 8080, "Port to run the server on")
	var libPath *string = flag.String("libpath", "", "Path to the shared library")
	flag.Parse()

	if libPath == nil || *libPath == "" {
		*libPath = os.Getenv("BONGO_LIBPATH")
		if *libPath == "" {
			log.Fatal("Please provide the path to the shared library using '-libpath' or as environment variable 'BONGO_LIBPATH")
			os.Exit(1)
		}
	}

	var instance src.BongoInstance
	instance.Start(*portnum, *libPath)

	os.Exit(0)
}
