package particles

import (
	"math"
	"math/rand"
	"project-particles/config"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.

var spawnrate float64 = config.General.SpawnRate //On peut tester avec 0.017 pour avoir environ 1 particule par seconde
var col float64
var buttonPressed bool

func (s *System) Update() {

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if x >= 0 && x <= 110 && y >= 0 && y <= 30 {
			if !buttonPressed {
				buttonPressed = true
				Gravity = !Gravity
			}
		}
		if x >= 0 && x <= 110 && y >= 40 && y <= 70 {
			if !buttonPressed {
				buttonPressed = true
				Bounce = !Bounce
				if !Bounce {
					ColorBounce = false
				}
			}
		}
		if x >= 135 && x <= 245 && y >= 40 && y <= 70 && Bounce {
			if !buttonPressed {
				buttonPressed = true
				ColorBounce = !ColorBounce
			}
		}
		if x >= 0 && x <= 110 && y >= 80 && y <= 110 {
			if !buttonPressed {
				buttonPressed = true
				RandomSpeed = !RandomSpeed
			}
		}
		if x >= 0 && x <= 110 && y >= 120 && y <= 150 {
			if !buttonPressed {
				buttonPressed = true
				RVBchange = !RVBchange
			}
		}
		if x >= 0 && x <= 110 && y >= 160 && y <= 190 {
			if !buttonPressed {
				buttonPressed = true
				SpawnAtMouse = !SpawnAtMouse
			}
		}
		if x >= 0 && x <= 110 && y >= 200 && y <= 230 {
			if !buttonPressed {
				buttonPressed = true
				SpeedFix = !SpeedFix
			}
		}
	} else {
		buttonPressed = false
	}

	X = s.Content.Len()

	//On compte le nombre de particules mortes de la liste
	countdead := 0
	for e := s.Content.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Particle)
		if dead(p) {
			countdead += 1
		}
	}
	//On reset la liste si toutes les particules sont mortes ou si on appuie sur Tab
	if countdead == s.Content.Len() || ebiten.IsKeyPressed(ebiten.KeyTab) {
		s.Content.Init()
	}

	//On génère la seed random à partir du temps pour qu'elle soit différente à chaque fois
	rand.Seed(time.Now().UnixNano())

	//On définit e, le début de la liste de particules
	e := s.Content.Front()

	//Tant que e n'est pas nul, on fait les actions suivantes
	for e != nil {

		//On fait avancer la particule
		if config.General.Orbital {
			//On fait avancer la particule en mode orbital
			updateOrbit(e.Value.(*Particle), float64(config.General.WindowSizeX)/2, float64(config.General.WindowSizeY)/2, 5)
			rotateParticle(e.Value.(*Particle), float64(config.General.WindowSizeX)/2, float64(config.General.WindowSizeY)/2, 5)
		} else {
			upPosition(e.Value.(*Particle))
		}

		//Si le lifespan est activé, on enlève 1 à la durée de vie de la particule
		if e.Value.(*Particle).Lifespan != -1 {
			decreaseLife(e.Value.(*Particle))
		}

		//Si l'optimisation est activée, on met la particules à la fin de la liste uniquement si elle est morte
		if config.General.Optimisation {
			if dead(e.Value.(*Particle)) {

				//La mettre à la fin
				e.Value.(*Particle).Opacity = 0
				s.Content.MoveToBack(e)
			}

			//Et si elle n'est pas activée, on met son opacité à 0 uniquement si elle est morte
		} else if !Bounce {
			if dead(e.Value.(*Particle)) {
				e.Value.(*Particle).Opacity = 0
			}
		} else {
			bounce(e.Value.(*Particle))
		}

		//On passe à la particule suivante
		e = e.Next()
	}

	//On regarde si l'option SpawnAtMouse est activée
	if SpawnAtMouse {
		setColor()
		//On génère une particule à la position de la souris si le click gauche est enfoncé
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {

			//On compte la durée du click pour gérer le fade
			count()

			//On génère SpawnPerClick particules à la position de la souris
			for i := 0; i < config.General.SpawnPerClick; i++ {
				if col < 1 || !config.General.Fade {
					s.Content.PushFront(CreateParticule())
				}
			}

			//Si le click gauche n'est pas enfoncé, on remet la durée du click à 0
		} else {
			col = 0
		}
		//Si le SpawnRate est inférieur à 1, alors on le garde en mémoire et on l'ajoute à un compteur, pour par exemple générer une particules tout les deux updates (pour 0.5 par exemple)
	} else {
		if spawnrate < 1 {
			spawnrateadd()
		} else {

			//Sinon, on génère SpawnRate particules
			for spawnrate >= 1 {
				if config.General.Optimisation {
					if dead(s.Content.Back().Value.(*Particle)) {
						s.Content.Back().Value = CreateParticule()
					} else {
						s.Content.PushFront(CreateParticule())
					}
					spawnrate--
				} else {
					s.Content.PushFront(CreateParticule())
					spawnrate--
				}
			}
		}
	}
}
func UpdateLenList(i int) {
	X += i
}
func GetSpeedFixState() bool {
	return SpeedFix
}
func GetSpawnAtMouseState() bool {
	return SpawnAtMouse
}
func GetRandomSpeedState() bool {
	return RandomSpeed
}
func GetBounceState() bool {
	return Bounce
}
func GetGravityState() bool {
	return Gravity
}
func GetColorBounceState() bool {
	return ColorBounce
}
func GetRVBChangeState() bool {
	return RVBchange
}
func rotateParticle(particle *Particle, centerX, centerY, angle float64) {
	// Calcule la distance entre le centre de rotation et la particule
	distance := math.Sqrt(math.Pow(particle.PositionX-centerX, 2) + math.Pow(particle.PositionY-centerY, 2))

	// Calcule l'angle actuel de la particule par rapport au centre de rotation
	currentAngle := math.Atan2(particle.PositionY-centerY, particle.PositionX-centerX)

	// Applique la rotation à l'angle actuel de la particule
	newAngle := currentAngle + angle

	// Calcule les nouvelles coordonnées de la particule en utilisant la distance et l'angle mis à jour
	particle.PositionX = centerX + distance*math.Cos(newAngle)
	particle.PositionY = centerY + distance*math.Sin(newAngle)
}
func updateOrbit(particle *Particle, centerX, centerY, orbitSpeed float64) {
	// Calcule la distance entre le centre de l'orbite et la particule
	distance := math.Sqrt(math.Pow(particle.PositionX-centerX, 2) + math.Pow(particle.PositionY-centerY, 2))

	// Calcule l'angle actuel de la particule par rapport au centre de l'orbite
	currentAngle := math.Atan2(particle.PositionY-centerY, particle.PositionX-centerX)

	// Met à jour l'angle de la particule en fonction de la vitesse d'orbite
	newAngle := currentAngle + orbitSpeed

	// Calcule les nouvelles coordonnées de la particule en utilisant la distance et l'angle mis à jour
	particle.PositionX = centerX + distance*math.Cos(newAngle)
	particle.PositionY = centerY + distance*math.Sin(newAngle)
}

