package network

import (
	"github.com/bloxapp/ssv/ibft/proto"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/libp2p/go-libp2p-core/peer"
	"io"
)

// SyncChanObj is a wrapper object for streaming of sync messages
type SyncChanObj struct {
	Msg    *SyncMessage
	Stream SyncStream
}

// SyncStream is a interface for all stream related functions for the sync process.
type SyncStream interface {
	io.Reader
	io.Writer
	io.Closer

	// CloseWrite closes the stream for writing but leaves it open for
	// reading.
	//
	// CloseWrite does not free the stream, users must still call Close or
	// Reset.
	CloseWrite() error

	// RemotePeer returns a string identifier of the remote peer connected to this stream
	RemotePeer() string
}

// Network represents the behavior of the network
type Network interface {
	// Broadcast propagates a signed message to all peers
	Broadcast(msg *proto.SignedMessage) error

	// ReceivedMsgChan is a channel that forwards new propagated messages to a subscriber
	ReceivedMsgChan() <-chan *proto.SignedMessage

	// BroadcastSignature broadcasts the given signature for the given lambda
	BroadcastSignature(msg *proto.SignedMessage) error

	// ReceivedSignatureChan returns the channel with signatures
	ReceivedSignatureChan() <-chan *proto.SignedMessage

	// BroadcastDecided broadcasts a decided instance with collected signatures
	BroadcastDecided(msg *proto.SignedMessage) error

	// ReceivedDecidedChan returns the channel for decided messages
	ReceivedDecidedChan() <-chan *proto.SignedMessage

	// GetHighestDecidedInstance sends a highest decided request to peers and returns answers.
	// If peer list is nil, broadcasts to all.
	GetHighestDecidedInstance(peers []peer.ID, msg *SyncMessage) (*Message, error)

	// RespondToHighestDecidedInstance responds to a GetHighestDecidedInstance
	RespondToHighestDecidedInstance(stream SyncStream, msg *SyncMessage) error

	// ReceivedSyncMsgChan returns the channel for sync messages
	ReceivedSyncMsgChan() <-chan *SyncChanObj

	// SubscribeToValidatorNetwork subscribing and listen to validator network
	SubscribeToValidatorNetwork(validatorPk *bls.PublicKey) error
}
