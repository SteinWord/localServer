package main

import (
	"flag"
	"fmt"
	"net/http"
	"text/template"
)

type View struct {
	Port string
}

var Gport string

var Body string = `
var host = localhost:{{.Port}}
`

func NewServer(dir string, port string) {
	Gport = port
	
	http.Handle("/", http.FileServer(http.Dir(dir)))
	http.HandleFunc("/go/go.js", viewJS)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

func viewJS(w http.ResponseWriter, r *http.Request) {
	js := View{Gport}
	tmpl, err := template.New("new").Parse(Body)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, js)
	if err != nil {
		panic(err)
	}
}

func main() {
	var p = flag.String("p", "default", "port")
	var d = flag.String("d", "default", "directory")
	flag.Parse()
	fmt.Println("start server : localhost:", *p, " <=== ", *d)
	fmt.Println("to stop server, pls ctrl-c")
	NewServer(*d, *p)
}
