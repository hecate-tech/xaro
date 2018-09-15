package player

import (
	"fmt"
	"strings"
	"time"

	"engo.io/engo"
)

var (
	directions = []string{"left", "right", "up", "down"}
	timeStamp  int64
	timer      time.Timer
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
	if !p.IsShooting {
		p.IsShooting = (engo.Input.Mouse.Action == engo.Press && engo.Input.Mouse.Button == engo.MouseButtonLeft)
	} else {
		p.IsShooting = !(engo.Input.Mouse.Action == engo.Release && engo.Input.Mouse.Button == engo.MouseButtonLeft)
	}

	return p.IsShooting
}

func (p *Player) updateAnimation() {
	p.Ase.PlaySpeed = 1.0
	for _, dir := range directions {
		// Plays the walking animation of the current
		// direction you are currently pressing
		if engo.Input.Button(dir).Down() {
			p.Ase.Play(dir)
			break
		}
	}
}

func (p *Player) updateAction(dt float32) {
	p.Ase.PlaySpeed = p.ShootSpeed
	p.Ase.Play(getCurrentDirection(p.Ase.CurrentAnimation.Name) + "action")

	// Check if timer has run out and action if so. If not then continue.
	if timeStamp += int64(dt); timeStamp <= time.Now().UnixNano()/int64(time.Millisecond) {
		nextActionTime := int64((float32(p.Ase.CurrentAnimation.End-(p.Ase.CurrentAnimation.Start-1)) * 100) / p.ShootSpeed)
		timeStamp = time.Now().UnixNano()/int64(time.Millisecond) + nextActionTime

		p.action() // Perform player action when timer is up
	}
}

func (p *Player) action() {
	fmt.Println("SHOOT ~ >>>----------|>")
}

func (p *Player) updateIdleAnimation() {
	// if player velocity is x:0, y:0 then play idle animation
	if p.Velocity.Equal(engo.Point{}) {
		p.Ase.Play(getCurrentDirection(p.Ase.CurrentAnimation.Name) + "idle")
	}
}

func getCurrentDirection(animName string) string {
	for _, dir := range directions {
		// returns the direction that the current animation is
		if strings.Contains(animName, dir) {
			return dir
		}
	}
	return ""
}
