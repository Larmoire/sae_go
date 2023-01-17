package particles

import (
	"math/rand"
	Extensions "project-particles/Extension"
	"project-particles/config"
)

//Fonction qui fait rebondir la particule sur les bords de la fenêtre
func (p *Particle) Bounce() {
	//On observe si la particule touche les bords de la fenêtre en X
	if p.PositionX <= 0 || p.PositionX+p.SpeedX+10*config.General.ScaleX >= float64(config.General.WindowSizeX) {
		//Si c'est le cas, on inverse la vitesse en X
		p.SpeedX = -p.SpeedX
		//Si l'option ColorBounce est activée, on change la couleur de la particule aussi avec des valeurs aléatoires
		if Extensions.ColorBounce {
			p.ColorRed, p.ColorGreen, p.ColorBlue = rand.Float64(), rand.Float64(), rand.Float64()
		}
	}
	//On observe maintenant si la particule touche les bords de la fenêtre en Y
	if p.PositionY <= 0 || p.PositionY+p.SpeedY+10*config.General.ScaleY >= float64(config.General.WindowSizeY) {
		//Si c'est le cas, on inverse maintenant la vitesse en Y
		p.SpeedY = -p.SpeedY
		//Si l'option ColorBounce est activée, on change la couleur de la particule aussi avec des valeurs aléatoires
		if Extensions.ColorBounce {
			p.ColorRed, p.ColorGreen, p.ColorBlue = rand.Float64(), rand.Float64(), rand.Float64()
		}
	}
}
