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
func NewSystem() System {
	l := list.New()
	var PosX float64
	var PosY float64

	for i := 0; i < (config.General.InitNumParticles); i++ {
		if !config.General.RandomSpawn {
			PosX = float64(config.General.SpawnX)
			PosY = float64(config.General.SpawnY)
		} else {
			PosX = float64(rand.Float64() * (float64(config.General.WindowSizeX) - 2))
			PosY = float64(rand.Float64() * (float64(config.General.WindowSizeY) - 2))
		}
		l.PushFront(&Particle{
			PositionX: PosX,
			PositionY: PosY,
			ScaleX:    1, ScaleY: 1,
			ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
			Opacity: 1,
		})
	}
	return System{Content: l}
}
