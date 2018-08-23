package utils

import "sync/atomic"

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
func (manager *IDManager) NextPlayerID() uint32 {
	return manager.nextID(manager.playerIDs)
}

// NextBotID returns a new unique bot ID
func (manager *IDManager) NextBotID() uint32 {
	return manager.nextID(manager.botsIDs)
}

// IsPlayerID returns if id exists already
func (manager *IDManager) IsPlayerID(id uint32) bool {
	_, ok := manager.playerIDs[id]
	return ok
}

// IsBotID returns if id exists already
func (manager *IDManager) IsBotID(id uint32) bool {
	_, ok := manager.botsIDs[id]
	return ok
}

func (manager *IDManager) nextID(idsMap map[uint32]bool) uint32 {
	id := atomic.AddUint32(&manager.lastID, 1)
	idsMap[id] = true
	return id
}
