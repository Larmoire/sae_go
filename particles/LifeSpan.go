package particles

func (p *Particle) DecreaseLife() {
	//Si le Lifespan est actif, on retire 1 à la durée de vie de la particule
	if p.Lifespan > 0 {
		p.Lifespan -= 1
		p.ChangeOpacity()
	}

}
