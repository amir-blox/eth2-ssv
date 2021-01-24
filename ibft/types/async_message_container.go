package types

import "sync"

type MessagesContainer struct {
	messages map[uint64]map[uint64]SignedMessage
	lock     sync.Mutex
}

func NewMessagesContainer() *MessagesContainer {
	return &MessagesContainer{
		messages: make(map[uint64]map[uint64]SignedMessage),
		lock:     sync.Mutex{},
	}
}

func (c *MessagesContainer) ReadOnlyMessagesByRound(round uint64) map[uint64]SignedMessage {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.messages[round]
}

func (c *MessagesContainer) AddMessage(msg SignedMessage) {
	c.lock.Lock()
	defer c.lock.Unlock()

	roundMsgs, found := c.messages[msg.Message.Round]
	if !found {
		roundMsgs = make(map[uint64]SignedMessage)
	}
	roundMsgs[msg.IbftId] = msg
	c.messages[msg.Message.Round] = roundMsgs
}
