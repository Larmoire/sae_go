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

var PosX float64
var PosY float64

var Speedx float64
var Speedy float64

func NewSystem() System {
	l := list.New()
	for i := 0; i < (config.General.InitNumParticles); i++ {
		l.PushFront(createParticule())
	}
	return System{Content: l}
}

func createParticule() *Particle {
	var ParticuleAMettre *Particle
	setSpeed()
	setSpawn()
	ParticuleAMettre = (&Particle{
		PositionX: PosX,
		PositionY: PosY,
		ScaleX:    1, ScaleY: 1,
		ColorRed: 1, ColorGreen: 0.5, ColorBlue: 0.5,
		Opacity:   1,
		SpeedX:    Speedx,
		SpeedY:    Speedy,
		SpawnRate: config.General.SpawnRate,
		Lifespan:  config.General.Lifespan,
	})
	return ParticuleAMettre
}

func setSpeed() {
	Speedx = rand.Float64()*(config.General.SpeedXmax-config.General.SpeedXmin) + config.General.SpeedXmin
	Speedy = rand.Float64()*(config.General.SpeedYmax-config.General.SpeedYmin) + config.General.SpeedYmin

}

func setSpawn() {
	if config.General.RandomSpawn {
		PosX = rand.Float64() * (float64(config.General.WindowSizeX) - 2)
		PosY = rand.Float64() * (float64(config.General.WindowSizeY) - 2)
	} else {
		PosX = float64(config.General.SpawnX)
		PosY = float64(config.General.SpawnY)
	}
}
