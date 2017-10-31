package httplib

import (
	"net"
	"net/http"
	"time"
)

var (
	httpClient *http.Client
)

const (
	CONNECTION_TIMEOUT      = 20
	RESPONSE_HEADER_TIMEOUT = 60
	HTTP_TIMEOUT            = 90
	MAX_IDLE_CONNECTIONS    = 20
)

// init HTTPClient
func init() {
	//httpClient = NewHttpClient()
}

func GetHttp() *http.Client {
	//return httpClient
	return NewHttpClient()
}

//
// Clients and Transports are safe for concurrent use by multiple goroutines
// and for efficiency should only be created once and re-used.
//
// The Client's Transport typically has internal state (cached TCP
// connections), so Clients should be reused instead of created as
// needed. Clients are safe for concurrent use by multiple goroutines.
//
func NewHttpClient() *http.Client {
	return NewTimeoutClient(CONNECTION_TIMEOUT*time.Second, RESPONSE_HEADER_TIMEOUT*time.Second)
}

func NewTimeoutClient(connectTimeout time.Duration, responseHeaderTimeout time.Duration) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			//Dial: TimeoutDialer(connectTimeout),
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 120 * time.Second,
			}).Dial,
			ResponseHeaderTimeout: responseHeaderTimeout,
			MaxIdleConnsPerHost:   MAX_IDLE_CONNECTIONS, // NEW
			//	TLSHandshakeTimeout: 10 * time.Second,
		},
		Timeout: HTTP_TIMEOUT * time.Second,
	}
}

func TimeoutDialer(connectTimeout time.Duration) func(net, addr string) (net.Conn, error) {
	return func(network, address string) (net.Conn, error) {
		conn, err := net.DialTimeout(network, address, connectTimeout)
		return conn, err
	}
}

/* exmaple:
https://groups.google.com/forum/#!topic/golang-nuts/kK74jkgfnEQ

var client = &http.Client{
  Transport: &http.Transport{
    Dial: func(network, addr string) (net.Conn, error) {
      log.Println("dial!")
      return net.Dial(network, addr)
    },
    MaxIdleConnsPerHost: 50,
  },
}
*/
