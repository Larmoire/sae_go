package main

import (
	"fmt"
	"image/color"
	Extensions "project-particles/Extension"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particles"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Draw se charge d'afficher à l'écran l'état actuel du système de particules
// g.system. Elle est appelée automatiquement environ 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction pourra être légèrement modifiée quand
// c'est précisé dans le sujet.

//fonction présente : Draw, DrawGravity, DrawBounce, DrawRandomSpeed, DrawRGB, DrawSpawnAtMouse, DrawSpeedFix

func (g *game) Draw(screen *ebiten.Image) {
	for e := g.system.Content.Front(); e != nil; e = e.Next() {
		p, ok := e.Value.(*particles.Particle)
		//Si la particule est toujours en vie, ou que le Lifespan est désactivé, on l'affiche
		if p.Lifespan > 0 || p.Lifespan == -1 {
			if ok {
				options := ebiten.DrawImageOptions{}
				options.GeoM.Rotate(p.Rotation)
				options.GeoM.Scale(p.ScaleX, p.ScaleY)
				options.GeoM.Translate(p.PositionX, p.PositionY)
				options.ColorM.Scale(p.ColorRed, p.ColorGreen, p.ColorBlue, p.Opacity)
				screen.DrawImage(assets.ParticleImage, &options)
			}
			//Si la particule est en dehors de l'écran et que l'optimisation est active, on la supprime du système
			if particles.Outwindows(p) && config.General.Optimisation {
				go g.system.Content.Remove(e)
				//On enlève 1 à la longeur de la liste de particules
				particles.UpdateLenList(-1)
			}
		}
	}
	//Si le GUI est actif, on affiche les boutons
	if config.General.GUI {
		DrawRandomSpeed(screen)
		DrawnRandomSpawn(screen)
		DrawSpawnAtMouse(screen)
		DrawRGB(screen)
		DrawGravity(screen)
		DrawBounce(screen)
		DrawRotate(screen)
		DrawSpeedFix(screen)
	}
	//Si le debug est actif, on affiche les fps et la longeur de la liste de particules
	if config.General.Debug {
		ebitenutil.DebugPrintAt(screen, fmt.Sprintln(ebiten.ActualTPS()), 5, config.General.WindowSizeY-40)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint(particles.X), 5, config.General.WindowSizeY-60)
	}
	//On affiche le message de reset
	ebitenutil.DebugPrintAt(screen, "Press TAB for a reset !", 5, config.General.WindowSizeY-20)
}

//Fonction qui affiche le bouton de RandomSpeed
func DrawRandomSpeed(screen *ebiten.Image) {
	//On créee une image de 110x30 pixels
	button := ebiten.NewImage(110, 30)
	//On la remplie de rouge
	button.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	//On créee un pointeur vers DrawImageOptions
	op := &ebiten.DrawImageOptions{}
	//On dessine l'image sur l'écran
	screen.DrawImage(button, op)
	//Si l'extension est activée, on affiche "RandomSpeed: On" sinon "RandomSpeed: Off"
	if Extensions.RandomSpeed {
		ebitenutil.DebugPrintAt(screen, "RandomSpeed: On", 6, 5)
	} else {
		ebitenutil.DebugPrintAt(screen, "RandomSpeed: Off", 6, 5)
	}
}

//Fonction qui affiche le bouton de RandomSpawn
func DrawnRandomSpawn(screen *ebiten.Image) {
	//On créee une image de 110x30 pixels
	button := ebiten.NewImage(110, 30)
	//On la remplie de rouge
	button.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	//On créee un pointeur vers DrawImageOptions
	op := &ebiten.DrawImageOptions{}
	//On décale l'image de 40 pixels en y
	op.GeoM.Translate(0, 40)
	//On dessine l'image sur l'écran
	screen.DrawImage(button, op)
	//Si l'extension est activée, on affiche "RandomSpawn: On" sinon "RandomSpawn: Off"
	if Extensions.RandomSpawn {
		ebitenutil.DebugPrintAt(screen, "RandomSpawn: On", 6, 45)
	} else {
		ebitenutil.DebugPrintAt(screen, "RandomSpawn: Off", 6, 45)
	}
}

//Fonction qui affiche le bouton de SpawnAtMouse
func DrawSpawnAtMouse(screen *ebiten.Image) {
	//On créee une image de 110x30 pixels
	button := ebiten.NewImage(110, 30)
	//On la remplie de rouge
	button.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	//On créee un pointeur vers DrawImageOptions
	op := &ebiten.DrawImageOptions{}
	//On décale l'image de 80 pixels en y
	op.GeoM.Translate(0, 80)
	//On dessine l'image sur l'écran
	screen.DrawImage(button, op)
	//Si l'extension est activée, on affiche "SpawnAtMouse: On"
	if Extensions.SpawnAtMouse {
		ebitenutil.DebugPrintAt(screen, "SpawnAtMouse: On", 6, 85)
		//On décale l'image de 120 pixels en x
		op.GeoM.Translate(120, 0)
		//On dessine l'image sur l'écran
		screen.DrawImage(button, op)
		//Si l'extension Fade est activée, on affiche "Fade: On" sinon "Fade: Off"
		if Extensions.Fade {
			ebitenutil.DebugPrintAt(screen, "Fade: On", 126, 85)
		} else {
			ebitenutil.DebugPrintAt(screen, "Fade: Off", 126, 85)
		}
		//Si l'extension est désactivée, on affiche "SpawnAtMouse: Off"
	} else {
		ebitenutil.DebugPrintAt(screen, "SpawnAtMouse: Off", 6, 85)
	}
}

