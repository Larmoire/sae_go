package particles

import (
	"math"
	"project-particles/config"
	"testing"
)

func TestUpdateOrbit(t *testing.T) {
	//On initialise la particule
	p := Particle{PositionX: 0, PositionY: 0, Rotation: 0}
	//On calcul la nouvelle position de la particule
	distance := math.Sqrt(math.Pow(p.PositionX-float64(config.General.WindowSizeX)/2, 2) + math.Pow(p.PositionY-float64(config.General.WindowSizeY)/2, 2))
	currentAngle := math.Atan2(p.PositionY-float64(config.General.WindowSizeY)/2, p.PositionX-float64(config.General.WindowSizeX)/2)
	newAngle := currentAngle + math.Pi - 1
	//On fait bouger la particule
	p.UpdateOrbit()
	//On vérifie que la particule a bien bougé
	PositionX := float64(config.General.WindowSizeX)/2 + distance*math.Cos(newAngle)
	PositionY := float64(config.General.WindowSizeY)/2 + distance*math.Sin(newAngle)
	if p.Rotation != -0.05 || p.PositionX != PositionX || p.PositionY != PositionY {
		t.Errorf("UpdateOrbit() failed")
	}
}

func TestUpdateOrbit2(t *testing.T) {
	//On initialise la particule
	p := Particle{PositionX: 500, PositionY: 300, Rotation: 1}
	//On calcul la nouvelle position de la particule
	distance := math.Sqrt(math.Pow(p.PositionX-float64(config.General.WindowSizeX)/2, 2) + math.Pow(p.PositionY-float64(config.General.WindowSizeY)/2, 2))
	currentAngle := math.Atan2(p.PositionY-float64(config.General.WindowSizeY)/2, p.PositionX-float64(config.General.WindowSizeX)/2)
	newAngle := currentAngle + math.Pi - 1
	PositionX := float64(config.General.WindowSizeX)/2 + distance*math.Cos(newAngle)
	PositionY := float64(config.General.WindowSizeY)/2 + distance*math.Sin(newAngle)
	//On fait bouger la particule
	p.UpdateOrbit()
	//On vérifie que la particule a bien bougé
	if p.Rotation != 0.95 || p.PositionX != PositionX || p.PositionY != PositionY {
		t.Errorf("UpdateOrbit() failed")
	}
}

func TestUpdateOrbit3(t *testing.T) {
	//On initialise la particule en dehors de l'écran
	p := Particle{PositionX: 1000, PositionY: 1000, Rotation: 0.5}
	//On calcul les nouvelles coordonnées de la particule
	distance := math.Sqrt(math.Pow(p.PositionX-float64(config.General.WindowSizeX)/2, 2) + math.Pow(p.PositionY-float64(config.General.WindowSizeY)/2, 2))
	currentAngle := math.Atan2(p.PositionY-float64(config.General.WindowSizeY)/2, p.PositionX-float64(config.General.WindowSizeX)/2)
	newAngle := currentAngle + math.Pi - 1
	PositionX := float64(config.General.WindowSizeX)/2 + distance*math.Cos(newAngle)
	PositionY := float64(config.General.WindowSizeY)/2 + distance*math.Sin(newAngle)
	//On fait bouger la particule
	p.UpdateOrbit()
	//On vérifie que la particule a bien bougé au bon endroit
	if p.Rotation != 0.45 || p.PositionX != PositionX || p.PositionY != PositionY {
		t.Errorf("UpdateOrbit() failed")
	}
}
