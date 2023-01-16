package particles

import "project-particles/config"

func (p *Particle) DecreaseLife() {
	p.Lifespan -= 1
	//On adapte l'Opacité à la durée de vie
	p.ChangeOpacity()
	if p.Lifespan == 0 {
		p.Opacity = 0
	}
}
func (p *Particle) ChangeOpacity() {
	if config.General.Opacity == 1 {
		p.Opacity -= 1 / config.General.Lifespan
	} else {
		p.Opacity += 1 / config.General.Lifespan
	}
}
