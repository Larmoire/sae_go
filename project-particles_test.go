package main

import (
	"project-particles/config"
	"project-particles/particles"
	"testing"
)


func TestSpawn(t *testing.T) {

	config.Get("config4.json")
	Particles := particles.NewSystem()
	
	//On vérifie que les particules sont bien créées
	if Particles.Content.Len() != config.General.InitNumParticles {
		t.Errorf("Spawn failed, got: %d, want: %d.", Particles.Content.Len(), config.General.InitNumParticles)
	}
}

func TestOnScreen(t *testing.T) {

	config.Get("config4.json")
	Particles := particles.NewSystem()

	for e := Particles.Content.Front(); e != nil; e = e.Next() {
		if !config.General.RandomSpawn {
			//Si le spawn n'est pas random, on vérifie que les particules sont bien placées à SpawnX et SpawnY
			if int(e.Value.(*particles.Particle).PositionX) != config.General.SpawnX {
				t.Errorf("Spawn failed, got: %f, want: %d.", e.Value.(*particles.Particle).PositionX, config.General.SpawnX)
			}
	
			if int(e.Value.(*particles.Particle).PositionY) != config.General.SpawnY {
				t.Errorf("Spawn failed, got: %f, want: %d.", e.Value.(*particles.Particle).PositionY, config.General.SpawnY)
			}
		} else {
			//Si le spawn est random, on vérifie que les particules sont bien dans l'ecran
			if int(e.Value.(*particles.Particle).PositionX) > config.General.WindowSizeX {
				t.Errorf("Spawn failed, got: %f, want: %d.", e.Value.(*particles.Particle).PositionX, config.General.WindowSizeX)
			}
			if int(e.Value.(*particles.Particle).PositionY) > config.General.WindowSizeY {
				t.Errorf("Spawn failed, got: %f, want: %d.", e.Value.(*particles.Particle).PositionY, config.General.WindowSizeY)
			}
		}
		
	}
}
