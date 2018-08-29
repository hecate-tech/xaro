package manager

import (
	"log"
	"strings"

	comm "engo.io/engo/common"
	"github.com/damienfamed75/engo-xaro/src/gameobjects/player"
	pb "github.com/damienfamed75/engo-xaro/src/proto"
)

func (m *Manager) updateConnectedPlayers(sPlayers *pb.Players, dt float32) {
	for sID, sP := range sPlayers.Players {
		if _, ok := m.ServerPlayers[sID]; !ok {
			m.newServerPlayer(sID)
			log.Println(sP.Username, "has connected...")
		}
		m.updateServerPlayer(sID, sP)
	}

	for ID := range m.ServerPlayers {
		if _, ok := sPlayers.Players[ID]; !ok {
			log.Println(m.ServerPlayers[ID].Username, "has disconnected...")
			m.removeServerPlayer(ID)
		}
	}
}

func (m *Manager) updateServerPlayer(index uint32, sp *pb.Player) {
	m.ServerPlayers[index].Ase.Play(sp.AnimName)
	m.ServerPlayers[index].Position.Set(sp.GetPosition().X, sp.GetPosition().Y)
	m.ServerPlayers[index].Drawable = m.ServerPlayers[index].Spritesheet.Drawable(int(m.ServerPlayers[index].Ase.CurrentFrame))

	if strings.Contains(m.ServerPlayers[index].Ase.CurrentAnimation.Name, "action") {
		m.ServerPlayers[index].Ase.PlaySpeed = m.ServerPlayers[index].ShootSpeed
	} else {
		m.ServerPlayers[index].Ase.PlaySpeed = 1.0
	}
}

func (m *Manager) newServerPlayer(index uint32) {
	m.ServerPlayers[index] = player.New(m.world)
	m.ServerPlayers[index].IsPlaying = false

	for _, system := range m.world.Systems() {
		switch sys := system.(type) {
		case *comm.RenderSystem:
			sys.Add(&m.ServerPlayers[index].BasicEntity, &m.ServerPlayers[index].RenderComponent, &m.ServerPlayers[index].SpaceComponent)
		}
	}
}

func (m *Manager) removeServerPlayer(index uint32) {
	for _, system := range m.world.Systems() {
		switch sys := system.(type) {
		case *comm.RenderSystem:
			sys.Remove(m.ServerPlayers[index].BasicEntity)
		}
	}
	delete(m.ServerPlayers, index)
}
