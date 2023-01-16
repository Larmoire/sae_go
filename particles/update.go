package particles

import (
	"math/rand"
	Extensions "project-particles/Extension"
	"project-particles/config"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.

//fonction présente : Update,UpdateLenList, GetSpeedFixState, GetSpawnAtMouseState, GetRandomSpeedState, GetBounceState, GetGravityState, GetColorBounceState,
//GetRGBChangeState, rotateParticle, updateOrbit,bounce, count, GetLen, spawnrateadd, dead, Outwindows, outLife, upPosition, decreaseLife, decreaseOpacity, increaseOpacity

var spawnrate float64 = config.General.SpawnRate //On peut tester avec 0.017 pour avoir environ 1 particule par seconde
var col float64
var Countdead int

func (s *System) Update() {

	Extensions.GUI()

	X = s.Content.Len()
	Countdead = 0
	//On compte le nombre de particules mortes de la liste
	for e := s.Content.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Particle)
		if p.Lifespan == 0 {
			Countdead += 1
		}
	}
	//On reset la liste si toutes les particules sont mortes ou si on appuie sur Tab
	if Countdead == s.Content.Len() || ebiten.IsKeyPressed(ebiten.KeyTab) {
		s.Content.Init()

	}

	//On génère la seed random à partir du temps pour qu'elle soit différente à chaque fois
	rand.Seed(time.Now().UnixNano())

	//On définit e, le début de la liste de particules
	e := s.Content.Front()
	//Tant que e n'est pas nul, on fait les actions suivantes
	for e != nil {
		p := (e.Value.(*Particle))
		//On fait avancer la particule
		if Extensions.Rotate {
			//On fait avancer la particule en mode rotation
			p.UpdateOrbit()

		} else {
			p.UpdatePos()
		}

		//Si le lifespan est activé, on enlève 1 à la durée de vie de la particule
		if p.Lifespan != -1 {
			p.DecreaseLife()
		}
		if !Extensions.Bounce {
			p.DeathOfParticle(s, e)
		} else {
			p.Bounce()
		}

		//On passe à la particule suivante
		e = e.Next()
	}

	//On regarde si l'option SpawnAtMouse est activée
	if Extensions.SpawnAtMouse {
		setColor()
		//On génère une particule à la position de la souris si le click gauche est enfoncé
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {

			//On compte la durée du click pour gérer le fade
			count()

			//On génère SpawnPerClick particules à la position de la souris
			for i := 0; i < config.General.SpawnPerClick; i++ {
				if col < 1 || !Extensions.Fade {
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

				s.Content.PushFront(CreateParticule())
				spawnrate--
			}
		}
	}
}
func UpdateLenList(i int) {
	X += i
}

// Le compteur de durée du click pour le fade
func count() {
	col += 0.01
}

// Pour le debug, on affiche le nombre de particules
func GetLen() int {
	return X
}

// Pour la mise en mémoire du reste du spawnrate
func spawnrateadd() {
	spawnrate += config.General.SpawnRate
}
