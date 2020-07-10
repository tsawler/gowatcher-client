package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	// StatusHealthy is healthy status
	StatusHealthy = 1
	// StatusProblem is problem status
	StatusProblem = 2
	// StatusHPending is pending status
	StatusPending = 3
	// StatusWarning is warning status
	StatusWarning = 4
)

// Status is the response sent back to goWatcher (as JSON)
type Status struct {
	Action      string    `json:"action"`
	OK          bool      `json:"ok"`
	Status      string    `json:"status"`
	Data        string    `json:"data"`
	DateTime    time.Time `json:"date_time"`
	NewStatusID int       `json:"new_status_id"`
}

// JsonRequest is the json format sent to us by goWatcher
type JsonRequest struct {
	Parameters string `json:"parameters"`
}

// ReportStatus performs a check, and returns JSON response
func ReportStatus(app App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ensure request comes from valid ip address
		remoteIP := GetIP(r)
		if _, ok := app.AllowFrom[remoteIP]; !ok {
			infoLog.Println("Denying access")
			DenyAccess(w, "", "Access denied")
			return
		}

		// parse json
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			errorLog.Println(err)
			DenyAccess(w, "", "Error reading request body")
			return
		}

		var j JsonRequest
		err = json.Unmarshal(body, &j)
		if err != nil {
			errorLog.Println(err)
			DenyAccess(w, "", "Error parsing json")
			return
		}

		// get the action verb
		action := chi.URLParam(r, "action")

		var okay bool
		var status Status
		var msg, data string
		var newStatusID int

		switch action {

		case "disk-space":
			// checking disk space
			ok, m, d, statusID, err := checkDiskSpace(j.Parameters)
			if err != nil {
				DenyAccess(w, action, err.Error())
				return
			}
			okay = ok
			msg = m
			data = d
			newStatusID = statusID
			break

		case "memory":
			// checking memory
			ok, m, d, statusID, err := checkMemory()
			if err != nil {
				DenyAccess(w, action, err.Error())
				return
			}
			okay = ok
			msg = m
			data = d
			newStatusID = statusID
			break

		case "postgres":
			ok, m, d, statusID, err := checkPostgres(j.Parameters)
			if err != nil {
				DenyAccess(w, action, err.Error())
				return
			}
			okay = ok
			msg = m
			data = d
			newStatusID = statusID
			break

		case "mariadb":
			ok, m, d, statusID, err := checkMariaDB(j.Parameters)
			if err != nil {
				DenyAccess(w, action, err.Error())
				return
			}
			okay = ok
			msg = m
			data = d
			newStatusID = statusID
			break

		case "redis":
			// checking redis status
			ok, m, d, statusID, err := checkRedis(j.Parameters)
			if err != nil {
				DenyAccess(w, action, err.Error())
				return
			}
			okay = ok
			msg = m
			data = d
			newStatusID = statusID
			break

		case "cpu":
			ok, m, d, statusID, err := checkCPU()
			if err != nil {
				DenyAccess(w, action, err.Error())
				return
			}
			okay = ok
			msg = m
			data = d
			newStatusID = statusID
			break

		case "load":
			ok, m, d, statusID, err := checkLoad()
			if err != nil {
				DenyAccess(w, action, err.Error())
				return
			}
			okay = ok
			msg = m
			data = d
			newStatusID = statusID
			break

		case "test":
			// performing connectivity test
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
		status.DateTime = time.Now()
		status.NewStatusID = newStatusID

		out, _ := json.MarshalIndent(status, "", "    ")

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(out)
	}
}

// GetIP gets a requests IP address by reading off the forwarded-for
// header (for proxies) and falls back to use the remote address.
func GetIP(r *http.Request) string {
	if !*inProduction {
		return "127.0.0.1"
	}

	testIP := r.RemoteAddr
	ip := net.ParseIP(testIP)
	infoLog.Println("IP is", testIP)

	if ip.To4() != nil {
		// this is an ipv4 address, but might have port on end. Split by :
		forwarded := r.Header.Get("X-FORWARDED-FOR")
		if forwarded != "" {
			ex := strings.Split(forwarded, ":")
			return ex[0]
		}
		ex := strings.Split(r.RemoteAddr, ":")
		return ex[0]
	}

	// this is an ipv6 address, possibly in form of [ip]:port. Use custom delimiters
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		ex := strings.FieldsFunc(forwarded, split)
		return ex[0]
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
		OK:       false,
		Status:   msg,
		DateTime: time.Now(),
	}

	out, _ := json.MarshalIndent(status, "", "    ")

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(out)
}
