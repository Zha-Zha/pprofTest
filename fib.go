package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func init() {
	http.HandleFunc("/fib/", fibHandler)
}

type fibResponse struct {
	Number    int `json:"number"`
	Fibonacci int `json:"fibonacci"`
}

func fibHandler(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.FormValue("n"))
	if err != nil {
		n = 1
	}

	t := r.FormValue("type")
	resp := fibResponse{Number: n, Fibonacci: 0}

	switch t {
	case "recursive":
		resp.Fibonacci = fibRecursive(resp.Number)
	case "iterative":
		resp.Fibonacci = fibIterative(resp.Number)
	default:
		resp.Fibonacci = fibRecursive(resp.Number)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func fibRecursive(n int) int {
	if n < 2 {
		return n
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}

func fibIterative(n int) int {
	if n < 2 {
		return n
	}

	fib := 1
	prevFib := fib
	tmp := 0

	for i := 2; i < n; i++ {
		tmp = fib
		fib += prevFib
		prevFib = tmp
	}

	return fib
}
