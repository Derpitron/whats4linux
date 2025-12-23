package mstore

import (
	"sync"
)

type MessageStore struct {
	mu     sync.RWMutex
	msgMap map[string][]string
}

func NewMessageStore() *MessageStore {
	return &MessageStore{
		msgMap: make(map[string][]string),
	}
}

func (ms *MessageStore) AddMessage(jid string, msgID string) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.msgMap[jid] = append(ms.msgMap[jid], msgID)
}
