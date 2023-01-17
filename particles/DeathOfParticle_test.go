package particles

import (
	"container/list"
	"project-particles/config"
	"testing"
)

func TestDeath(t *testing.T) {
	config.Get("../config.json")
	//On initialise la particule dans la fenêtre mais avec le lifespan à 1
	p := &Particle{PositionX: 400, PositionY: 300, Lifespan: 1, Opacity: 1}

	//On update la particule
	p.DecreaseLife()
	p.DeathOfParticle(&System{}, &list.Element{})

	//On vérifie que la particule a bien disparu
	if p.Opacity != 0 {
		t.Errorf("Death() failed")
	}
}

func TestDeath2(t *testing.T) {
	config.Get("../config.json")
	//On initialise la particule en dehors de la fenêtre
	p := &Particle{PositionX: -10 * config.General.ScaleX, PositionY: -10 * config.General.ScaleY, Lifespan: 10, Opacity: 1}

	//On update la particule
	p.DecreaseLife()
	p.DeathOfParticle(&System{}, &list.Element{})

	//On vérifie que la particule a bien disparu
	if p.Opacity != 0 {
		t.Errorf("Death() failed")
	}
}

func TestDeath3(t *testing.T) {
	config.Get("../config.json")
	//On initialise la particule avec le lifespan désactivé et en dehors de la fenêtre
	p := &Particle{PositionX: -10 * config.General.ScaleX, PositionY: -10 * config.General.ScaleY, Lifespan: -1, Opacity: 1}

	//On update la particule
	p.DecreaseLife()
	p.DeathOfParticle(&System{}, &list.Element{})
	//On vérifie que la particule n'a pas disparu
	if p.Opacity == 0 {
		t.Errorf("Death() failed")
	}
}

func TestDeath4(t *testing.T) {
	config.Get("../config.json")
	//On initialise la particule avec le lifespan désactivé et dans la fenêtre
	p := &Particle{PositionX: 400, PositionY: 300, Lifespan: -1, Opacity: 1}

	//On update la particule
	p.DecreaseLife()
	p.DeathOfParticle(&System{}, &list.Element{})

	//On vérifie que la particule n'a pas disparu
	if p.Opacity == 0 {
		t.Errorf("Death() failed")
	}
}
