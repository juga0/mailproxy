// event.go - Katzenpost client mailproxy API events.
// Copyright (C) 2018  Yawning Angel.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Package event implements the event types returned by the API's event
// listener.
package event

import (
	"encoding/hex"
	"fmt"

	"github.com/katzenpost/core/crypto/ecdh"
)

// Event is the generic event sent over the event listener channel.
type Event interface {
	// String returns a string representation of the Event.
	String() string
}

// ConnectionStatusEvent is the event sent when an account's connection status
// changes.
type ConnectionStatusEvent struct {
	// AccountID is the account identifier for the account associated with
	// the event.
	AccountID string

	// IsConnected is true iff the account is connected to the provider.
	IsConnected bool
}

// String returns a string representation of the ConnectionStatusEvent.
func (e *ConnectionStatusEvent) String() string {
	return fmt.Sprintf("ConnectionStatus[%v]: %v", e.AccountID, e.IsConnected)
}

// MessageReceivedEvent is the event sent when a new message is received.
type MessageReceivedEvent struct {
	// AccountID is the account identifier for the account associated with
	// the event.
	AccountID string

	// SenderKey is the message sender's public key, if any.
	SenderKey *ecdh.PublicKey

	// MessageID is the local unique identifier for the message.
	MessageID []byte
}

// String returns a string representation of the MessageReceivedEvent.
func (e *MessageReceivedEvent) String() string {
	return fmt.Sprintf("MessageReceived[%v]: %v %v", e.AccountID, e.SenderKey, hex.EncodeToString(e.MessageID))
}

// KaetzchenReplyEvent is the event sent when a Kaetzchen request completes.
type KaetzchenReplyEvent struct {
	// AccountID is the account identifier for the account associated with the
	// event.
	AccountID string

	// MessageID is the unique identifier for the request associated with the
	// reply.
	MessageID []byte

	// Payload is the reply payload if any.
	Payload []byte

	// Err is the error encountered when servicing the request if any.
	Err error
}

// String returns a string representation of the KaetzchenReplyEvent.
func (e *KaetzchenReplyEvent) String() string {
	if e.Err != nil {
		return fmt.Sprintf("KaetzchenReply[%v]: %v failed: %v", e.AccountID, hex.EncodeToString(e.MessageID), e.Err)
	}
	return fmt.Sprintf("KaetzchenReply[%v]: %v (%v bytes)", e.AccountID, hex.EncodeToString(e.MessageID), len(e.Payload))
}
