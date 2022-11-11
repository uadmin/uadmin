package uadmin

import (
	"fmt"
	"math"
	"net/http"
	"strings"
	"time"
)

func trailAPIHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the user is an admin
	session := IsAuthenticated(r)
	if session == nil {
		w.WriteHeader(403)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "access denied",
		})
		return
	}
	Preload(session)
	if !session.User.Admin {
		w.WriteHeader(403)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "access denied",
		})
		return
	}
	f, ok := w.(http.Flusher)

	w.Header().Set("Connection", "Keep-Alive")
	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Content-type", "text/html; charset=utf-8")

	limit := 4096

	if len(trailBytes) < limit {
		temp := append(trailBytes, []byte(strings.Repeat(" ", (limit)-len(trailBytes)))...)
		// temp = append(temp, '\r', '\n')
		// temp = append([]byte(fmt.Sprintf("%x\r\n", len(temp))), temp...)
		w.Write(temp)
	} else {
		// temp := append(trailBytes, '\r', '\n')
		// w.Write(append([]byte(fmt.Sprintf("%x\r\n", len(temp))), temp...))
		// w.Write(trailBytes)
		for i := 0; i < len(trailBytes); i += limit {
			temp := trailBytes[i:int(math.Min(float64(i+limit), float64(len(trailBytes)-1)))]
			if len(temp) < 4096 {
				w.Write(append(temp, []byte(strings.Repeat(" ", limit-len(temp)))...))
			} else {
				w.Write(temp)
			}
		}
	}

	if ok {
		f.Flush()
	} else {
		fmt.Println("Not flusher")
	}

	// time.Sleep(time.Second * 1)

	c := make(chan string, 250)
	RegisterTrailChan(c)
	msg := ""

	for {
		t := time.After(time.Second * 1)
		select {
		case msg = <-c:
			// fmt.Fprintln(w, msg)
			if len(msg) < (limit) {
				msg += strings.Repeat(" ", (limit)-len(msg))
			}
			// msg += "\r\n"
			// w.Write(append([]byte(fmt.Sprintf("%x\r\n", len(msg)-2)), msg...))
			w.Write([]byte(msg))
			if ok {
				f.Flush()
			}
			// fmt.Println("flushed " + msg)
		case <-t:
		}
	}
}
