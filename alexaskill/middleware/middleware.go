package middleware

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var appID string

func AppID() {

}

func AlexaValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		certURL := r.Header.Get("SignatureCertChainUrl")

		if !verifyCertURL(certURL) {
			w.WriteHeader(404)
			w.Write([]byte("Cert URL error"))
			return
		}

		publicKey, err := verifySignatureCert(certURL)
		if err != nil {
			w.WriteHeader(404)
			return
		}

		if verifySignature(publicKey, r) {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(404)
			return
		}

	})
}

func verifySignatureCert(certURL string) (*rsa.PublicKey, error) {
	if !verifyCertURL(certURL) {
		return nil, errors.New("URL validation failed")
	}

	resp, err := http.Get(certURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	certBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	publicKey := verifyCert(certBody)

	return publicKey, nil
}

func verifySignature(publicKey *rsa.PublicKey, r *http.Request) bool {
	signature, _ := base64.StdEncoding.DecodeString(r.Header.Get("Signature"))

	requestCopy, _ := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewReader(requestCopy))

	derivedHash := getDerivedHash(requestCopy)
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA1, derivedHash, signature)

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func verifyCert(b []byte) *rsa.PublicKey {
	block, _ := pem.Decode(b)
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Println(err)
		return nil
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(cert)

	opts := x509.VerifyOptions{
		DNSName: "echo-api.amazon.com",
		Roots:   certPool,
	}

	if _, err = cert.Verify(opts); err != nil {
		log.Println(err)
		return nil
	}

	pkey := cert.PublicKey.(*rsa.PublicKey)

	return pkey
}

func verifyCertURL(path string) bool {
	link, _ := url.Parse(path)

	if link.Scheme != "https" {
		return false
	}

	if link.Host != "s3.amazonaws.com" && link.Host != "s3.amazonaws.com:443" {
		return false
	}

	if !strings.HasPrefix(link.Path, "/echo.api/") {
		return false
	}

	return true
}

func getDerivedHash(b []byte) []byte {
	h := sha1.New()
	h.Write(b)

	return h.Sum(nil)
}
