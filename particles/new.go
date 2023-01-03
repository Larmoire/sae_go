package particles

import (
	"container/list"
	"math/rand"
	"project-particles/config"
	"time"
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

var NbPart int

func NewSystem() System {
	rand.Seed(time.Now().UnixNano())
	l := list.New()
	for i := 0; i < (config.General.InitNumParticles); i++ {
		l.PushFront(createParticule())
	}
	NbPart = config.General.InitNumParticles
	return System{Content: l}
}

func createParticule() *Particle {
	var ParticuleAMettre *Particle
	setSpeed()
	setSpawn()
	ParticuleAMettre = (&Particle{
		PositionX: PosX,
		PositionY: PosY,
		ScaleX:    config.General.ScaleX, ScaleY: config.General.ScaleY,
		ColorRed: config.General.ColorRed, ColorGreen: config.General.ColorGreen, ColorBlue: config.General.ColorBlue,
		Opacity:   config.General.Opacity,
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
