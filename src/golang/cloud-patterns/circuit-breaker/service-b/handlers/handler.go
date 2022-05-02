package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const defaultPort = "9008"

type IHandler interface {
	ListenAndServe() error
}

type handler struct {
	port string
}

func (h *handler) ListenAndServe() error {
	http.HandleFunc("/get", h.GetValues)
	return http.ListenAndServe(":"+h.port, nil)
}

func (h *handler) GetValues(w http.ResponseWriter, _ *http.Request) {
	rand.Seed(time.Now().UnixNano())
	seconds := rand.Intn(30)
	time.Sleep(time.Duration(seconds))

	data, _ := json.Marshal(`{
		"message": "success",
		"data": true
	}`)
	w.Write(data)
	w.WriteHeader(http.StatusOK)
}

func NewHandler() IHandler {
	port := ""
	if port = os.Getenv("var_service_b_port"); port == "" {
		port = defaultPort
	}

	return &handler{port}
}
