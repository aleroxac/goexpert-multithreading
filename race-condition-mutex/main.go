package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var visits uint64 = 0

func main() {
	m := sync.Mutex{}
	http.HandleFunc("/mutex", func(w http.ResponseWriter, r *http.Request) {
		m.Lock()
		visits++
		m.Unlock()
		time.Sleep(10 * time.Millisecond)
		w.Write([]byte("Visitante: " + fmt.Sprint(visits)))
	})
	http.ListenAndServe(":8080", nil)
}
