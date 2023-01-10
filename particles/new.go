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

var Red, Green, Blue float64

var Gravity bool
var Bounce bool
var ColorBounce bool
var RandomSpeed bool
var RVBchange bool
var SpeedFix bool
var SpawnAtMouse bool

func NewSystem() System {

	Gravity = config.General.Gravity
	Bounce = config.General.Bounce
	ColorBounce = config.General.ColorBounce
	RandomSpeed = config.General.RandomSpeed
	RVBchange = config.General.RVBchange
	SpeedFix = config.General.SpeedFix
	SpawnAtMouse = config.General.SpawnAtMouse

	rand.Seed(time.Now().UnixNano())
	l := list.New()
	for i := 0; i < (config.General.InitNumParticles); i++ {
		if !GetSpawnAtMouseState() {
			l.PushFront(CreateParticule())
		}
	}
	NbPart = config.General.InitNumParticles
	return System{Content: l}
}

//Fonction pour générer une particule
func CreateParticule() *Particle {
	var ParticuleAMettre *Particle
	//On définit les variables vitesse
	setSpeed()
	//On définit les variables position de base
	setSpawn()
	//On définit les variables couleur
	setColor()
	//On créee la particule en fonctions de variables définies juste avant
	ParticuleAMettre = (&Particle{
		PositionX: PosX,
		PositionY: PosY,
		ScaleX:    config.General.ScaleX, ScaleY: config.General.ScaleY,
		ColorRed: Red, ColorGreen: Green, ColorBlue: Blue,
		Opacity:  config.General.Opacity,
		SpeedX:   Speedx,
		SpeedY:   Speedy,
		Lifespan: config.General.Lifespan,
	})
	return ParticuleAMettre
}

//La vitesse est aléatoire entre les valeurs min et max
func setSpeed() {
	if GetRandomSpeedState() {
		Speedx = rand.Float64()*(config.General.SpeedXmax-config.General.SpeedXmin) + config.General.SpeedXmin
		Speedy = rand.Float64()*(config.General.SpeedYmax-config.General.SpeedYmin) + config.General.SpeedYmin
	} else {
		Speedx = config.General.SpeedX
		Speedy = config.General.SpeedY
	}

}
func setColor() {
	//Si le fade est activé, on définit la couleur en fonction de col, la durée du click
	if config.General.Fade {
		Red = config.General.ColorRed - col
		Green = config.General.ColorGreen - col
		Blue = config.General.ColorBlue - col
	} else if GetRVBChangeState() {
		//Changement de la couleur en fonction des touches du clavier
		if ebiten.IsKeyPressed(ebiten.KeyR) {
			Red, Green, Blue = 1, 0, 0
		}
		if ebiten.IsKeyPressed(ebiten.KeyG) {
			Red, Green, Blue = 0, 1, 0
		}
		if ebiten.IsKeyPressed(ebiten.KeyB) {
			Red, Green, Blue = 0, 0, 1
		}
		if ebiten.IsKeyPressed(ebiten.KeyY) {
			Red, Green, Blue = 1, 1, 0
		}
		if ebiten.IsKeyPressed(ebiten.KeyC) {
			Red, Green, Blue = 0, 1, 1
		}
	} else {
		//Sinon, on définit la couleur en fonction du config
		Red = config.General.ColorRed
		Green = config.General.ColorGreen
		Blue = config.General.ColorBlue
	}
}
func setSpawn() {
	if config.General.RandomSpawn {
		//Si randomspawn est true, on définit la position de la particule aléatoirement
		PosX = rand.Float64() * ((float64(config.General.WindowSizeX)) - 10*config.General.ScaleX)
		PosY = rand.Float64() * ((float64(config.General.WindowSizeY)) - 10*config.General.ScaleY)
	} else {
		//Sinon, on la met à une valeur fixe du config
		PosX = float64(config.General.SpawnX)
		PosY = float64(config.General.SpawnY)
	}
	if config.General.SpawnAtMouse && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		//Sinon, on regarde si SpawnAtMouse est activé pour mettre la position à celle de la souris
		acX, acY = ebiten.CursorPosition()
		PosX = float64(acX)
		PosY = float64(acY)
	}
}
