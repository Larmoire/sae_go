package particles

import (
	Extensions "project-particles/Extension"
	"project-particles/config"
)

func (p *Particle) UpdatePos() {
	//Si SpeedFix n'est pas actif, on met à jour la position de la particule par sa vitesse
	if !Extensions.SpeedFix {
		//On met à jour la position en Y
		p.PositionX += p.SpeedX
		//Si la particule est en mode gravity, on augmente sa vitesse en y pour gérer une accélération
		if Extensions.Gravity {
			p.SpeedY += config.General.GravityVal
		}
		//Si la particule est en mode bounce et gravité, on vérifie si elle touche le sol
		if Extensions.Bounce && Extensions.Gravity && (p.PositionY+p.SpeedY+10*config.General.ScaleY >= float64(config.General.WindowSizeY)) {
			//Si c'est le cas, on met leur vitesse à 0.
			//On pourrait aussi inverser leur vitesse en Y pour qu'elles rebondissent mais ça ne marche pas très bien
			p.SpeedY = 0
			p.SpeedX = 0
		}
		//On met à jour la position en Y
		p.PositionY += p.SpeedY
	}
}
