package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
)

func healthcheckHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func registerPProfHandlers(r *http.ServeMux) {
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)
}

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/healthz", healthcheckHandler)
	registerPProfHandlers(m)
	if err := http.ListenAndServe(":8080", m); err != nil {
		log.Fatal(err)
	}
}
