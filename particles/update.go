package particles

import "project-particles/config"

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.

var spawnrate float64 = config.General.SpawnRate //On peut tester avec 0.17 pour avoir environ 1 particule par seconde

func (s *System) Update() {
	e := s.Content.Front()
	for e != nil {
		e.Value.(*Particle).PositionX += e.Value.(*Particle).SpeedX
		e.Value.(*Particle).PositionY += e.Value.(*Particle).SpeedY

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
