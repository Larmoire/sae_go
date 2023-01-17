package particles

import (
	"container/list"
	"project-particles/config"
)

//Fonction qui met la particule à mort si elle sort de la fenêtre ou si son temps de vie est écoulé
func (p *Particle) DeathOfParticle(s *System, e *list.Element) {
	//Si le lifespan est à 0 ou si la particule sort de la fenêtre on met son opacité à 0
	if p.Lifespan == 0 || Outwindows(p) && p.Lifespan != -1 {
		p.Opacity = 0
		//Si l'option Optimisation est activée, on enlève la particule de la liste des particules
		if config.General.Optimisation && s != nil {
			go s.Content.Remove(e)
			//On incrémente le compteur de particules mortes
			Countdead += 1
		}
	}
}
func Outwindows(p *Particle) bool {
	//On vérifie si la particule sort de la fenêtre en X et en Y
	return (p.PositionX >= float64(config.General.WindowSizeX) || p.PositionX <= 0-10*config.General.ScaleX || p.PositionY >= float64(config.General.WindowSizeY) || p.PositionY <= 0-10*config.General.ScaleY)
}
