package particles

import (
	"container/list"
	"math/rand"
	"project-particles/config"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.
func createParticule() *Particle {
	var PosX float64
	var PosY float64
	var ParticuleAMettre *Particle
	if !config.General.RandomSpawn {
		PosX = float64(config.General.SpawnX)
		PosY = float64(config.General.SpawnY)
	} else {
		PosX = float64(rand.Float64() * (float64(config.General.WindowSizeX) - 2))
		PosY = float64(rand.Float64() * (float64(config.General.WindowSizeY) - 2))
	}
	ParticuleAMettre = (&Particle{
		PositionX: PosX,
		PositionY: PosY,
		ScaleX:    1, ScaleY: 1,
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity:   1,
		SpeedX:    rand.Float64()*(config.General.SpeedXmax-config.General.SpeedXmin) + config.General.SpeedXmin,
		SpeedY:    rand.Float64()*(config.General.SpeedYmax-config.General.SpeedYmin) + config.General.SpeedYmin,
		SpawnRate: config.General.SpawnRate,
	})
	return ParticuleAMettre
}
func NewSystem() System {
	l := list.New()

	for i := 0; i < (config.General.InitNumParticles); i++ {
		l.PushFront(createParticule())
	}
	return System{Content: l}
}
