package ntransport

import (
	"net"
	"time"

	"github.com/nknorg/nkn-sdk-go"
)

// NKNTransport is the transport layer based on NKN client
type NKNTransport struct {
	m *nkn.MultiClient
}

// NewNKNTransport creates a new NKN transport layer
func NewNKNTransport(m *nkn.MultiClient) *NKNTransport {
	return &NKNTransport{m: m}
}

// Dial connects to the remote address on the network "nkn"
func (nt *NKNTransport) Dial(addr string, dialTimeout time.Duration) (net.Conn, error) {
	return nt.m.DialWithConfig(addr, &nkn.DialConfig{DialTimeout: int32(dialTimeout / time.Millisecond)})
}

func (nt *NKNTransport) Listen(port uint16) (net.Listener, error) {
	err := nt.m.Listen(nil)
	if err != nil {
		return nil, err
	}
	return nt.m, nil
}

func (nt *NKNTransport) GetNetwork() string {
	return "nkn"
}

func (nt *NKNTransport) String() string {
	return "nkn"
}
