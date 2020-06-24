package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/willf/pad"
)

var (
	Info *log.Logger
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func Init(infoHandle io.Writer) {

	Info = log.New(infoHandle,
		"[INFO] ",
		log.Ldate|log.Ltime)
}

func ServeGetHttp(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	lrw := NewLoggingResponseWriter(w)
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`Tupy Server version 42 - response code: L3Zhci9ydW4vbmZxYXNpYS5nbwo=
Yay, congrats, use the code: >>> cat <<<<`))
	case "POST":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`hack3r`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`404 page not found`))
	}
	elapsed := time.Since(start)
	Info.Printf("| %d | %s | %s | %s%s\n", lrw.statusCode, pad.Left(elapsed.String(), 12, " "), pad.Left(r.RemoteAddr, 14, " "), pad.Right(r.Method, 10, " "), r.RequestURI)
}

func ServeHttpWithParams(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	lrw := NewLoggingResponseWriter(w)
	switch r.Method {
	case "POST":
		switch r.Header.Get("Nfq-Token") {
		case "42":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`Voilà, you finished everything, after finish your application take a look on these videos:
	 - https://www.youtube.com/watch?v=Y27Agvg56YE
	 - https://www.youtube.com/watch?v=dcQeOR9Aw_0
	 - https://www.youtube.com/watch?v=IWZ1ewE0Q0A
	 - https://www.youtube.com/watch?v=V5jHHYFqt9o

Also, paste this entire message as the response \o/`))
		default:
			w.Header().Set("Nfq-Token", "<answer-life-universe-everything?>")
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(`FORBIDDEN, use the code: >>> response-headers <<<`))
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`404 page not found`))
	}
	elapsed := time.Since(start)
	Info.Printf("| %d | %s | %s | %s%s\n", lrw.statusCode, pad.Left(elapsed.String(), 12, " "), pad.Left(r.RemoteAddr, 14, " "), pad.Right(r.Method, 10, " "), r.RequestURI)
}

func main() {
	Init(os.Stdout)
	data, err := ioutil.ReadFile("welcome")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
	r := mux.NewRouter()
	r.HandleFunc("/", ServeGetHttp)
	r.HandleFunc("/42", ServeHttpWithParams)
	log.Fatal(http.ListenAndServe(":8080", r))
}
