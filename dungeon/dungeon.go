package dungeon

import (
	"math/rand"

	"github.com/nakaeuna/but1-quadtree/quadtree"
)

/*
Structure contenant le donjon :
- Arbre BSP
- Quadtree du donjon
*/
type Dungeon struct {
	bspTree bspTree
	quad    quadtree.Quadtree
}

/*
Racine de l'arbre BSP : point de départ de la génération BSP
*/
type bspTree struct {
	width, height int
	topX, topY    int // par défaut (0,0) ?
	root          *bspNode
}

/*
Structure d'une node pour l'arbre BSP du donjon
*/
type bspNode struct {
	width, height int
	topX, topY    int
	nodeLeft      *bspNode
	nodeRight     *bspNode
}

/*
Structure d'un portail, chaque portail est lié à un autre portail qui sert point de sortie.
*/
type Portal struct {
	PosX, PosY int 
	ExitX, ExitY int
	LinkedPortail *Portal
}

var (
	alea *rand.Rand // seed du donjon
)
