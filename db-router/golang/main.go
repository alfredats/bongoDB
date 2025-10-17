package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"strings"

	BongoCore "github.com/alfredats/bongoDB/db-core/golang/src"
)

func run_replica(libPath *string, port int) {
	var instance BongoCore.BongoInstance
	instance.Start(port, *libPath)
}

func main() {
	var cliReplicas *int = flag.Int("replicas", 3, "Number of replicas")
	var cliPorts *string = flag.String("ports", "8080", "Ports to run the database instances on")
	var cliLibPath *string = flag.String("libpath", "", "Path to the BongoDB-cpp library")
	flag.Parse()

	if cliLibPath == nil || *cliLibPath == "" {
		*cliLibPath = os.Getenv("BONGO_LIBPATH")
		if *cliLibPath == "" {
			log.Fatal("Please provide path to the shared library using '-libpath' or with the 'BONGO_LIBPATH' env var")
			os.Exit(1)
		}
	}

	if cliPorts == nil || *cliPorts == "" {
		log.Fatal("Please provide the ports to be used through -ports")
		return
	}
	strPorts := strings.Split(*cliPorts, ",")

	intPorts := make([]int, *cliReplicas)
	if len(strPorts) == 1 {
		leadPort, _ := strconv.Atoi(strPorts[0])
		for i := 0; i < *cliReplicas; i++ {
			intPorts[i] = leadPort + i
		}
	} else if len(strPorts) == *cliReplicas {
		for i := 0; i < *cliReplicas; i++ {
			parsed_portnum, err := strconv.Atoi(strPorts[i])
			if err != nil {
				log.Fatal("Failed to parse portnums")
				os.Exit(1)
			}
			intPorts[i] = parsed_portnum
		}
	} else {
		log.Fatal("Input portnums does not match number of replicas requested")
		os.Exit(1)
	}

	log.Printf("Config:\n  numReplicas: %d\n  ports:%v", *cliReplicas, intPorts)

	// TODO: use goroutines here for parallelism
	for i := 0; i < *cliReplicas; i++ {
		run_replica(cliLibPath, intPorts[i])
	}

}
