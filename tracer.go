package bitswap

import (
	peer "github.com/libp2p/go-libp2p-core/peer"
	bsmsg "github.com/uss2022sayahi/go-bitswap/message"
)

// Tracer provides methods to access all messages sent and received by Bitswap.
// This interface can be used to implement various statistics (this is original intent).
type Tracer interface {
	MessageReceived(peer.ID, bsmsg.BitSwapMessage)
	MessageSent(peer.ID, bsmsg.BitSwapMessage)
}

// Configures Bitswap to use given tracer.
func WithTracer(tap Tracer) Option {
	return func(bs *Bitswap) {
		bs.tracer = tap
	}
}
