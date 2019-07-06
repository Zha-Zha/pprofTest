package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func init() {
	http.HandleFunc("/concat/", concatHandler)
}

type concatResponse struct {
	Count    int    `json:"count"`
	String   string `json:"string"`
	Original string `json:"original"`
}

func concatHandler(w http.ResponseWriter, r *http.Request) {
	str := r.FormValue("str")
	count, err := strconv.Atoi(r.FormValue("count"))
	if err != nil {
		count = 10
	}

	resp := concatResponse{Count: count, Original: str}

	t := r.FormValue("type")
	switch t {
	case "buffer":
		resp.String = concatWithBuffer(str, count)
	case "plus":
		resp.String = concatWithPlus(str, count)
	case "sprintf":
		resp.String = concatWithSprintf(str, count)
	default:
		resp.String = concatWithPlus(str, count)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func concatWithPlus(str string, count int) (ret string) {
	for i := 0; i < count; i++ {
		ret = str + ret
	}
	return ret
}

func concatWithBuffer(str string, count int) (ret string) {
	data := bytes.NewBuffer([]byte{})
	for i := 0; i < count; i++ {
		data.WriteString(str)
	}
	return data.String()
}

func concatWithSprintf(str string, count int) (ret string) {
	for i := 0; i < count; i++ {
		ret = fmt.Sprintf("%s%s", ret, str)
	}
	return ret
}
