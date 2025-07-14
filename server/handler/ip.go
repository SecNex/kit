package handler

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

type IPResponse struct {
	IP   string `json:"ip"`
	Type string `json:"type"`
}

// Handler to get the IP address of the client (IPv4 and IPv6)
func (h *Handler) IP(w http.ResponseWriter, r *http.Request) {
	// Try to get real IP from headers first (for reverse proxies)
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.Header.Get("X-Real-IP")
	}

	// If no headers, use RemoteAddr
	if ip == "" {
		ip = r.RemoteAddr
	}

	if ip == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("No IP address found"))
		return
	}

	// Use net.SplitHostPort to properly separate IP from port
	// This works correctly for both IPv4 and IPv6
	host, _, err := net.SplitHostPort(ip)
	if err != nil {
		// If SplitHostPort fails, it might be just an IP without port
		host = ip
	}

	ipType := "ipv4"
	if strings.Contains(host, ":") {
		ipType = "ipv6"
	}

	if r.Header.Get("Accept") == "application/json" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{"ip":"%s", "type":"%s"}`, host, ipType)))
	} else {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(host))
	}
}
