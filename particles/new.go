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

func NewSystem() System {
	rand.Seed(time.Now().UnixNano())
	l := list.New()
	for i := 0; i < (config.General.InitNumParticles); i++ {
		l.PushFront(CreateParticule())
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
	Speedx = rand.Float64()*(config.General.SpeedXmax-config.General.SpeedXmin) + config.General.SpeedXmin
	Speedy = rand.Float64()*(config.General.SpeedYmax-config.General.SpeedYmin) + config.General.SpeedYmin
}
func setColor() {
	//Si le fade est activé, on définit la couleur en fonction de col, la durée du click
	if config.General.Fade {
		Red = config.General.ColorRed - col
		Green = config.General.ColorGreen - col
		Blue = config.General.ColorBlue - col
	} else if config.General.RVBchange {
		//Changement de la couleur en fonction des touches du clavier
		if ebiten.IsKeyPressed(ebiten.KeyR) {
			Red -= 0.0001
		}
		if ebiten.IsKeyPressed(ebiten.KeyV) {
			Green -= 0.0001
		}
		if ebiten.IsKeyPressed(ebiten.KeyB) {
			Blue -= 0.0001
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
		PosX = rand.Float64() * (float64(config.General.WindowSizeX) - 2)
		PosY = rand.Float64() * (float64(config.General.WindowSizeY) - 2)
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
