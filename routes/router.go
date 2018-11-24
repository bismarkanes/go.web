package router

import (
  "fmt"
  "github.com/julienschmidt/httprouter"
  log "github.com/sirupsen/logrus"
  "net/http"
)

func writeErrorResponse(w http.ResponseWriter, errorMsg []byte, statusCode int) {
  w.WriteHeader(statusCode)
  _, err := fmt.Fprintf(w, "%s", errorMsg)
  if err != nil {
    log.Warning("ERR_WRITE_TO_RESPONSE")
  }
}

func writeResponse(w http.ResponseWriter, data []byte) {
  _, err := fmt.Fprintf(w, "%s", data)
  if err != nil {
    log.Warning("ERR_WRITE_TO_RESPONSE")
  }
}

// Ping sample
func Ping(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  writeResponse(w, []byte("PONG"))
}

// Error sample
func Error(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  writeErrorResponse(w, []byte("ERR_NOT_FOUND"), http.StatusNotFound)
}
