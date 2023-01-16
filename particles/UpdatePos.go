package particles

import (
	Extensions "project-particles/Extension"
	"project-particles/config"
)

func (p *Particle) UpdatePos() {
	if !Extensions.SpeedFix {
		p.PositionX += p.SpeedX
		//Si la particule est en mode gravity, on augmente sa vitesse en y
		if Extensions.Gravity {
			p.SpeedY += config.General.GravityVal
		}
		if Extensions.Bounce && Extensions.Gravity && (p.PositionY+p.SpeedY+10*config.General.ScaleY >= float64(config.General.WindowSizeY)) {
			p.SpeedY = 0
			p.SpeedX = 0
		}
		p.PositionY += p.SpeedY
	}
}
