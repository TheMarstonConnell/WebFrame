package main

import (
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	p := "static/index.html"
	http.ServeFile(w, r, p)
}
func getWASM(w http.ResponseWriter, r *http.Request) {
	p := "static/page.wasm"
	w.Header().Add("Content-Type", "application/wasm")
	http.ServeFile(w, r, p)
}
func getWASMJS(w http.ResponseWriter, r *http.Request) {
	p := "static/wasm_exec.js"
	http.ServeFile(w, r, p)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/page.wasm", getWASM)
	mux.HandleFunc("/wasm_exec.js", getWASMJS)

	err := http.ListenAndServe(":3333", mux)
	if err != nil {
		return
	}

}
