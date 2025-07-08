package middlewares

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/secnex/kit/models"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
	size       int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(b)
	rw.size += size
	return size, err
}

func getClientIP(r *http.Request) string {
	// Zuerst X-Forwarded-For prüfen (bei Proxies/Load Balancern)
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	if xForwardedFor != "" {
		// X-Forwarded-For kann mehrere IPs enthalten, erste ist Client-IP
		ips := strings.Split(xForwardedFor, ",")
		clientIP := strings.TrimSpace(ips[0])
		if clientIP != "" && isValidIP(clientIP) {
			return clientIP
		}
	}

	// X-Real-IP prüfen (oft von nginx verwendet)
	xRealIP := r.Header.Get("X-Real-IP")
	if xRealIP != "" && isValidIP(xRealIP) {
		return xRealIP
	}

	// CF-Connecting-IP prüfen (Cloudflare)
	cfConnectingIP := r.Header.Get("CF-Connecting-IP")
	if cfConnectingIP != "" && isValidIP(cfConnectingIP) {
		return cfConnectingIP
	}

	// RemoteAddr als Fallback (direkte Verbindung)
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

func isValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

func formatNginxTime(t time.Time) string {
	return t.Format("[02/Jan/2006:15:04:05 -0700]")
}

func (m *Middleware) LogToFile(message string) {
	file, err := os.OpenFile(m.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(message + "\n")
	if err != nil {
		log.Printf("Failed to write to log file: %v", err)
	}
}

func (m *Middleware) LogToDatabase(ipAddress string, method string, path string, proto string, duration time.Duration, userAgent string, status int) {
	m.DB.DB.Create(&models.HTTPLog{
		IPAddress: ipAddress,
		Method:    method,
		Path:      path,
		Proto:     proto,
		Duration:  duration,
		UserAgent: userAgent,
		Status:    status,
	})
}

func (m *Middleware) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		clientIP := getClientIP(r)

		wrapper := &responseWriter{
			ResponseWriter: w,
			statusCode:     200, // Default status code
			size:           0,
		}

		next.ServeHTTP(wrapper, r)
		duration := time.Since(start)

		userAgent := r.Header.Get("User-Agent")
		if userAgent == "" {
			userAgent = "-"
		}

		referer := r.Header.Get("Referer")
		if referer == "" {
			referer = "-"
		}

		// NGINX Combined Log Format
		// $remote_addr - $remote_user [$time_local] "$request" $status $body_bytes_sent "$http_referer" "$http_user_agent"
		nginxLogEntry := fmt.Sprintf(`%s - - %s "%s %s %s" %d %d "%s" "%s"`,
			clientIP,               // $remote_addr
			formatNginxTime(start), // $time_local
			r.Method,               // $request method
			r.RequestURI,           // $request URI
			r.Proto,                // $request protocol
			wrapper.statusCode,     // $status
			wrapper.size,           // $body_bytes_sent
			referer,                // $http_referer
			userAgent,              // $http_user_agent
		)

		fmt.Printf("🔎 %s %s %s %s - %s\n", r.Method, r.RequestURI, r.Proto, duration, clientIP)

		go m.LogToFile(nginxLogEntry)
		go m.LogToDatabase(clientIP, r.Method, r.RequestURI, r.Proto, duration, userAgent, wrapper.statusCode)
	})
}
