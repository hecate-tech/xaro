package utils

import (
	"sync/atomic"
)

// IDManager is used to manage unique ID's of players and bots
type IDManager struct {
	playerIDs map[uint32]bool
	botsIDs   map[uint32]bool

	lastID uint32
}

// NewIDManager is a basic constructor for an ID manager
func NewIDManager() *IDManager {
	return &IDManager{
		playerIDs: make(map[uint32]bool),
		botsIDs:   make(map[uint32]bool),
		lastID:    0,
	}
}

// NextPlayerID returns a new unique player ID
func (m *IDManager) NextPlayerID() uint32 {
	return m.nextID(m.playerIDs)
}

// NextBotID returns a new unique bot ID
func (m *IDManager) NextBotID() uint32 {
	return m.nextID(m.botsIDs)
}

// IsPlayerID returns if id exists already
func (m *IDManager) IsPlayerID(id uint32) bool {
	_, ok := m.playerIDs[id]
	return ok
}

// IsBotID returns if id exists already
func (m *IDManager) IsBotID(id uint32) bool {
	_, ok := m.botsIDs[id]
	return ok
}

func (m *IDManager) nextID(idsMap map[uint32]bool) uint32 {
	id := atomic.AddUint32(&m.lastID, 1)
	idsMap[id] = true
	return id
}
