package particles

import (
	"math/rand"
	"project-particles/config"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.

var spawnrate float64 = config.General.SpawnRate //On peut tester avec 0.17 pour avoir environ 1 particule par seconde
var col float64

func (s *System) Update() {
	X = s.Content.Len()
	// Si toutes les particules sont mortes, vider la liste
	countdead := 0
	for e := s.Content.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Particle)
		if dead(p) {
			countdead += 1
		}
	}
	if countdead == s.Content.Len() {
		s.Content.Init()
	}
	rand.Seed(time.Now().UnixNano())
	e := s.Content.Front()

	for e != nil {

		upPosition(e.Value.(*Particle))

		if e.Value.(*Particle).Lifespan != -1 {
			decreaseLife(e.Value.(*Particle))
		}
		if config.General.Optimisation {
			if dead(e.Value.(*Particle)) {
				//La mettre à la fin
				e.Value.(*Particle).Opacity = 0
				s.Content.MoveToBack(e)
			}
		} else {
			if dead(e.Value.(*Particle)) {
				e.Value.(*Particle).Opacity = 0
			}
		}

		e = e.Next()
	}
	if config.General.SpawnAtMouse {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			count()
			for i := 0; i < config.General.SpawnPerClick; i++ {
				s.Content.PushFront(createParticule())
			}
		} else {
			col = 0
		}
	} else {
		if spawnrate < 1 {
			spawnrateadd()
		} else {
			for spawnrate >= 1 {
				if config.General.Optimisation {
					if dead(s.Content.Back().Value.(*Particle)) {
						s.Content.Back().Value = createParticule()
					} else {
						s.Content.PushFront(createParticule())
					}
					spawnrate--
				} else {
					s.Content.PushFront(createParticule())
					spawnrate--
				}
			}
		}
	}

}
func count() {
	col += 0.01
}
func GetLen() int {
	return X
}
func spawnrateadd() {
	spawnrate += config.General.SpawnRate
}

func dead(p *Particle) bool {
	return (outwindows(p) || outLife(p))
}

func outwindows(p *Particle) bool {
	if p.PositionX > float64(config.General.WindowSizeX) || p.PositionX < 0-10*config.General.ScaleX || p.PositionY > float64(config.General.WindowSizeY) || p.PositionY < 0-10*config.General.ScaleY {
		return true
	}
	return false
}
func outLife(p *Particle) bool {
	return p.Lifespan == 0
}

func upPosition(p *Particle) {
	p.PositionX += p.SpeedX
	if config.General.Gravity {
		p.SpeedY += config.General.GravityVal
	}
	p.PositionY += p.SpeedY

}

func decreaseLife(p *Particle) {
	p.Lifespan -= 1
	if config.General.Opacity == 1 {
		decreaseOpacity(p)
	} else {
		increaseOpacity(p)
	}
}

func decreaseOpacity(p *Particle) {
	p.Opacity -= 1 / config.General.Lifespan
}
func increaseOpacity(p *Particle) {
	p.Opacity += 1 / config.General.Lifespan
	if p.Lifespan == 0 {
		p.Opacity = 0
	}
}
