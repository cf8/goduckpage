package main

import (
    "fmt"
    "regexp"
    "net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
    query := req.FormValue("q")
    r := regexp.MustCompile(`!\w+`)
    ddg_url := "https://duckduckgo.com/?q=%s"
    startpage_url := "https://startpage.com/do/search/?query=%s"

    if r.MatchString(query) == true {
      http.Redirect(w, req, fmt.Sprintf(ddg_url, query), http.StatusFound)
    } else {
      http.Redirect(w, req, fmt.Sprintf(startpage_url, query), http.StatusFound)
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServeTLS("127.0.0.1:9001", "/etc/nginx/ssl/server.crt", "/etc/nginx/ssl/server.key", nil)
}
