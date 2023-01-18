package main

import (
	"log"
	"os"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particles"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/urfave/cli"
)

// main est la fonction principale du projet. Elle commence par lire le fichier
// de configuration, puis elle charge en mémoire l'image d'une particule. Elle
// initialise ensuite la fenêtre d'affichage, puis elle crée un système de
// particules encapsulé dans un "game" et appelle la fonction RunGame qui se
// charge de faire les mise-à-jour (Update) et affichages (Draw) de manière
// régulière.

var app = cli.NewApp()

func info() {
	app.Name = "Particules"
	app.Author = "Julien Evrard, Marine Rouzic"
	app.Version = "2.0"
	app.Usage = ""
	app.UsageText = "A CLI Application to generate particles"
	app.HideHelp, app.HideVersion = true, true
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "Custom",
			Aliases: []string{"c"},
			Usage:   "Use a custom with GUI config file\n",
			Action: func(c *cli.Context) {
				ExecPart("./Extension/Custom.json")
			},
		},
		{
			Name:    "Draw",
			Aliases: []string{"d"},
			Usage:   "Use a Draw config file\n",
			Action: func(c *cli.Context) {
				ExecPart("./Extension/Draw.json")
			},
		},
		{
			Name:    "Rotate",
			Aliases: []string{"r"},
			Usage:   "You can Spawn particles by your left clic, the longer the clic, the darker is the particle\n",
			Action: func(c *cli.Context) {
				ExecPart("./Extension/Rotate.json")
			},
		},
		{
			Name:    "Arrows",
			Aliases: []string{"a"},
			Usage:   "You can move the spawn point with your arrows keys !\n",
			Action: func(c *cli.Context) {
				ExecPart("./Extension/Arrows.json")
			},
		},
		{
			Name:    "Bounce",
			Aliases: []string{"b"},
			Usage:   "Create huge particle by a clic, that will bounce and change color\n",
			Action: func(c *cli.Context) {
				ExecPart("./Extension/Bouncing.json")
			},
		},
		{
			Name:    "Gravity",
			Aliases: []string{"g"},
			Usage:   "Create a particle by your left clic, affected by a gravity\n",
			Action: func(c *cli.Context) {
				ExecPart("./Extension/Gravity.json")
			},
		},
	}
}

func ExecPart(arg string) {
	config.Get(arg)
	assets.Get()
	ebiten.SetWindowTitle(config.General.WindowTitle)
	ebiten.SetWindowSize(config.General.WindowSizeX, config.General.WindowSizeY)
	//Ajout de la possibilité de resize la fenêtre
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	g := game{system: particles.NewSystem()}

	err := ebiten.RunGame(&g)
	if err != nil {
		log.Print(err)
	}
}

func main() {
	info()
	commands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
