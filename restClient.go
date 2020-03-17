package bcoin

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
)

type restClient struct {
	serverAddr string
	user       string
	passwd     string
	httpClient *http.Client
	timeout    int
}

type HttpVerb int

const (
	Connect HttpVerb = iota
	Delete
	Get
	Head
	Options
	Path
	Post
	Put
	Trace
)

func (v HttpVerb) String() string {
	return [...]string{"Connect", "Delete", "Get", "Head", "Options", "Path", "Post", "Put", "Trace"}[v]
}

func newClient(host string, port int, user, passwd string, useSSL bool, timeout int) (c *restClient, err error) {
	if len(host) == 0 {
		err = errors.New("missing argument host")
	}

	var serverAddr string
	var httpClient *http.Client

	if useSSL {
		serverAddr = "https://"
		t := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		httpClient = &http.Client{Transport: t}
	} else {
		serverAddr = "http://"
		httpClient = &http.Client{}
	}

	c = &restClient{serverAddr: fmt.Sprintf("%s%s:%d", serverAddr, host, port), user: user, passwd: passwd, httpClient: httpClient, timeout: timeout}
	return
}

func (c *restClient) call(method HttpVerb, params interface{}) (r, err error) {

}