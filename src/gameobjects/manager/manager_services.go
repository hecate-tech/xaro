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
			m.ServerPlayers[sID] = player.New(m.world)
			m.ServerPlayers[sID].IsPlaying = false

			log.Println(m.ServerPlayers[sID].Username, "has connected...")

			for _, system := range m.world.Systems() {
				switch sys := system.(type) {
				case *comm.RenderSystem:
					sys.Add(&m.ServerPlayers[sID].BasicEntity, &m.ServerPlayers[sID].RenderComponent, &m.ServerPlayers[sID].SpaceComponent)
				}
			}
		}

		m.ServerPlayers[sID].Position.Set(sP.GetPosition().X, sP.GetPosition().Y)
		m.ServerPlayers[sID].Ase.Play(sP.AnimName)
		m.ServerPlayers[sID].Ase.Update(dt)
		m.ServerPlayers[sID].Drawable = m.ServerPlayers[sID].Spritesheet.Drawable(int(m.ServerPlayers[sID].Ase.CurrentFrame))
		if strings.Contains(m.ServerPlayers[sID].Ase.CurrentAnimation.Name, "action") {
			m.ServerPlayers[sID].Ase.PlaySpeed = m.ServerPlayers[sID].ShootSpeed
		} else {
			m.ServerPlayers[sID].Ase.PlaySpeed = 1.0
		}
	}

	for ID := range m.ServerPlayers {
		if _, ok := sPlayers.Players[ID]; !ok {
			log.Println(m.ServerPlayers[ID].Username, "has disconnected...")

			for _, system := range m.world.Systems() {
				switch sys := system.(type) {
				case *comm.RenderSystem:
					sys.Remove(m.ServerPlayers[ID].BasicEntity)
				}
			}

			delete(m.ServerPlayers, ID)
		}
	}
}
