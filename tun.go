package main

import (
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"gopkg.in/mflag.v1"
)

var (
	port    int
	verbose bool
	useURL  bool
)

func init() {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)
	//
	mflag.IntVar(&port, []string{"p", "-port"}, 5051, "a port used by the server")
	mflag.BoolVar(&useURL, []string{"u", "-url"}, false, "take an url as source")
	mflag.BoolVar(&verbose, []string{"v", "-verbose"}, false, "be verbosive")
	mflag.Usage = func() {
		log.Printf("Usage: %s [option] <dir|url>", os.Args[0])
		log.Println("Spawns an http server that serves files from the specified directory or URL.")
		log.Println("\nOPTIONS:")
		mflag.PrintDefaults()
	}
}

func main() {
	mflag.Parse()
	if len(mflag.Args()) < 1 {
		mflag.Usage()
		os.Exit(1)
	}
	port := strconv.Itoa(port)
	if useURL {
		u := mflag.Arg(0)
		uri, err := url.Parse(u)
		if err != nil {
			log.Fatalln("unable to parse URL:", err)
		}
		if len(uri.Scheme) == 0 {
			uri.Scheme = "http"
		}
		http.Handle("/", http.HandlerFunc(proxyHandler(uri)))
	} else {
		dir := mflag.Arg(0)
		if info, err := os.Stat(dir); err != nil {
			log.Fatalln(err)
		} else if !info.IsDir() {
			log.Fatalf("stat %s: not a directory", dir)
		}
		http.Handle("/", http.FileServer(http.Dir(dir)))
	}
	if verbose {
		if ip, err := getIP(); err != nil {
			log.Println("unable to list entrypoints:", err)
		} else {
			log.Println("entrypoint:", "http://"+ip.String()+":"+port)
		}
	}
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

func proxyHandler(u *url.URL) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(u.String() + r.RequestURI)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
		resp.Body.Close()
	}
}

func isHTTPS(u *url.URL) bool {
	return u.Scheme == "https"
}
