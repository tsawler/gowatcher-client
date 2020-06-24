package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net"
	"net/http"
	"regexp"
	"strings"
)

// Status is the response sent back to goWatcher (as JSON)
type Status struct {
	Action string `json:"action"`
	OK     bool   `json:"ok"`
	Status string `json:"status"`
	Data   string `json:"data"`
}

// ReportStatus performs a check based on a verb, and returns JSON response
func ReportStatus(app App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ensure request comes from valid ip address
		remoteIP := GetIP(r)
		infoLog.Println("IP address is", remoteIP)
		if _, ok := app.AllowFrom[remoteIP]; !ok {
			DenyAccess(w, "", "Access denied")
			return
		}

		// get the action verb
		action := chi.URLParam(r, "action")
		var okay bool
		var status Status
		var msg, data string

		infoLog.Println("Action:", action)

		switch action {

		case "disk-space":
			// checking disk space
			ok, m, d, err := checkDiskSpace(diskToCheck)
			if err != nil {
				DenyAccess(w, action, err.Error())
				return
			}
			okay = ok
			msg = m
			data = d
			break

		case "memory":
			// checking memory
			ok, m, d, err := checkMemory()
			if err != nil {
				DenyAccess(w, action, err.Error())
				return
			}
			okay = ok
			msg = m
			data = d
			break

		case "test":
			// performing connectivity test
			infoLog.Println("Handling a test")
			okay = true
			msg = "Success"
			data = ""
			break

		default:
			DenyAccess(w, action, "Unknown request")
			return
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
	if !inProduction {
		return "127.0.0.1"
	}

	testIP := r.RemoteAddr
	ip := net.ParseIP(testIP)

	if ip.To4() != nil {
		forwarded := r.Header.Get("X-FORWARDED-FOR")
		if forwarded != "" {
			ex := strings.Split(forwarded, ":")
			return ex[1]
		}
		ex := strings.Split(r.RemoteAddr, ":")
		return ex[1]
	}

	ex := strings.FieldsFunc(testIP, split)
	return ex[0]
}

func split(r rune) bool {
	return r == '[' || r == ']'
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

func isIPV4(ip string) bool {
	if strings.Contains(ip, ".") {
		return true
	}
	return false
}
