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
	Addr := fmt.Sprintf(":%d", s.Key("Port").MustInt(80))
	URL_path := s.Key("URL_path").String()
	URL_path = "/" + strings.Trim(URL_path, "/") + "/"
	Work_dir := s.Key("Work_dir").MustString(".well-known")
	Https_host := s.Key("Https_host").String()

	ToHttps := func(w http.ResponseWriter, r *http.Request) {
		host := Https_host
		if host == "" {
			host = strings.Split(r.Host, ":")[0]
		}
		host = "https://" + host + r.RequestURI
		http.Redirect(w, r, host, 301)
	}

	FileHandler := http.StripPrefix(URL_path,
		http.FileServer(http.Dir(Work_dir)))
	FileHandleFunc := func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.Method, "-", r.RemoteAddr, r.URL)
		FileHandler.ServeHTTP(w, r)
	}

	if URL_path != "//" {
		http.HandleFunc(URL_path, FileHandleFunc)
	}
	http.HandleFunc("/", ToHttps)
	http.ListenAndServe(Addr, nil)
}
