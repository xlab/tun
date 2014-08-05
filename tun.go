package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var port int

func init() {
	flag.IntVar(&port, "p", 5051, "a port used by the server")
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
	name := flag.Arg(0)
	if info, err := os.Stat(name); err != nil {
		log.Fatalln(err)
	} else if !info.IsDir() {
		log.Fatalf("stat %s: not a directory", name)
	}
	http.Handle("/", http.FileServer(http.Dir(name)))
	if err := http.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
		log.Fatalln(err)
	}
}
