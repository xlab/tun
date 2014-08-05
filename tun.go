package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

var (
	port    int
	verbose bool
)

func init() {
	flag.IntVar(&port, "p", 5051, "a port used by the server")
	flag.BoolVar(&verbose, "v", false, "be verbosive")
	flag.Usage = func() {
		name := os.Args[0]
		fmt.Fprintf(os.Stderr, "Usage: %s [option] <dir>\n", name)
		fmt.Fprintf(os.Stderr, "Spawns an http server that serves files from the specified directory.\n")
		fmt.Fprintf(os.Stderr, "\nOPTIONS:\n")
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(-1)
	}
	port := strconv.Itoa(port)
	dir := flag.Arg(0)
	if info, err := os.Stat(dir); err != nil {
		log.Fatalln(err)
	} else if !info.IsDir() {
		log.Fatalf("stat %s: not a directory", dir)
	}
	if verbose {
		if ip, err := getIP(); err != nil {
			log.Fatalln("unable to list entrypoints:", err)
		} else {
			println("entrypoint:", "http://"+ip.String()+":"+port)
		}
	}
	http.Handle("/", http.FileServer(http.Dir(dir)))
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalln(err)
	}
}

func getIP() (addr net.IP, err error) {
	var addrs []net.Addr
	if addrs, err = net.InterfaceAddrs(); err != nil {
		return
	}
	for i := range addrs {
		if ipnet, ok := addrs[i].(*net.IPNet); ok {
			if ipnet.IP.IsGlobalUnicast() && ipnet.IP.To4() != nil {
				return ipnet.IP, nil
			}
		}
	}
	return nil, errors.New("no public ip found")
}