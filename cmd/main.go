package main

import (
	"flag"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	// "github.com/Afissard/project-quadtree/configuration"
	// "github.com/Afissard/project-quadtree/game"
	// "github.com/Afissard/project-quadtree/assets"
	"github.com/nakaeuna/but1-quadtree/assets"
	"github.com/nakaeuna/but1-quadtree/configuration"
	"github.com/nakaeuna/but1-quadtree/game"
)

func main() {
	var configFileName string
	flag.StringVar(&configFileName, "config", "config.json", "select configuration file")

	configuration.Load(configFileName)

	assets.Load()

	g := &game.Game{}
	g.Init()

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	ebiten.SetWindowTitle("Quadtree")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}
