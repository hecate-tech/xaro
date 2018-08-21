package player

import (
	"strings"

	"engo.io/engo"
)

var (
	directions = []string{"left", "right", "up", "down"}
)

func (p *Player) updateMovement() {
	if engo.Input.Button("left").Down() {
		p.Velocity.X = -p.MoveSpeed
	}
	if engo.Input.Button("right").Down() {
		p.Velocity.X = +p.MoveSpeed
	}
	if engo.Input.Button("up").Down() {
		p.Velocity.Y = -p.MoveSpeed
	}
	if engo.Input.Button("down").Down() {
		p.Velocity.Y = +p.MoveSpeed
	}
}

func (p *Player) inAction() bool {
	return false
}

func (p *Player) updateAnimation() {
	if !p.inAction() {
		p.Ase.PlaySpeed = 1.0
		for _, dir := range directions {
			if engo.Input.Button(dir).Down() {
				p.Ase.Play(dir)
				break
			}
		}
	}
}

func (p *Player) updateIdleAnimation() {
	if p.Velocity.X == 0 && p.Velocity.Y == 0 && !strings.HasSuffix(p.Ase.CurrentAnimation.Name, "idle") && !p.inAction() {
		if strings.HasSuffix(p.Ase.CurrentAnimation.Name, "action") {
			fix := strings.TrimSuffix(p.Ase.CurrentAnimation.Name, "action")
			p.Ase.Play(fix + "idle")
		} else {
			p.Ase.Play(p.Ase.CurrentAnimation.Name + "idle")
		}
	}
}
