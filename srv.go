package main

import (
        "html/template"
        "net/http"
)

var t *template.Template

func init() {
        t = template.Must(template.ParseFiles("static/index.html"))
}

func Index(w http.ResponseWriter, r *http.Request) {
        t.ExecuteTemplate(w, "static/index.html", http.Dir("static"))
}

func main() {
        http.HandleFunc("/", Index)
        //http.ListenAndServe(":9090", nil)
        //http.ListenAndServe(":9090", http.FileServer("index.html"))
        http.ListenAndServe(":9090", http.FileServer(http.Dir("static")))
}
