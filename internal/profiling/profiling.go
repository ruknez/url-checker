package profiling

import (
	"fmt"
	"net/http"
	"net/http/pprof"
)

func GoPProfStart(port int) {
	r := http.NewServeMux()

	// Регистрация pprof-обработчиков
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	go http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
