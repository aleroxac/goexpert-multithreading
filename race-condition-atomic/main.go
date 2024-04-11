package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var visits uint64 = 0

func main() {
	http.HandleFunc("/atomic", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddUint64(&visits, 1)
		time.Sleep(10 * time.Millisecond)
		w.Write([]byte("Visitante: " + fmt.Sprint(visits)))
	})
	http.ListenAndServe(":8080", nil)
}
