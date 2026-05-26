// Package gtls provides grpc secure connectivity, supporting both server-only authentication and client-server authentication.
package gtls

import (
	"google.golang.org/grpc/credentials"
)

// GetServerTLSCredentialsByCA two-way authentication via CA-issued root certificate
func GetServerTLSCredentialsByCA(caFile string, certFile string, keyFile string) (credentials.TransportCredentials, error) {
	_ = "STUB: not implemented"
	//read and parse the information from the certificate file to obtain the certificate public key, key pair
	return *new(credentials.TransportCredentials), nil
}

// create an empty CertPool

//attempts to parse the incoming PEM-encoded certificate. If the parsing is successful it will be added to the CertPool for later use

//building TLS-based TransportCredentials options

// set up a certificate chain that allows the inclusion of one or more
// requirement to verify the client's certificate
// set the set of root certificates and use the mode set in ClientAuth for verification

// GetServerTLSCredentials server-side authentication
func GetServerTLSCredentials(certFile string, keyFile string) (credentials.TransportCredentials, error) {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials), nil
}
