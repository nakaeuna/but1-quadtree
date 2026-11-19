package camera

import (
	// "github.com/Afissard/project-quadtree/configuration"
	"encoding/json"
	"log"
	"os"

	"github.com/nakaeuna/but1-quadtree/configuration"
)

// Init met en place une caméra.
func (c *Camera) Init() {
	if configuration.Global.CameraMode == Static {
		c.X = configuration.Global.ScreenCenterTileX
		c.Y = configuration.Global.ScreenCenterTileY
	}
}

func (c *Camera) Load(configurationFileName string) {
	content, err := os.ReadFile(configurationFileName)
	if err != nil {
		log.Fatal("Error while opening configuration file: ", err)
	}

	err = json.Unmarshal(content, &c)
	if err != nil {
		log.Fatal("Error while reading configuration file: ", err)
	}

	if configuration.Global.CameraMode == Static {
		c.X = configuration.Global.ScreenCenterTileX
		c.Y = configuration.Global.ScreenCenterTileY
	}
}
