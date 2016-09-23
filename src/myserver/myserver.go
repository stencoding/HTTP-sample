package myserver

import (
	"fmt"
	"net/http"
	"sync"
)

type MyHandler struct {
	sync.Mutex
	count int
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var count int

	h.Lock()
	count = ++h.count
	h.Unlock()

	fmt.Fprintf(w, "Visitor count: %d.", count)
}