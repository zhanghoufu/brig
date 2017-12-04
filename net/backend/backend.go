package backend

import (
	stdnet "net"
	"time"

	"github.com/disorganizer/brig/net/peer"
)

// Pinger
type Pinger interface {
	// LastSeen returns a timestamp of when this peer last responded.
	LastSeen() time.Time

	// Roundtrip returns the time needed to send a small package to a peer.
	Roundtrip() time.Duration

	// Err returns a non-nil value if the last try to contact this peer failed.
	Err() error

	// Close shuts down this pinger.
	Close() error
}

// Backend defines all required methods needed from the underyling implementation
// in order to talk with other nodes.
type Backend interface {
	// ResolveName resolves a human readable `name` to a list of peers.
	// Each of these can be contacted to check their credentials.
	// If the backend support exact lookups, this method will only
	// return one peer on success always.
	ResolveName(name peer.Name) ([]peer.Info, error)

	// Identity resolves our own name to an addr that we could pass to Dial.
	// It is used as part of the brig identifier for others.
	Identity() (peer.Info, error)

	// Dial builds up a connection to another peer.
	// If only ever one protocol is used, just pass the same string always.
	Dial(peerAddr, protocol string) (stdnet.Conn, error)

	// Listen returns a listener, that will yield incoming connections
	// from other peers when calling Accept.
	Listen(protocol string) (stdnet.Listener, error)

	// Ping returns a Pinger interface for the peer at `peerAddr`.
	// It should not create a full
	Ping(peerAddr string) (Pinger, error)
}
