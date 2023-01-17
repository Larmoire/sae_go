package Extensions

import (
	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	Gravity       bool = config.General.Gravity
	Bounce        bool = config.General.Bounce
	ColorBounce   bool = config.General.ColorBounce
	RandomSpeed   bool = config.General.RandomSpeed
	RGBchange     bool = config.General.RGBchange
	SpeedFix      bool = config.General.SpeedFix
	SpawnAtMouse  bool = config.General.SpawnAtMouse
	Rotate        bool = config.General.Rotate
	Fade          bool = config.General.Fade
	RandomSpawn   bool = config.General.RandomSpawn
	buttonPressed bool
)

func GUI() {
	//On regarde si la souris est cliquée et si l'extension GUI est activée
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && config.General.GUI {
		GravityState()
		BounceState()
		ColorBounceState()
		RandomSpeedState()
		RGBChangeState()
		SpawnAtMouseState()
		SpeedFixState()
		RotateChangeState()
		FadeState()
		RandomSpawnState()
	} else {
		//On reset le clic de la souris
		buttonPressed = false
	}
}

//Si la souris est sur le bouton, on change l'état de l'extension
func RandomSpeedState() {
	x, y := ebiten.CursorPosition()
	if x >= 0 && x <= 110 && y >= 0 && y <= 30 {
		if !buttonPressed {
			buttonPressed = true
			RandomSpeed = !RandomSpeed
		}
	}
}

//Si la souris est sur le bouton, on change l'état de l'extension
func RandomSpawnState() {
	x, y := ebiten.CursorPosition()
	if x >= 0 && x <= 110 && y >= 40 && y <= 70 {
		if !buttonPressed {
			buttonPressed = true
			RandomSpawn = !RandomSpawn
		}
	}
	//Si SpawnAtMouse est activé et le RandomSpawn aussi, on le désactive
	if SpawnAtMouse && RandomSpawn {
		SpawnAtMouse = !SpawnAtMouse
	}
}

//Si la souris est sur le bouton, on change l'état de l'extension
func SpawnAtMouseState() {
	x, y := ebiten.CursorPosition()
	if x >= 0 && x <= 110 && y >= 80 && y <= 110 {
		if !buttonPressed {
			buttonPressed = true
			SpawnAtMouse = !SpawnAtMouse
		}
	}
	if !SpawnAtMouse {
		Fade = false
	}
	//Si SpawnAtMouse est activé et le RandomSpawn aussi, on désactive RandomSpawn
	if SpawnAtMouse && RandomSpawn {
		RandomSpawn = false
	}
}

//Si la souris est sur le bouton, on change l'état de l'extension
func FadeState() {
	x, y := ebiten.CursorPosition()
	if x >= 135 && x <= 245 && y >= 80 && y <= 110 && SpawnAtMouse {
		if !buttonPressed {
			buttonPressed = true
			Fade = !Fade
		}
	}
}

//Si la souris est sur le bouton, on change l'état de l'extension
func RGBChangeState() {
	x, y := ebiten.CursorPosition()
	if x >= 0 && x <= 110 && y >= 120 && y <= 150 {
		if !buttonPressed {
			buttonPressed = true
			RGBchange = !RGBchange

		}
	}
}

//Si la souris est sur le bouton, on change l'état de l'extension
func GravityState() {
	x, y := ebiten.CursorPosition()
	if x >= 0 && x <= 110 && y >= 160 && y <= 190 {
		if !buttonPressed {
			buttonPressed = true
			Gravity = !Gravity
		}
	}
}

//Si la souris est sur le bouton, on change l'état de l'extension
func BounceState() {
	x, y := ebiten.CursorPosition()
	if x >= 0 && x <= 110 && y >= 200 && y <= 230 {
		if !buttonPressed {
			buttonPressed = true
			Bounce = !Bounce
		}
	}
	//Si Bounce est désactivé, on désactive ColorBounce
	if !Bounce {
		ColorBounce = false
	}
}

//Si la souris est sur le bouton, on change l'état de l'extension
func ColorBounceState() {
	x, y := ebiten.CursorPosition()
	if x >= 135 && x <= 245 && y >= 200 && y <= 230 && Bounce {
		if !buttonPressed {
			buttonPressed = true
			ColorBounce = !ColorBounce
		}
	}
}

//Si la souris est sur le bouton, on change l'état de l'extension
func RotateChangeState() {
	x, y := ebiten.CursorPosition()
	if x >= 0 && x <= 110 && y >= 240 && y <= 270 {
		if !buttonPressed {
			buttonPressed = true
			Rotate = !Rotate

		}
	}
}

//Si la souris est sur le bouton, on change l'état de l'extension
func SpeedFixState() {
	x, y := ebiten.CursorPosition()
	if x >= 0 && x <= 110 && y >= 280 && y <= 310 {
		if !buttonPressed {
			buttonPressed = true
			SpeedFix = !SpeedFix
		}
	}
}