func bounce(p *Particle) {
	//la faire rebondir sur les bords
	if p.PositionX <= 0 || p.PositionX+p.SpeedX+10*config.General.ScaleX >= float64(config.General.WindowSizeX) {
		p.SpeedX = -p.SpeedX
		if ColorBounce {
			p.ColorRed, p.ColorGreen, p.ColorBlue = rand.Float64(), rand.Float64(), rand.Float64()
		}
	}
	if p.PositionY <= 0 || p.PositionY+p.SpeedY+10*config.General.ScaleY >= float64(config.General.WindowSizeY) {
		p.SpeedY = -p.SpeedY
		if ColorBounce {
			p.ColorRed, p.ColorGreen, p.ColorBlue = rand.Float64(), rand.Float64(), rand.Float64()
		}
	}
}

//Le compteur de durée du click pour le fade
func count() {
	col += 0.01
}

//Pour le debug, on affiche le nombre de particules
func GetLen() int {
	return X
}

//Pour la mise en mémoire du reste du spawnrate
func spawnrateadd() {
	spawnrate += config.General.SpawnRate
}

//Check si la particule est morte
func dead(p *Particle) bool {
	return (Outwindows(p) || outLife(p))
}

//Renvoie true si la particule sort complétement de l'écran
func Outwindows(p *Particle) bool {
	if p.PositionX > float64(config.General.WindowSizeX) || p.PositionX < 0-10*config.General.ScaleX || p.PositionY > float64(config.General.WindowSizeY) || p.PositionY < 0-10*config.General.ScaleY {
		return true
	}
	return false
}

//Renvoie true si la particule n'a plus de LifeSpan
func outLife(p *Particle) bool {
	return p.Lifespan == 0
}

//Augmente la position de la particule
func upPosition(p *Particle) {
	if !SpeedFix {
		p.PositionX += p.SpeedX
		//Si la particule est en mode gravity, on augmente sa vitesse en y
		if Gravity {
			p.SpeedY += config.General.GravityVal
		}
		if Bounce && Gravity && (p.PositionY+p.SpeedY+10*config.General.ScaleY >= float64(config.General.WindowSizeY)) {
			p.SpeedY = 0
			p.SpeedX = 0
		}
		p.PositionY += p.SpeedY
	}
}

//Réduit le lifespan de la particule
func decreaseLife(p *Particle) {
	p.Lifespan -= 1
	//Si l'opacity de base est à 1, on la réduit en fonction de la durée de vie restante
	if config.General.Opacity == 1 {
		decreaseOpacity(p)
		//Sinon, on l'augmente en fonction de la durée de vie restante
	} else {
		increaseOpacity(p)
	}
}

//Enlève l'opacity en fonction de la durée de vie restante
func decreaseOpacity(p *Particle) {
	p.Opacity -= 1 / config.General.Lifespan
}

//Augmente l'opacity en fonction de la durée de vie restante
func increaseOpacity(p *Particle) {
	p.Opacity += 1 / config.General.Lifespan
	if p.Lifespan == 0 {
		p.Opacity = 0
	}
}
