package particles

import (
	"math"
	Extensions "project-particles/Extension"
	"project-particles/config"
)

func (p *Particle) UpdateOrbit() {
	distance := math.Sqrt(math.Pow(p.PositionX-float64(config.General.WindowSizeX)/2, 2) + math.Pow(p.PositionY-float64(config.General.WindowSizeY)/2, 2))

	// Calcule l'angle actuel de la particule par rapport au centre de l'orbite
	currentAngle := math.Atan2(p.PositionY-float64(config.General.WindowSizeY)/2, p.PositionX-float64(config.General.WindowSizeX)/2)

	// Met à jour l'angle de la particule en fonction de la vitesse d'orbite
	newAngle := currentAngle + math.Pi - 1

	// Calcule les nouvelles coordonnées de la particule en utilisant la distance et l'angle mis à jour
	if !Extensions.SpeedFix {
		p.PositionX = float64(config.General.WindowSizeX)/2 + distance*math.Cos(newAngle)
		p.PositionY = float64(config.General.WindowSizeY)/2 + distance*math.Sin(newAngle)
		p.Rotation -= 0.05
	}
}
