package particles

import (
	"container/list"
	"math/rand"
	Extensions "project-particles/Extension"
	pictures "project-particles/Extension/Pictures"
	"project-particles/config"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.

// Fonction présente : NewSystem, CreateParticule, setSpeed, setColor, setSpawn

var acX int
var acY int

var PosX float64
var PosY float64

var Speedx float64
var Speedy float64

var NbPart int
var X int

var Red, Green, Blue float64

var NombreDeParticules int = config.General.InitNumParticles

func NewSystem() System {

	l := list.New()
	ImageIn(l)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < (config.General.InitNumParticles); i++ {
		if !Extensions.SpawnAtMouse {
			l.PushFront(CreateParticule())
		}
	}
	return System{Content: l}
}

// Fonction pour générer une particule
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
		CibleX:   float64(PosX),
		CibleY:   float64(PosY),
	})
	return ParticuleAMettre
}

// La vitesse est aléatoire entre les valeurs min et max
func setSpeed() {
	if Extensions.RandomSpeed {
		Speedx = rand.Float64()*(config.General.SpeedXmax-config.General.SpeedXmin) + config.General.SpeedXmin
		Speedy = rand.Float64()*(config.General.SpeedYmax-config.General.SpeedYmin) + config.General.SpeedYmin
	} else {
		Speedx = config.General.SpeedX
		Speedy = config.General.SpeedY
	}

}
func setColor() {
	//Si le fade est activé, on définit la couleur en fonction de col, la durée du click
	if Extensions.Fade {
		Red = config.General.ColorRed - col
		Green = config.General.ColorGreen - col
		Blue = config.General.ColorBlue - col
	} else if Extensions.RGBchange {
		//Changement de la couleur en fonction des touches du clavier
		if ebiten.IsKeyPressed(ebiten.KeyR) { //Rouge en pressant R
			Red, Green, Blue = 1, 0, 0
		}
		if ebiten.IsKeyPressed(ebiten.KeyG) { //Vert en pressant G
			Red, Green, Blue = 0, 1, 0
		}
		if ebiten.IsKeyPressed(ebiten.KeyB) { //Bleu en pressant B
			Red, Green, Blue = 0, 0, 1
		}
		if ebiten.IsKeyPressed(ebiten.KeyY) { //Jaune en pressant Y
			Red, Green, Blue = 1, 1, 0
		}
		if ebiten.IsKeyPressed(ebiten.KeyC) { //Cyan en pressant C
			Red, Green, Blue = 0, 1, 1
		}
		if ebiten.IsKeyPressed(ebiten.KeyP) { //Violet en pressant P
			Red, Green, Blue = 1, 0, 1
		}
	} else {
		//Sinon, on définit la couleur en fonction du config
		Red = config.General.ColorRed
		Green = config.General.ColorGreen
		Blue = config.General.ColorBlue
	}
}
func setSpawn() {
	if Extensions.RandomSpawn {
		//Si randomspawn est true, on définit la position de la particule aléatoirement
		PosX = rand.Float64() * ((float64(config.General.WindowSizeX)) - 10*config.General.ScaleX)
		PosY = rand.Float64() * ((float64(config.General.WindowSizeY)) - 10*config.General.ScaleY)
	} else {
		//Sinon, on la met à une valeur fixe du config
		PosX = float64(config.General.SpawnX)
		PosY = float64(config.General.SpawnY)
	}
	if Extensions.SpawnAtMouse && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		//Sinon, on regarde si SpawnAtMouse est activé pour mettre la position à celle de la souris
		acX, acY = ebiten.CursorPosition()
		PosX = float64(acX)
		PosY = float64(acY)
	}
}
func ImageIn(l *list.List) {
	if config.General.Pictures != "" {

		h, w, image := pictures.Gettabpixels()
		rand.Seed(time.Now().UnixNano())
		for k := 0; k < len(image); k++ {
			l := rand.Intn(len(image))
			image[k], image[l] = image[l], image[k]
		}

		precision := 10.0
		NombreDeParticules = h * w

		config.General.ScaleX = config.General.ScaleX / precision
		config.General.ScaleY = config.General.ScaleY / precision

		for i := 0; i < NombreDeParticules; i++ {
			NbPart += 1
			l.PushFront(CreateParticule())
			l.Front().Value.(*Particle).CibleX, l.Front().Value.(*Particle).CibleY = float64(image[i][0])*config.General.ScaleX*precision, float64(image[i][1])*config.General.ScaleY*precision
			l.Front().Value.(*Particle).ColorRed, l.Front().Value.(*Particle).ColorGreen, l.Front().Value.(*Particle).ColorBlue = float64(image[i][2])/255, float64(image[i][3])/255, float64(image[i][4])/255
			l.Front().Value.(*Particle).Opacity = float64(image[i][5]) / 255
			if config.General.Spawnimg {
				l.Front().Value.(*Particle).PositionX, l.Front().Value.(*Particle).PositionY = l.Front().Value.(*Particle).CibleX, l.Front().Value.(*Particle).CibleY

			}
		}

	}
}
