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

var spawnrate float64 = config.General.SpawnRate

func (s *System) Update() {

	//On génère la seed random à partir du temps pour qu'elle soit différente à chaque fois
	rand.Seed(time.Now().UnixNano())

	//On définit e, le début de la liste de particules
	e := s.Content.Front()

	//Tant que e n'est pas nul, on fait les actions suivantes
	for e != nil {

		//On fait avancer la particule
		upPosition(e.Value.(*Particle))

		//On passe à la particule suivante
		e = e.Next()
	}

	//Si le SpawnRate est inférieur à 1, alors on le garde en mémoire et on l'ajoute à un compteur, pour par exemple générer une particules tout les deux updates (pour 0.5 par exemple)
	if spawnrate < 1 {
		spawnrateadd()
	} else {

	//Sinon, on génère SpawnRate particules
	for spawnrate >= 1 {
			s.Content.PushFront(CreateParticule())
			spawnrate--
		}
	}
}

//Pour la mise en mémoire du reste du spawnrate
func spawnrateadd() {
	spawnrate += config.General.SpawnRate
}


//Augmente la position de la particule
func upPosition(p *Particle) {
	p.PositionX += p.SpeedX
	p.PositionY += p.SpeedY

}

