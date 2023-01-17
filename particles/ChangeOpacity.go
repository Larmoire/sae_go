package particles

import "project-particles/config"

func (p *Particle) ChangeOpacity() {
	//Si l'opacité de base est à 1, on diminue l'opacité de la particule
	if p.Opacity != 0 {
		if config.General.Opacity == 1 {
			p.Opacity -= 1 / config.General.Lifespan
			//Sinon, on l'augmente
		} else {
			p.Opacity += 1 / config.General.Lifespan
		}
	}
}
