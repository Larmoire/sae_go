package particles

import (
	Extensions "project-particles/Extension"
	"project-particles/config"
	"testing"
)

func TestUpdatePos(t *testing.T) {
	//On initialise la particule
	p := Particle{PositionX: 0, PositionY: 0, SpeedX: 1, SpeedY: 1}
	//On fait bouger la particule
	p.UpdatePos()
	//On vérifie que la particule a bien bougé
	if p.PositionX != 1 || p.PositionY != 1 {
		t.Errorf("UpdatePos() failed")
	}
}

func TestUpdatePos2(t *testing.T) {
	//On initialise la particule
	p := Particle{PositionX: 0, PositionY: 0, SpeedX: -1, SpeedY: -1}
	//On fait bouger la particule
	p.UpdatePos()
	//On vérifie que la particule a bien bougé
	if p.PositionX != -1 || p.PositionY != -1 {
		t.Errorf("UpdatePos() failed")
	}
}

func TestUpdatePos5(t *testing.T) {
	//On initialise la particule
	p := Particle{PositionX: 0, PositionY: 0, SpeedX: 0, SpeedY: 0}
	//On fait bouger la particule
	p.UpdatePos()
	//On vérifie que la particule a bien bougé
	if p.PositionX != 0 || p.PositionY != 0 {
		t.Errorf("UpdatePos() failed")
	}
}

func TestUpdatePos3(t *testing.T) {
	config.Get("../config.json")
	Extensions.Gravity = true
	//On initialise la particule
	p := Particle{PositionX: 0, PositionY: 0, SpeedX: 1, SpeedY: 1}
	//On fait bouger la particule
	p.UpdatePos()
	//On vérifie que la particule a bien bougé
	if p.PositionX != 1 || p.PositionY != 1+config.General.GravityVal {
		t.Errorf("UpdatePos() failed")
	}
}

func TestUpdatePos4(t *testing.T) {
	config.Get("../config.json")
	Extensions.Gravity = true
	//On initialise la particule
	p := Particle{PositionX: 0, PositionY: 0, SpeedX: -1, SpeedY: -1}
	//On fait bouger la particule
	p.UpdatePos()
	//On vérifie que la particule a bien bougé
	if p.PositionX != -1 || p.PositionY != -1+config.General.GravityVal {
		t.Errorf("UpdatePos() failed")
	}
}
