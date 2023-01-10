package main

import (
	"fmt"
	"image/color"
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
				go g.system.Content.Remove(e)
				particles.UpdateLenList(-1)
			}
		}
	}

	if config.General.Debug {
		//Si le debug est actif, on affiche les fps et la longeur de la liste de particules
		DrawGravity(screen)
		DrawBounce(screen)
		DrawRandomSpeed(screen)
		DrawRGB(screen)
		DrawSpawnAtMouse(screen)
		DrawSpeedFix(screen)

		ebitenutil.DebugPrintAt(screen, fmt.Sprintln(ebiten.ActualTPS()), config.General.WindowSizeX-108, 0)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintln(particles.GetLen()), config.General.WindowSizeX-15, 15)
		ebitenutil.DebugPrintAt(screen, "Press TAB for a reset !", 5, config.General.WindowSizeY-20)
	}

}
func DrawGravity(screen *ebiten.Image) {

	button := ebiten.NewImage(110, 30)
	button.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	op := &ebiten.DrawImageOptions{}

	if particles.GetGravityState() {
		screen.DrawImage(button, op)
		ebitenutil.DebugPrintAt(screen, "Gravity: On", 6, 5)

	} else {
		screen.DrawImage(button, op)
		ebitenutil.DebugPrintAt(screen, "Gravity: Off", 6, 5)
	}
}
func DrawBounce(screen *ebiten.Image) {

	button := ebiten.NewImage(110, 30)
	button.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	op := &ebiten.DrawImageOptions{}

	if particles.GetBounceState() {
		op.GeoM.Translate(0, 40)
		screen.DrawImage(button, op)
		ebitenutil.DebugPrintAt(screen, "Bounce: On", 6, 45)
		if particles.GetColorBounceState() {
			op.GeoM.Translate(120, 0)
			screen.DrawImage(button, op)
			ebitenutil.DebugPrintAt(screen, "Color: On", 126, 45)
		} else {
			op.GeoM.Translate(120, 0)
			screen.DrawImage(button, op)
			ebitenutil.DebugPrintAt(screen, "Color: Off", 126, 45)
		}
	} else {
		op.GeoM.Translate(0, 40)
		screen.DrawImage(button, op)
		ebitenutil.DebugPrintAt(screen, "Bounce: Off", 6, 45)
	}
}
func DrawRandomSpeed(screen *ebiten.Image) {

	button := ebiten.NewImage(110, 30)
	button.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	op := &ebiten.DrawImageOptions{}

	if particles.GetRandomSpeedState() {

		op.GeoM.Translate(0, 80)
		screen.DrawImage(button, op)
		ebitenutil.DebugPrintAt(screen, "RandomSpeed: On", 6, 85)
	} else {

		op.GeoM.Translate(0, 80)
		screen.DrawImage(button, op)
		ebitenutil.DebugPrintAt(screen, "RandomSpeed: Off", 6, 85)
	}
}
func DrawRGB(screen *ebiten.Image) {

	button := ebiten.NewImage(110, 30)
	button.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	op := &ebiten.DrawImageOptions{}

	if particles.GetRGBChangeState() {
		op.GeoM.Translate(0, 120)
		screen.DrawImage(button, op)
		ebitenutil.DebugPrintAt(screen, "RGB: On", 6, 125)
	} else {
		op.GeoM.Translate(0, 120)
		screen.DrawImage(button, op)
		ebitenutil.DebugPrintAt(screen, "RGB: Off", 6, 125)
	}
}
func DrawSpawnAtMouse(screen *ebiten.Image) {

	button := ebiten.NewImage(110, 30)
	button.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	op := &ebiten.DrawImageOptions{}

	if particles.GetSpawnAtMouseState() {
		op.GeoM.Translate(0, 160)
		screen.DrawImage(button, op)
		ebitenutil.DebugPrintAt(screen, "SpawnAtMouse: On", 6, 165)
	} else {
		op.GeoM.Translate(0, 160)
		screen.DrawImage(button, op)
		ebitenutil.DebugPrintAt(screen, "SpawnAtMouse: Off", 6, 165)
	}
}
func DrawSpeedFix(screen *ebiten.Image) {

	button := ebiten.NewImage(110, 30)
	button.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	op := &ebiten.DrawImageOptions{}

	if particles.GetSpeedFixState() {
		op.GeoM.Translate(0, 200)
		screen.DrawImage(button, op)
		ebitenutil.DebugPrintAt(screen, "SpeedFix: On", 6, 205)
	} else {
		op.GeoM.Translate(0, 200)
		screen.DrawImage(button, op)
		ebitenutil.DebugPrintAt(screen, "SpeedFix: Off", 6, 205)
	}
}
