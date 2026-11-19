package character

import (
	"encoding/json"
	"log"
	"os"

	"github.com/nakaeuna/but1-quadtree/camera"
	"github.com/nakaeuna/but1-quadtree/configuration"
)

// Init met en place un personnage. Pour le moment
// cela consiste simplement à initialiser une variable
// responsable de définir l'étape d'animation courante.
func (c *Character) Init() {
	c.AnimationStep = 1

	if configuration.Global.CameraMode == camera.Static {
		c.X = configuration.Global.ScreenCenterTileX
		c.Y = configuration.Global.ScreenCenterTileY
	}
}

func (c *Character) Load(configurationFileName string) {
	content, err := os.ReadFile(configurationFileName)
	if err != nil {
		log.Fatal("Error while opening configuration file: ", err)
	}

	err = json.Unmarshal(content, &c)
	if err != nil {
		log.Fatal("Error while reading configuration file: ", err)
	}

	if configuration.Global.CameraMode == camera.Static && (configuration.Global.ScreenCenterTileX*2 < c.X || configuration.Global.ScreenCenterTileY*2 < c.Y) {
		c.X = configuration.Global.ScreenCenterTileX
		c.Y = configuration.Global.ScreenCenterTileY
	}

	c.AnimationStep = 1
}
