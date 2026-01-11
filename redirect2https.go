// http4Lets
package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"gopkg.in/ini.v1"
)

func main() {
	cfg, err := ini.Load("redirect2https.ini")
	if err != nil {
		log.Fatalf("Fail to read ini file: %v\n", err)
	}
	s := cfg.Section("")
	addr := fmt.Sprintf(":%d", s.Key("Port").MustInt(80))
	urlPath := s.Key("URL_path").String()
	urlPath = "/" + strings.Trim(urlPath, "/") + "/"
	workDir := s.Key("Work_dir").MustString(".well-known")
	httpsHost := s.Key("Https_host").String()

	toHttps := func(w http.ResponseWriter, r *http.Request) {
		host := httpsHost
		if host == "" {
			host = strings.Split(r.Host, ":")[0]
		}
		host = "https://" + host + r.RequestURI
		http.Redirect(w, r, host, http.StatusMovedPermanently)
	}

	f := http.StripPrefix(urlPath, http.FileServer(http.Dir(workDir)))
	fileHandlerFunc := func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.Method, "-", r.RemoteAddr, r.URL)
		f.ServeHTTP(w, r)
	}

	if urlPath != "//" {
		http.HandleFunc(urlPath, fileHandlerFunc)
	}
	http.HandleFunc("/", toHttps)
	http.ListenAndServe(addr, nil)
}
