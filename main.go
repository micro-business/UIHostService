package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/microbusinesses/Micro-Businesses-Core/common/diagnostics"
	"github.com/microbusinesses/UIHostService/config"
)

var consulAddress string
var consulScheme string
var listeningPort int
var rootDirectory string

func main() {
	flag.StringVar(&consulAddress, "consul-address", "", "The consul address in form of host:port. The default value is empty string.")
	flag.StringVar(&consulScheme, "consul-scheme", "", "The consul scheme. The default value is empty string.")
	flag.IntVar(&listeningPort, "listening-port", 0, "The port the application is serving HTTP request on. The default is zero.")
	flag.StringVar(&rootDirectory, "root-directory", "", "The root directory the of where files to be served are located. The default value is empty string.")
	flag.Parse()

	consulConfigurationReader := config.ConsulConfigurationReader{ConsulAddress: consulAddress, ConsulScheme: consulScheme}

	setConsulConfigurationValuesRequireToBeOverriden(&consulConfigurationReader)

	http.Handle("/", http.FileServer(http.Dir(rootDirectory)))

	if portToListen, err := consulConfigurationReader.GetListeningPort(); err != nil {
		log.Fatal(err.Error())
	} else {
		log.Fatal(http.ListenAndServe(":"+strconv.Itoa(portToListen), nil))
	}
}

func setConsulConfigurationValuesRequireToBeOverriden(consulConfigurationReader *config.ConsulConfigurationReader) {
	diagnostics.IsNotNil(consulConfigurationReader, "consulConfigurationReader", "consulConfigurationReader is nil.")

	if listeningPort != 0 {
		consulConfigurationReader.ListeningPortToOverride = listeningPort
	} else if port, err := strconv.Atoi(os.Getenv("PORT")); err == nil && port != 0 {
		consulConfigurationReader.ListeningPortToOverride = port
	}
}
