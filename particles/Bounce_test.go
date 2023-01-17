package particles

import (
	"testing"
)

func TestBounce(t *testing.T) {
	//On initialise la particule
	p := Particle{PositionX: 0, PositionY: 0, SpeedX: 1, SpeedY: 1}
	//On fait rebondir la particule
	p.Bounce()
	//On vérifie que la particule a bien rebondi
	if p.SpeedX != -1 || p.SpeedY != -1 {
		t.Errorf("Bounce() failed")
	}
}

func TestBounce2(t *testing.T) {
	//On initialise la particule
	p := Particle{PositionX: 0, PositionY: 0, SpeedX: -1, SpeedY: -1}
	//On fait rebondir la particule
	p.Bounce()
	//On vérifie que la particule a bien rebondi
	if p.SpeedX != 1 || p.SpeedY != 1 {
		t.Errorf("Bounce() failed")
	}
}

func TestBounce3(t *testing.T) {
	//On initialise la particule
	p := Particle{PositionX: 0, PositionY: 0, SpeedX: 1, SpeedY: -1}
	//On fait rebondir la particule
	p.Bounce()
	//On vérifie que la particule a bien rebondi
	if p.SpeedX != -1 || p.SpeedY != 1 {
		t.Errorf("Bounce() failed")
	}
}

func TestBounce4(t *testing.T) {
	//On initialise la particule
	p := Particle{PositionX: 0, PositionY: 0, SpeedX: -1, SpeedY: 1}
	//On fait rebondir la particule
	p.Bounce()
	//On vérifie que la particule a bien rebondi
	if p.SpeedX != 1 || p.SpeedY != -1 {
		t.Errorf("Bounce() failed")
	}
}

func TestBounce5(t *testing.T) {
	//On initialise la particule
	p := Particle{PositionX: 0, PositionY: 0, SpeedX: 0, SpeedY: 0}
	//On fait rebondir la particule
	p.Bounce()
	//On vérifie que la particule a bien rebondi
	if p.SpeedX != 0 || p.SpeedY != 0 {
		t.Errorf("Bounce() failed")
	}
}
