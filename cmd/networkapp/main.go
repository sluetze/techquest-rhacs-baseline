package main

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	// Sleep for 70 minutes before doing anything
	time.Sleep(70 * time.Minute)

	transport := &http.Transport{
		// InsecureSkipVerify disables certificate checks
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// Override dial to ensure IPv4 ("tcp4")
	transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		return (&net.Dialer{}).DialContext(ctx, "tcp4", addr)
	}
	client := &http.Client{Transport: transport}

	for {
		resp, err := client.Get("https://wirt.pleasedonotuse.win:443")
		if err != nil {
			log.Printf("http connection error: %v", err)
		} else {
			log.Printf("http-connection-established, status: %s", resp.Status)
			resp.Body.Close()
		}
		time.Sleep(5 * time.Minute)
	}
}
