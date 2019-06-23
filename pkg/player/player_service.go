package player

import (

	// "github.com/EngoEngine/engo"
	"github.com/EngoEngine/ecs"
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
