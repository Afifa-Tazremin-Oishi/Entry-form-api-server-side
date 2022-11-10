package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/xyz/service"
)

var (
	_version = "default"
)

func main() {
	fmt.Println("Starting Account Service ", _version)
	defer fmt.Println("Done....")
	port := flag.Int("p", 8000, "Service listen port")
	// bindAddress := flag.String("b", "localhost", "Bind address")
	bindAddress := flag.String("b", "0.0.0.0", "Bind address")
	verbose := flag.Bool("v", false, "Verbose output")
	flag.Parse()
	if *verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Config file missing")
		fmt.Println("account [flags] <path to config file> ")
		flag.Usage()
		os.Exit(1)
	}
	//Read the config file
	configBytes, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Println("Unable to read config file ", err)
		os.Exit(1)
	}
	if accountService := service.NewMyAllRestService(configBytes, *verbose); accountService != nil {
		stopSignal := make(chan bool)
		termination := make(chan os.Signal)
		signal.Notify(termination, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			<-termination
			fmt.Println("SIGTERM/SIGINT received from os")
			stopSignal <- true
		}()
		accountService.Serve(*bindAddress, *port, stopSignal)
	} else {
		fmt.Println("Unable to start the service ...")
		os.Exit(2)
	}

}
