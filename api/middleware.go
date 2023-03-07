package api

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
)

type userIPType string

var userIP userIPType = "user_ip"

func (s *server) LogMiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware invoked : before handler")
		defer log.Println("Middleinvoked : after handler")
		h.ServeHTTP(w, r)
	})

}

func (s *server) IPExtractMiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("IP Middleware invoked ")

		var uIP string
		uIP, err := getIP(r)
		if err != nil {
			uIP = "unknown"

		}
		ctx := context.WithValue(context.Background(), userIP, uIP)

		h.ServeHTTP(w, r.WithContext(ctx))
	})

}

func getIP(r *http.Request) (string, error) {
	// Check if its forwarded request
	forwardIP := r.Header.Get("X-Forwarded-For")

	if len(forwardIP) != 0 {
		return forwardIP, nil
	}

	// extract the IP from request
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "unknown", fmt.Errorf("Ip address not found ")
	}
	srcIP := net.ParseIP(ip)
	if len(srcIP) == 0 {
		return "unknown", nil
	}
	return srcIP.String(), nil

}

func getIPFromContext(ctx context.Context) string {
	return ctx.Value(userIP).(string)
}
