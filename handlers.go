package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"strings"
)

// Status is the response sent back to goWatcher (as JSON)
type Status struct {
	Action string `json:"action"`
	OK     bool   `json:"ok"`
	Status string `json:"status"`
	Data   string `json:"data"`
}

// ReportStatus perfroms a check based on a verb, and returns JSON response
func ReportStatus(app App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		remoteIP := GetIP(r)
		infoLog.Println("Request came from", remoteIP)

		if _, ok := app.AllowFrom[remoteIP]; !ok {
			DenyAccess(w, "", "Access denied")
			return
		}

		action := chi.URLParam(r, "action")
		var okay bool
		var status Status
		var msg, data string

		switch action {
		case "disk-space":
			infoLog.Println("disk space")
			ok, m, d, err := checkDiskSpace(diskToCheck)
			if err != nil {
				DenyAccess(w, action, err.Error())
			}
			okay = ok
			msg = m
			data = d
		case "memory":
			infoLog.Println("Memory")
		default:
			DenyAccess(w, action, "Unknown request")
		}

		status.Action = action
		status.OK = okay
		status.Status = msg
		status.Data = data

		out, _ := json.MarshalIndent(status, "", "    ")

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(out)
	}
}

// GetIP gets a requests IP address by reading off the forwarded-for
// header (for proxies) and falls back to use the remote address.
func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		ex := strings.Split(forwarded, ":")
		return ex[0]
	}
	ex := strings.Split(r.RemoteAddr, ":")
	return ex[0]
}

// DenyAccess generates a JSON error response
func DenyAccess(w http.ResponseWriter, action, msg string) {
	status := Status{
		OK:     false,
		Status: msg,
	}

	out, _ := json.MarshalIndent(status, "", "    ")

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(out)
}
