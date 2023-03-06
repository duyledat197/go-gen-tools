package ws

import (
	"net/http"
)

func NewWebSocket(mux *http.ServeMux) {
	hub := newHub()
	go hub.run()
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
}
