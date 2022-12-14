package particles

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
func (s *System) Update() {

	e := s.Content.Front()
	for e != nil {
		e.Value.(*Particle).PositionX += e.Value.(*Particle).SpeedX
		e.Value.(*Particle).PositionY += e.Value.(*Particle).SpeedY

		e = e.Next()
	}
}
