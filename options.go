package darktrace

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
)

// Disable TLS validation for endpoints without valid SSL certificates
func ClientDisableTLSValidation(c *Client) error {
	c.t.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	return nil
}

// Add root ca to current request store
func ClientAddRootCA(localCertFile string) func(c *Client) error {
	return func(c *Client) error {
		// Get the current client CAs
		rootCAs := c.t.TLSClientConfig.RootCAs

		// If client doesn't have CAs, get the SystemCertPool
		if rootCAs == nil {
			rootCAs, _ = x509.SystemCertPool()
		}

		// Continue with an empty pool if not able to get SystemCertPool
		if rootCAs == nil {
			rootCAs = x509.NewCertPool()
		}

		// Read in the cert file
		certs, err := ioutil.ReadFile(localCertFile)
		if err != nil {
			return fmt.Errorf("failed to append %q to RootCAs: %v", localCertFile, err)
		}

		// Append our cert to the system pool
		if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
			return fmt.Errorf("no certs appended, using system certs only")
		}

		// Trust the augmented cert pool in our client
		c.t.TLSClientConfig.RootCAs = rootCAs
		return nil
	}
}

func Param(key, value string) func() (string, string) {
	return func() (string, string) {
		return key, value
	}
}
