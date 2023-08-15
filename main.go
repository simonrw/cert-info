package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	var (
		hostname     = flag.String("hostname", "", "Hostname to connect to")
		jsonOutput   = flag.Bool("json", false, "Output JSON")
		noServerName = flag.Bool("noservername", false, "Do not set server-name in TLS configuration")
	)
	flag.Parse()

	if *hostname == "" {
		log.Fatal("no hostname specified")
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:443", *hostname))
	if err != nil {
		log.Fatalf("cannot connect to %s: %v", *hostname, err)
	}

	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	if !*noServerName {
		config.ServerName = *hostname
	}
	client := tls.Client(conn, config)

	if err := client.Handshake(); err != nil {
		log.Fatalf("cannot perform handshake with %s: %v", *hostname, err)
	}

	s := client.ConnectionState()
	certs := s.PeerCertificates

	var cert *x509.Certificate
	if len(certs) == 0 {
		log.Fatalf("no peer certificates found")
	} else {
		cert = certs[0]
	}

	if *jsonOutput {
		renderJson(cert, os.Stdout)
	} else {
		renderPretty(cert, os.Stdout)
	}
}

func renderJson(cert *x509.Certificate, writer io.Writer) {
	b, err := json.Marshal(cert)
	if err != nil {
		log.Fatalf("cannot output json: %v", err)
	}
	fmt.Printf("%s\n", string(b))

}

func renderPretty(cert *x509.Certificate, writer io.Writer) {

	for _, name := range cert.DNSNames {
		fmt.Printf("SAN: %s\n", name)
	}
	for _, email := range cert.EmailAddresses {
		fmt.Printf("email: %s\n", email)
	}
	for _, ip := range cert.IPAddresses {
		fmt.Printf("ip: %s\n", ip)
	}
	for _, uri := range cert.URIs {
		fmt.Printf("uri: %s\n", uri)
	}
	fmt.Printf("valid from %s to %s\n", cert.NotBefore, cert.NotAfter)
	fmt.Printf("issuer: %s\n", cert.Issuer)
	fmt.Printf("is ca: %v\n", cert.IsCA)
}
