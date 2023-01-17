package particles

import "testing"

func TestLifeSpan(t *testing.T) {
	//On initialise la particule
	p := Particle{Lifespan: 2}
	//On fait diminuer la durée de vie de la particule
	p.DecreaseLife()
	//On vérifie que la durée de vie a bien diminué
	if p.Lifespan != 1 {
		t.Errorf("DecreaseLife() failed")
	}
}

func TestLifeSpan2(t *testing.T) {
	//On initialise la particule
	p := Particle{Lifespan: -1}
	//On fait diminuer la durée de vie de la particule
	p.DecreaseLife()
	//On vérifie que la durée de vie a bien diminué
	if p.Lifespan != -1 {
		t.Errorf("DecreaseLife() failed")
	}
}

func TestLifeSpan3(t *testing.T) {
	//On initialise la particule
	p := Particle{Lifespan: 0}
	//On fait diminuer la durée de vie de la particule
	p.DecreaseLife()
	//On vérifie que la durée de vie a bien diminué
	if p.Lifespan != 0 {
		t.Errorf("DecreaseLife() failed")
	}
}
