package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strings"
)

// Status is the response sent back to goWatcher (as JSON)
type Status struct {
	Action string `json:"action"`
	OK     bool   `json:"ok"`
	Status string `json:"status"`
}

// Home shows the home (login) screen
func ReportStatus(app App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		remoteIP := getIP(r)
		infoLog.Println("Request came from", remoteIP)

		if _, ok := app.AllowFrom[remoteIP]; !ok {
			denyAccess(w, "", "Access denied")
			return
		}

		okay := false
		action := chi.URLParam(r, "action")

		switch action {
		case "disk-space":
			infoLog.Println("disk space")
		case "memory":
			infoLog.Println("Memory")
		default:
			denyAccess(w, action, "Unknown request")
		}

		status := Status{
			Action: action,
			OK:     okay,
			Status: fmt.Sprintf("Everything's good for %s check", action),
		}

		out, _ := json.MarshalIndent(status, "", "    ")

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(out)
	}
}

// GetIP gets a requests IP address by reading off the forwarded-for
// header (for proxies) and falls back to use the remote address.
func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		ex := strings.Split(forwarded, ":")
		return ex[0]
	}
	ex := strings.Split(r.RemoteAddr, ":")
	return ex[0]
}

func denyAccess(w http.ResponseWriter, action, msg string) {
	status := Status{
		OK:     false,
		Status: msg,
	}

	out, _ := json.MarshalIndent(status, "", "    ")

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(out)
}
