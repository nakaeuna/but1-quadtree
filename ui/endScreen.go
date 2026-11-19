package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/nakaeuna/but1-quadtree/configuration"
)

/*
EndScreenUI est l'interface affiché si le joueur gagne.
*/
type EndScreenUI struct {
	// Commun aux structure d'UI
	IsActive bool
}

/*
 */
func (ui *UI) updateEndScreen() {

}

func (ui *UI) drawEndScreen(screen *ebiten.Image) {
	text.Draw(screen, "Victoire", ui.textFont, configuration.Global.ScreenWidth/2+16, configuration.Global.ScreenHeight/4, ui.textDeath)
	text.Draw(screen, "Jeu réalisé par", ui.textFont, configuration.Global.ScreenWidth/2-16, configuration.Global.ScreenHeight/2, ui.textDeath)
	text.Draw(screen, "naka_euna & Afissard", ui.textFont, configuration.Global.ScreenWidth/2-32, configuration.Global.ScreenHeight/2+16, ui.textDeath)
}
