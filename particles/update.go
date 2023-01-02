package particles

import (
	"math/rand"
	"project-particles/config"
	"time"
)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.

var spawnrate float64 = config.General.SpawnRate //On peut tester avec 0.17 pour avoir environ 1 particule par seconde

func (s *System) Update() {

	rand.Seed(time.Now().UnixNano())

	e := s.Content.Front()

	for e != nil {

		upPosition(e.Value.(*Particle))

		decreaseLife(e.Value.(*Particle))

		if outwindows(e.Value.(*Particle)) {
			s.Content.Remove(e)
		}
		e = e.Next()
	}

	for spawnrate >= 1 {
		s.Content.PushFront(createParticule())
		spawnrate--
	}

	if spawnrate < 1 {
		spawnrateadd()
	}

}

func spawnrateadd() {
	spawnrate += config.General.SpawnRate
}

func outwindows(p *Particle) bool {
	if p.PositionX > float64(config.General.WindowSizeX) || p.PositionX < 0 || p.PositionY > float64(config.General.WindowSizeY) || p.PositionY < 0 {
		return true
	}
	return false
}

func upPosition(p *Particle) {
	p.PositionX += p.SpeedX
	p.PositionY += p.SpeedY
	if config.General.Gravity {
		p.PositionY += config.General.GravityVal
	}

}

func decreaseLife(p *Particle) {
	p.Lifespan -= 1
	decreaseOpacity(p)
}

func decreaseOpacity(p *Particle) {
	p.Opacity -= 1 / config.General.Lifespan
}
func increaseOpacity(p *Particle) {
	p.Opacity += 0.005
}
