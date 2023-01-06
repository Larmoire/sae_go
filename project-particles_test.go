package main

import (
	"project-particles/config"
	"project-particles/particles"
	"testing"
)


func TestSpawn(t *testing.T) {

	config.Get("config_part4.json")
	Particles := particles.NewSystem()

	//On vérifie que les particules sont bien créées
	if Particles.Content.Len() != config.General.InitNumParticles {
		t.Errorf("Spawn failed, got: %d, want: %d.", Particles.Content.Len(), config.General.InitNumParticles)
	}
}

func TestOnScreen(t *testing.T) {

	config.Get("config_part4.json")
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

func TestSpawnRate(t *testing.T) {
	
	config.Get("config_part4.json")
	Particles := particles.NewSystem()
	Particles.Update()
	//Vérification que le nombre de particules est bien égal à InitNumParticles + SpawnRate au bout d'un update si SpawnRate > 1
	if config.General.SpawnRate >= 1 {
		for i := 0; i < int(config.General.SpawnRate); i++ {
			if Particles.Content.Len() != config.General.InitNumParticles+i+1 {
				t.Errorf("Spawn failed, got: %d, want: %d.", Particles.Content.Len(), config.General.InitNumParticles+i+1)
			}
		}
	//Si le SpawnRate est inférieur à 1, on vérifie qu'une particule est générée au bout de deux appels de Update pour un SpawnRate de 0.5 par exemple
	} else {
		var count float64 
		if config.General.SpawnRate != 0 {
			//On regarde si le SpawnRate est inférieur à 1, tant qu'il l'est on l'incrémente de lui-même à chaque update, et on compte le nombre de fois que cela est produit
			for spawnr := config.General.SpawnRate ; spawnr <= 1.1; spawnr += config.General.SpawnRate {
				Particles.Update()
				count++
			}
			//On regarde enfin si le nombre de particules générées est bien égal à InitNumParticles + le nombre de fois que le SpawnRate a été incrémenté
			if Particles.Content.Len() != config.General.InitNumParticles+int(count*config.General.SpawnRate) {
				t.Errorf("Spawn failed, got: %d, want: %d.", Particles.Content.Len(), config.General.InitNumParticles+int(count*config.General.SpawnRate))
			}
		}
	}
}

func TestSpeed(t *testing.T) {

	config.Get("config_part4.json")
	Particles := particles.NewSystem()
	//On créée un tableau contenant les positions des particules avant l'update
	var tabPos [][]float64
	for e := Particles.Content.Front(); e != nil; e = e.Next() {
		tabPos = append(tabPos, []float64{e.Value.(*particles.Particle).PositionX, e.Value.(*particles.Particle).PositionY})
	}
	var count int
	//On effectue une update
	Particles.Update()
	//On regarde si les positions actuelles des particules sont bien égales aux positions précédentes + la vitesse
	for e := Particles.Content.Front(); e != nil; e = e.Next() {
		//Le tabPos[count][0] correspond à la position X de la particule 'count' avant l'update
		if e.Value.(*particles.Particle).PositionX != tabPos[count][0]+e.Value.(*particles.Particle).SpeedX {
			t.Errorf("Spawn failed, got: %f, want: %f.", e.Value.(*particles.Particle).PositionX, tabPos[count][0]+e.Value.(*particles.Particle).SpeedX)
		}
		//Le tabPos[count][1] correspond à la position Y de la particule 'count' avant l'update
		if e.Value.(*particles.Particle).PositionY != tabPos[count][1]+e.Value.(*particles.Particle).SpeedY {
			t.Errorf("Spawn failed, got: %f, want: %f.", e.Value.(*particles.Particle).PositionY, tabPos[count][1]+e.Value.(*particles.Particle).SpeedY )
		}
		//Le compteur sert à parcourir le tableau contenant les positions précédentes
		count++
	}
}
