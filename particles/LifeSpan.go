package particles

import "project-particles/config"

func (p *Particle) DecreaseLife() {
	//Si le Lifespan est actif, on retire 1 à la durée de vie de la particule
	if p.Lifespan > 0 {
		p.Lifespan -= 1
		p.ChangeOpacity()
	}

}

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
