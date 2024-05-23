package smplhttpsvr

import (
	"log"
	"net/http"
	"strconv"
)

// Simple http server with friendly api and удобным контролем/управлением.
type httpsvr struct {
	server http.Server
}

// Get new httpsvr stuct object
func NewServer() *httpsvr {
	svr := httpsvr{}
	return &svr
}

func (h *httpsvr) Addr() string {
	return h.server.Addr
}

func (h *httpsvr) Start(ip string, port uint16) {
	var addrStr string = ""
	if len(ip) > 0 {
		addrStr += ip
	}
	log.Printf("Set http server ip: %s", ip)
	if port > 0 {
		addrStr += ":" + strconv.FormatUint(uint64(port), 10)
	}
	log.Printf("Set http server port: %d", port)
	h.server.Addr = addrStr
	h.server.ListenAndServe()
}

func (h *httpsvr) Stop() error {
	return h.server.Close()
}