//Fonction qui affiche le bouton de RGB
func DrawRGB(screen *ebiten.Image) {
	//On créee une image de 110x30 pixels
	button := ebiten.NewImage(110, 30)
	//On la remplie de rouge
	button.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	//On créee un pointeur vers DrawImageOptions
	op := &ebiten.DrawImageOptions{}
	//On décale l'image de 120 pixels en y
	op.GeoM.Translate(0, 120)
	//On dessine l'image sur l'écran
	screen.DrawImage(button, op)
	//Si l'extension est activée, on affiche "RGB: On" sinon "RGB: Off"
	if Extensions.RGBchange {
		ebitenutil.DebugPrintAt(screen, "RGB: On", 6, 125)
		ebitenutil.DebugPrintAt(screen, "R for Red", config.General.WindowSizeX-59, 0)
		ebitenutil.DebugPrintAt(screen, "G for Green", config.General.WindowSizeX-70, 15)
		ebitenutil.DebugPrintAt(screen, "B for Blue", config.General.WindowSizeX-64, 30)
		ebitenutil.DebugPrintAt(screen, "C for Cyan", config.General.WindowSizeX-64, 45)
		ebitenutil.DebugPrintAt(screen, "Y for Yellow", config.General.WindowSizeX-75, 60)
		ebitenutil.DebugPrintAt(screen, "P for Purple", config.General.WindowSizeX-75, 75)
	} else {
		ebitenutil.DebugPrintAt(screen, "RGB: Off", 6, 125)
	}
}

//Fonction qui affiche le bouton de Gravity
func DrawGravity(screen *ebiten.Image) {
	//On créee une image de 110x30 pixels
	button := ebiten.NewImage(110, 30)
	//On la remplie de rouge
	button.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	//On créee un pointeur vers DrawImageOptions
	op := &ebiten.DrawImageOptions{}
	//On décale l'image de 160 pixels en y
	op.GeoM.Translate(0, 160)
	//On dessine l'image sur l'écran
	screen.DrawImage(button, op)
	//Si l'extension est activée, on affiche "Gravity: On" sinon "Gravity: Off"
	if Extensions.Gravity {
		ebitenutil.DebugPrintAt(screen, "Gravity: On", 6, 165)
	} else {
		ebitenutil.DebugPrintAt(screen, "Gravity: Off", 6, 165)
	}
}

//Fonction qui affiche le bouton de Bounce
func DrawBounce(screen *ebiten.Image) {
	//On créee une image de 110x30 pixels
	button := ebiten.NewImage(110, 30)
	//On la remplie de rouge
	button.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	//On créee un pointeur vers DrawImageOptions
	op := &ebiten.DrawImageOptions{}
	//On décale l'image de 200 pixels en y
	op.GeoM.Translate(0, 200)
	//On dessine l'image sur l'écran
	screen.DrawImage(button, op)
	//Si l'extension est activée, on affiche "Bounce: On"
	if Extensions.Bounce {
		ebitenutil.DebugPrintAt(screen, "Bounce: On", 6, 205)
		//On décale l'image de 120 pixels en x
		op.GeoM.Translate(120, 0)
		//On dessine l'image sur l'écran
		screen.DrawImage(button, op)
		//Si l'extension ColorBounce est activée, on affiche "Color: On" sinon "Color: Off"
		if Extensions.ColorBounce {
			ebitenutil.DebugPrintAt(screen, "Color: On", 126, 205)
		} else {
			ebitenutil.DebugPrintAt(screen, "Color: Off", 126, 205)
		}
		//Si l'extension est désactivée, on affiche "Bounce: Off"
	} else {
		ebitenutil.DebugPrintAt(screen, "Bounce: Off", 6, 205)
	}
}

//Fonction qui affiche le bouton de Rotate
func DrawRotate(screen *ebiten.Image) {
	//On créee une image de 110x30 pixels
	button := ebiten.NewImage(110, 30)
	//On la remplie de rouge
	button.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	//On créee un pointeur vers DrawImageOptions
	op := &ebiten.DrawImageOptions{}
	//On décale l'image de 240 pixels en y
	op.GeoM.Translate(0, 240)
	//On dessine l'image sur l'écran
	screen.DrawImage(button, op)
	//Si l'extension est activée, on affiche "Rotate: On" sinon "Rotate: Off"
	if Extensions.Rotate {
		ebitenutil.DebugPrintAt(screen, "Rotate: On", 6, 245)
	} else {
		ebitenutil.DebugPrintAt(screen, "Rotate: Off", 6, 245)
	}
}

//Fonction qui affiche le bouton de SpeedFix
func DrawSpeedFix(screen *ebiten.Image) {
	//On créee une image de 110x30 pixels
	button := ebiten.NewImage(110, 30)
	//On la remplie de rouge
	button.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	//On créee un pointeur vers DrawImageOptions
	op := &ebiten.DrawImageOptions{}
	//On décale l'image de 280 pixels en y
	op.GeoM.Translate(0, 280)
	//On dessine l'image sur l'écran
	screen.DrawImage(button, op)
	//Si l'extension est activée, on affiche "SpeedFix: On" sinon "SpeedFix: Off"
	if Extensions.SpeedFix {
		ebitenutil.DebugPrintAt(screen, "SpeedFix: On", 6, 285)
	} else {
		ebitenutil.DebugPrintAt(screen, "SpeedFix: Off", 6, 285)
	}
}
