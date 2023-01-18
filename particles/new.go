package particles

import (
	"container/list"
	"math/rand"
	Extensions "project-particles/Extension"
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
var x bool = true

var Red, Green, Blue float64

var NombreDeParticules int = config.General.InitNumParticles

func NewSystem() System {
	//On définit les variables de base des extensions
	Extensions.Gravity = config.General.Gravity
	Extensions.Bounce = config.General.Bounce
	Extensions.ColorBounce = config.General.ColorBounce
	Extensions.RandomSpeed = config.General.RandomSpeed
	Extensions.RGBchange = config.General.RGBchange
	Extensions.SpeedFix = config.General.SpeedFix
	Extensions.SpawnAtMouse = config.General.SpawnAtMouse
	Extensions.Rotate = config.General.Rotate
	Extensions.Fade = config.General.Fade
	Extensions.RandomSpawn = config.General.RandomSpawn
	//On définit la liste de particules
	l := list.New()
	//On génère la seed random à partir du temps pour qu'elle soit différente à chaque fois
	rand.Seed(time.Now().UnixNano())
	//On génère les particules en fonction du nombre de particules défini dans le fichier config
	for i := 0; i < (config.General.InitNumParticles); i++ {
		//On ajoute la particule à la liste si l'extension SpawnAtMouse n'est pas activée
		if !Extensions.SpawnAtMouse {
			l.PushFront(CreateParticule())
		}
	}
	//On renvoie le systeme de particules
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
		Rotation:  0,
		ScaleX:    config.General.ScaleX, ScaleY: config.General.ScaleY,
		ColorRed: Red, ColorGreen: Green, ColorBlue: Blue,
		Opacity:  config.General.Opacity,
		SpeedX:   Speedx,
		SpeedY:   Speedy,
		Lifespan: config.General.Lifespan,
	})
	return ParticuleAMettre
}

func setSpeed() {
	//La vitesse est aléatoire entre les valeurs min et max si l'extension RandomSpeed est activée
	if Extensions.RandomSpeed {
		Speedx = rand.Float64()*(config.General.SpeedXmax-config.General.SpeedXmin) + config.General.SpeedXmin
		Speedy = rand.Float64()*(config.General.SpeedYmax-config.General.SpeedYmin) + config.General.SpeedYmin
	} else {
		//Sinon la vitesse est fixe
		Speedx = config.General.SpeedX
		Speedy = config.General.SpeedY
	}

}
func setColor() {
	//Si l'extention arrows est active, la couleur de base est rouge pour se reperer
	if Extensions.Arrows {
		Red, Green, Blue = 1, 0, 0
		//Sinon, on regarde si le Fade est actif pour définir la couleur en fonction de col, la durée du click
	} else if Extensions.Fade {
		Red = config.General.ColorRed - col
		Green = config.General.ColorGreen - col
		Blue = config.General.ColorBlue - col
		//Sinon on regarde si le RGB est actif
	} else if Extensions.RGBchange {
		//On récupère la touche appuyé au clavier et on la traite dans un switch
		switch string(ebiten.AppendInputChars(nil)) {
		//Si r, alors rouge, si b, alors bleu, etc
		case "r": //red
			Red, Green, Blue = 1, 0, 0
		case "b": //blue
			Red, Green, Blue = 0, 0, 1
		case "y": //yellow
			Red, Green, Blue = 1, 1, 0
		case "g": //green
			Red, Green, Blue = 0.18, 0.54, 0.34
		case "c": //cyan
			Red, Green, Blue = 0, 1, 1
		case "m": //magenta
			Red, Green, Blue = 1, 0, 1
		case "w": //white
			Red, Green, Blue = 1, 1, 1
		case "o": //orange
			Red, Green, Blue = 1, 0.5, 0
		case "p": //purple
			Red, Green, Blue = 0.54, 0.17, 0.89
		case "l": //lime
			Red, Green, Blue = 0.4, 0.80, 0.67
		}
	} else {
		//Sinon, on définit la couleur en fonction du config
		Red = config.General.ColorRed
		Green = config.General.ColorGreen
		Blue = config.General.ColorBlue
	}
}

func setSpawn() {
	//Fonction pour définir les variables de spawn une seule fois.
	initvar()
	//Si randomspawn est true, on définit la position de la particule aléatoirement dans la fenêtre
	if Extensions.RandomSpawn {
		PosX = rand.Float64() * ((float64(config.General.WindowSizeX)) - 10*config.General.ScaleX)
		PosY = rand.Float64() * ((float64(config.General.WindowSizeY)) - 10*config.General.ScaleY)
	}
	//Si SpawnAtMouse est activé et qu'un clic est fait, on met la position à celle de la souris
	if Extensions.SpawnAtMouse && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		acX, acY = ebiten.CursorPosition()
		if !(float64(acX) >= float64(config.General.WindowSizeX)-10*config.General.ScaleX || float64(acX) <= 10*config.General.ScaleX || float64(acY) >= float64(config.General.WindowSizeY)-10*config.General.ScaleY || float64(acY) <= 10*config.General.ScaleY) {
			PosX = float64(acX)
			PosY = float64(acY)
		} else {
			PosX = -100
			PosY = -100
		}
	}
	//Si l'extension Arrows est active, on change le point de spawn avec les flèches directionelles
	if Extensions.Arrows {
		//PosX += 1 Si la touche droite est pressée
		if ebiten.IsKeyPressed(ebiten.KeyRight) && PosX < float64(config.General.WindowSizeX)-10*config.General.ScaleX-10 {
			PosX += 10
		}
		//PosX -= 1 Si la touche gauche est pressée
		if ebiten.IsKeyPressed(ebiten.KeyLeft) && PosX > 0+10*config.General.ScaleX {
			PosX -= 10
		}
		//PosY += 1 Si la touche bas est pressée
		if ebiten.IsKeyPressed(ebiten.KeyDown) && PosY < float64(config.General.WindowSizeY)-10*config.General.ScaleY-10 {
			PosY += 10
		}
		//PosX -= 1 Si la touche haute est pressée
		if ebiten.IsKeyPressed(ebiten.KeyUp) && PosY > 0+10*config.General.ScaleY {
			PosY -= 10
		}
	}
}
func initvar() {
	if x {
		PosX = float64(config.General.SpawnX)
		PosY = float64(config.General.SpawnY)
		x = !x
	}
}
