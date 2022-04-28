package server

import "net/http"

type server struct {
	addr string
	certFile string
	keyFile string 
}

func New(addr, certFile, keyFile string) *server {
	return &server {
		addr: addr, 
		certFile: certFile, 
		keyFile: keyFile,
	}
}

func (s *server) Start() error {
	err := http.ListenAndServeTLS(s.addr, s.certFile, s.keyFile, nil)
	if err != nil {
		return err
	}

	return nil
}