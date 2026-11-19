package game

import (
	// "github.com/Afissard/project-quadtree/configuration"
	"github.com/nakaeuna/but1-quadtree/configuration"
)

// Layout détermine la taille de l'image sur laquelle Ebitengine
// affiche les images du jeu en fonction de la taille de la fenêtre
// dans laquelle l'affichage a lieu. Pour le moment, cette taille
// ne dépend pas de celle de la fenêtre.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	screenWidth = int(float32(configuration.Global.ScreenWidth) * 1.5) // modifie la taille de la fenêtre
	screenHeight = int(float32(configuration.Global.ScreenHeight) * 1.0)
	if configuration.Global.DebugMode {
		screenWidth += configuration.Global.NumTileForDebug * configuration.Global.TileSize
		screenHeight += configuration.Global.TileSize
	}
	return screenWidth, screenHeight
}
