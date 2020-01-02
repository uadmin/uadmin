package uadmin

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func Handler(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		/*
			Prepare log message. Valid place holders:
			A perfect list (not fully inplemented): http://httpd.apache.org/docs/current/mod/mod_log_config.html
			 %a: Client IP address
			 %{remote}p: Client port
			 %A: Server hostname/IP
			 %{local}p: Server port
			 %U: Path
			 %c: All coockies
			 %{NAME}c: Cookie named 'NAME'
			 %{GET}f: GET request parameters
			 %{POST}f: POST request parameters
			 %B: Respnse length
			 %>s: Response code
			 %D: Time taken in microseconds
			 %T: Time taken in seconds
			 %I: Request length
			 %i: All headers
			 %{NAME}i: header named 'NAME'
		*/
		HTTP_LOG_MSG := HTTPLogFormat
		host, port, _ := net.SplitHostPort(r.RemoteAddr)
		HTTP_LOG_MSG = strings.Replace(HTTP_LOG_MSG, "%a", host, -1)
		HTTP_LOG_MSG = strings.Replace(HTTP_LOG_MSG, "%{remote}p", port, -1)
		host, port, _ = net.SplitHostPort(r.Host)
		HTTP_LOG_MSG = strings.Replace(HTTP_LOG_MSG, "%A", host, -1)
		HTTP_LOG_MSG = strings.Replace(HTTP_LOG_MSG, "%{local}p", port, -1)
		HTTP_LOG_MSG = strings.Replace(HTTP_LOG_MSG, "%U", r.URL.Path, -1)
		HTTP_LOG_MSG = strings.Replace(HTTP_LOG_MSG, "%I", fmt.Sprint(r.ContentLength), -1)

		// Process cookies
		if strings.Contains(HTTP_LOG_MSG, "%c") {
			v := []string{}
			for _, c := range r.Cookies() {
				v = append(v, c.Name+"="+c.Value)
			}
			HTTP_LOG_MSG = strings.Replace(HTTP_LOG_MSG, "%c", strings.Join(v, "&"), -1)
		}
		re := regexp.MustCompile(`%{[^ ;,}]*}c`)
		cookies := re.FindAll([]byte(HTTP_LOG_MSG), -1)
		for _, cookie := range cookies {
			cookieName := strings.TrimPrefix(string(cookie), "%{")
			cookieName = strings.TrimSuffix(cookieName, "}c")
			c, _ := r.Cookie(cookieName)
			if c != nil {
				HTTP_LOG_MSG = strings.Replace(HTTP_LOG_MSG, string(cookie), c.Name+"="+c.Value, -1)
			} else {
				HTTP_LOG_MSG = strings.Replace(HTTP_LOG_MSG, string(cookie), c.Name+"=", -1)
			}
		}

		// Process headers
		if strings.Contains(HTTP_LOG_MSG, "%i") {
			v := []string{}
			for k := range r.Header {
				v = append(v, k+"="+r.Header.Get(k))
			}
			HTTP_LOG_MSG = strings.Replace(HTTP_LOG_MSG, "%i", strings.Join(v, "&"), -1)
		}
		re = regexp.MustCompile(`%{[^ ;,}]*}i`)
		headers := re.FindAll([]byte(HTTP_LOG_MSG), -1)
		for _, header := range headers {
			headerName := strings.TrimPrefix(string(header), "%{")
			headerName = strings.TrimSuffix(headerName, "}i")
			h := r.Header.Get(headerName)
			HTTP_LOG_MSG = strings.Replace(HTTP_LOG_MSG, string(header), headerName+"="+h, -1)
		}

		// Process GET/POST parameters
		if strings.Contains(HTTP_LOG_MSG, "%{GET}f") {
			HTTP_LOG_MSG = strings.Replace(HTTP_LOG_MSG, "%{GET}f", r.URL.RawQuery, -1)
		}
		if strings.Contains(HTTP_LOG_MSG, "%{POST}f") {
			v := []string{}
			err := r.ParseMultipartForm(32 << 20)
			if err != nil {
				r.ParseForm()
			}
			for key, val := range r.PostForm {
				v = append(v, key+"=["+strings.Join(val, ",")+"]")
			}
			HTTP_LOG_MSG = strings.Replace(HTTP_LOG_MSG, "%{POST}f", strings.Join(v, "&"), -1)
		}

		// Add context with stime
		ctx := context.WithValue(r.Context(), CKey("start"), time.Now())
		r = r.WithContext(ctx)
		res := responseWriter{
			w: w,
		}

		// Execute the actual handler
		f(&res, r)

		// add etime
		ctx = context.WithValue(r.Context(), CKey("end"), time.Now())
		r = r.WithContext(ctx)

		// Add post execution context
		// response counter
		HTTP_LOG_MSG = strings.Replace(HTTP_LOG_MSG, "%B", fmt.Sprint(res.GetCounter()), -1)

		// response code
		HTTP_LOG_MSG = strings.Replace(HTTP_LOG_MSG, "%>s", fmt.Sprint(res.GetCode()), -1)

		// time taken
		sTime := ctx.Value(CKey("start")).(time.Time)
		eTime := ctx.Value(CKey("end")).(time.Time)
		if strings.Contains(HTTP_LOG_MSG, "%D") {
			HTTP_LOG_MSG = strings.Replace(HTTP_LOG_MSG, "%D", fmt.Sprint(eTime.Sub(sTime).Nanoseconds()/1000), -1)
		}
		if strings.Contains(HTTP_LOG_MSG, "%T") {
			HTTP_LOG_MSG = strings.Replace(HTTP_LOG_MSG, "%T", fmt.Sprintf("%0.3f", float64(eTime.Sub(sTime).Nanoseconds())/1000000000), -1)
		}

		// Log Metrics
		SetMetric("uadmin/http/responsetime", float64(eTime.Sub(sTime).Nanoseconds()/1000000))
		IncrementMetric("uadmin/http/requestrate")

		go func() {
			if LogHTTPRequests {
				// Send log to syslog
				Syslogf(INFO, HTTP_LOG_MSG)
			}
		}()
	}
}

type responseWriter struct {
	w       http.ResponseWriter
	counter uint64
	code    int
}

func (w *responseWriter) Header() http.Header {
	return w.w.Header()
}

func (w *responseWriter) Write(b []byte) (int, error) {
	w.counter += uint64(len(b))
	return w.w.Write(b)
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.w.WriteHeader(statusCode)
}

func (w *responseWriter) GetCode() int {
	if w.code == 0 {
		return 200
	}
	return w.code
}

func (w *responseWriter) GetCounter() uint64 {
	return w.counter
}
