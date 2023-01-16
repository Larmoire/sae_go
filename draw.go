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
		if p.Lifespan > 0 || p.Lifespan == -1 {
			if ok {
				options := ebiten.DrawImageOptions{}
				options.GeoM.Rotate(p.Rotation)
				options.GeoM.Scale(p.ScaleX, p.ScaleY)
				options.GeoM.Translate(p.PositionX, p.PositionY)
				options.ColorM.Scale(p.ColorRed, p.ColorGreen, p.ColorBlue, p.Opacity)
				screen.DrawImage(assets.ParticleImage, &options)
			}
			if particles.Outwindows(p) {
				if config.General.Optimisation {
					go g.system.Content.Remove(e)
					particles.UpdateLenList(-1)
				}
			}
		}
	}
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
	if config.General.Debug {
		//Si le debug est actif, on affiche les fps et la longeur de la liste de particules
		ebitenutil.DebugPrintAt(screen, fmt.Sprintln(ebiten.ActualTPS()), 5, config.General.WindowSizeY-40)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintln(particles.GetLen()), 5, config.General.WindowSizeY-60)
	}
	ebitenutil.DebugPrintAt(screen, "Press TAB for a reset !", 5, config.General.WindowSizeY-20)
}
func DrawRandomSpeed(screen *ebiten.Image) {

	button := ebiten.NewImage(110, 30)
	button.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(button, op)

	if Extensions.RandomSpeed {
		ebitenutil.DebugPrintAt(screen, "RandomSpeed: On", 6, 5)
	} else {
		ebitenutil.DebugPrintAt(screen, "RandomSpeed: Off", 6, 5)
	}
}

func DrawnRandomSpawn(screen *ebiten.Image) {
	button := ebiten.NewImage(110, 30)
	button.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 40)
	screen.DrawImage(button, op)

	if Extensions.RandomSpawn {
		ebitenutil.DebugPrintAt(screen, "RandomSpawn: On", 6, 45)
	} else {
		ebitenutil.DebugPrintAt(screen, "RandomSpawn: Off", 6, 45)
	}
}

func DrawSpawnAtMouse(screen *ebiten.Image) {

	button := ebiten.NewImage(110, 30)
	button.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 80)
	screen.DrawImage(button, op)

	if Extensions.SpawnAtMouse {
		ebitenutil.DebugPrintAt(screen, "SpawnAtMouse: On", 6, 85)
		if Extensions.Fade {
			op.GeoM.Translate(120, 0)
			screen.DrawImage(button, op)
			ebitenutil.DebugPrintAt(screen, "Fade: On", 126, 85)
		} else {
			op.GeoM.Translate(120, 0)
			screen.DrawImage(button, op)
			ebitenutil.DebugPrintAt(screen, "Fade: Off", 126, 85)
		}
	} else {
		op.GeoM.Translate(0, 80)
		screen.DrawImage(button, op)
		ebitenutil.DebugPrintAt(screen, "SpawnAtMouse: Off", 6, 85)
	}
}

func DrawRGB(screen *ebiten.Image) {

	button := ebiten.NewImage(110, 30)
	button.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	op := &ebiten.DrawImageOptions{}

	if Extensions.RGBchange {
		op.GeoM.Translate(0, 120)
		screen.DrawImage(button, op)
		ebitenutil.DebugPrintAt(screen, "RGB: On", 6, 125)
	} else {
		op.GeoM.Translate(0, 120)
		screen.DrawImage(button, op)
		ebitenutil.DebugPrintAt(screen, "RGB: Off", 6, 125)
	}
}
func DrawGravity(screen *ebiten.Image) {
	button := ebiten.NewImage(110, 30)
	button.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	op := &ebiten.DrawImageOptions{}

	if Extensions.Gravity {
		op.GeoM.Translate(0, 160)
		screen.DrawImage(button, op)
		ebitenutil.DebugPrintAt(screen, "Gravity: On", 6, 165)

	} else {
		op.GeoM.Translate(0, 160)
		screen.DrawImage(button, op)
		ebitenutil.DebugPrintAt(screen, "Gravity: Off", 6, 165)
	}
}
func DrawBounce(screen *ebiten.Image) {
	button := ebiten.NewImage(110, 30)
	button.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	op := &ebiten.DrawImageOptions{}

	if Extensions.Bounce {
		op.GeoM.Translate(0, 200)
		screen.DrawImage(button, op)
		ebitenutil.DebugPrintAt(screen, "Bounce: On", 6, 205)
		if Extensions.ColorBounce {
			op.GeoM.Translate(120, 0)
			screen.DrawImage(button, op)
			ebitenutil.DebugPrintAt(screen, "Color: On", 126, 205)
		} else {
			op.GeoM.Translate(120, 0)
			screen.DrawImage(button, op)
			ebitenutil.DebugPrintAt(screen, "Color: Off", 126, 205)
		}
	} else {
		op.GeoM.Translate(0, 200)
		screen.DrawImage(button, op)
		ebitenutil.DebugPrintAt(screen, "Bounce: Off", 6, 205)
	}
}

func DrawRotate(screen *ebiten.Image) {
	button := ebiten.NewImage(110, 30)
	button.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 240)
	screen.DrawImage(button, op)

	if Extensions.Rotate {
		ebitenutil.DebugPrintAt(screen, "Rotate: On", 6, 245)
	} else {
		ebitenutil.DebugPrintAt(screen, "Rotate: Off", 6, 245)
	}
}

func DrawSpeedFix(screen *ebiten.Image) {

	button := ebiten.NewImage(110, 30)
	button.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	op := &ebiten.DrawImageOptions{}

	if Extensions.SpeedFix {
		op.GeoM.Translate(0, 280)
		screen.DrawImage(button, op)
		ebitenutil.DebugPrintAt(screen, "SpeedFix: On", 6, 285)
	} else {
		op.GeoM.Translate(0, 280)
		screen.DrawImage(button, op)
		ebitenutil.DebugPrintAt(screen, "SpeedFix: Off", 6, 285)
	}
}
