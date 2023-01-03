package particles

import (
	"container/list"
	"math/rand"
	"project-particles/config"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.

var acX int
var acY int

var PosX float64
var PosY float64

var Speedx float64
var Speedy float64

var NbPart int
var X int

var Red, Green, Blue float64 = 1, 1, 1

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
	setColor()
	ParticuleAMettre = (&Particle{
		PositionX: PosX,
		PositionY: PosY,
		ScaleX:    config.General.ScaleX, ScaleY: config.General.ScaleY,
		ColorRed: Red, ColorGreen: Green, ColorBlue: Blue,
		Opacity: config.General.Opacity,
		SpeedX:  Speedx,
		SpeedY:  Speedy,

		Lifespan: config.General.Lifespan,
	})
	return ParticuleAMettre
}

func setSpeed() {
	Speedx = rand.Float64()*(config.General.SpeedXmax-config.General.SpeedXmin) + config.General.SpeedXmin
	Speedy = rand.Float64()*(config.General.SpeedYmax-config.General.SpeedYmin) + config.General.SpeedYmin

}
func setColor() {
	if config.General.Fade {
		Red = config.General.ColorRed - col
		Green = config.General.ColorGreen - col
		Blue = config.General.ColorBlue - col
	} else if ebiten.IsKeyPressed(ebiten.KeyR) {
		Red = Red - 0.0001
	} else if ebiten.IsKeyPressed(ebiten.KeyG) {
		Green = Green - 0.0001
	} else if ebiten.IsKeyPressed(ebiten.KeyB) {
		Blue = Blue - 0.0001
	}
}
func setSpawn() {
	if config.General.RandomSpawn {
		PosX = rand.Float64() * (float64(config.General.WindowSizeX) - 2)
		PosY = rand.Float64() * (float64(config.General.WindowSizeY) - 2)
	} else if !config.General.SpawnAtMouse {
		PosX = float64(config.General.SpawnX)
		PosY = float64(config.General.SpawnY)
	} else {
		acX, acY = ebiten.CursorPosition()
		PosX = float64(acX)
		PosY = float64(acY)
	}
}
