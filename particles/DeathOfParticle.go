package particles

import (
	"container/list"
	"project-particles/config"
)

func (p *Particle) DeathOfParticle(s *System, e *list.Element) {
	if p.Lifespan == 0 || Outwindows(p) {
		p.Opacity = 0
		if config.General.Optimisation {
			go s.Content.Remove(e)
			Countdead += 1
		}
	}
}
func Outwindows(p *Particle) bool {
	if p.PositionX > float64(config.General.WindowSizeX) || p.PositionX < 0-10*config.General.ScaleX || p.PositionY > float64(config.General.WindowSizeY) || p.PositionY < 0-10*config.General.ScaleY {
		return true
	}
	return false
}
