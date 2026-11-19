package game

import (
	"github.com/nakaeuna/but1-quadtree/configuration"
	"github.com/nakaeuna/but1-quadtree/dungeon"
)

func (g *Game) newLevel() {
	g.floor.QuadtreeContent, g.character.X, g.character.Y, g.listPortal = dungeon.NewDungeon(configuration.Global.FixedGenWidth, configuration.Global.FixedGenHeight, "fromCurrentSeed")
	if configuration.Global.CameraMode == 1 {
		g.camera.X = g.character.X
		g.camera.Y = g.character.Y
	}
	g.ui.DungeonLevel.DungeonLevel++
	g.ui.DungeonLevel.IsActive = true
	g.ui.DungeonLevel.TimeShowed = 60 * 1
}
