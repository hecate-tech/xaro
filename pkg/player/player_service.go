package player

import (
	"github.com/hecatetech/xaro/general"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
)

// Update gets called every frame
func (p *Player) Update(dt float32) {
	// keep playing the animation accordingly
	// p.Ase.Update(dt)

	p.Velocity.Set(0, 0)

	p.updateMovement()
	// if p.inAction() {
	// 	p.updateAction(dt)
	// } else {
	// 	p.updateAnimation()
	// 	p.updateIdleAnimation()
	// }

	p.Position.Add(*p.Velocity.MultiplyScalar(dt))

	// p.Drawable = p.Spritesheet.Drawable(int(p.Ase.CurrentFrame))
}

// Remove deletes the player and player systems
func (p *Player) Remove(ecs.BasicEntity) {
	p.Drawable.Close()
}

func (p *Player) updateMovement() {
	if engo.Input.Button(general.BtnLeft).Down() {
		p.Velocity.X = -p.MoveSpeed
	}
	if engo.Input.Button(general.BtnRight).Down() {
		p.Velocity.X = +p.MoveSpeed
	}
	if engo.Input.Button(general.BtnUp).Down() {
		p.Velocity.Y = -p.MoveSpeed
	}
	if engo.Input.Button(general.BtnDown).Down() {
		p.Velocity.Y = +p.MoveSpeed
	}
}
