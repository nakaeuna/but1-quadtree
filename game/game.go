package game

import (
	"github.com/nakaeuna/but1-quadtree/camera"
	"github.com/nakaeuna/but1-quadtree/character"
	"github.com/nakaeuna/but1-quadtree/floor"
	"github.com/nakaeuna/but1-quadtree/ui"
	"github.com/nakaeuna/but1-quadtree/dungeon"
)

// Game est le type permettant de représenter les données du jeu.
// Aucun champs n'est exporté pour le moment.
//
// Les champs non exportés sont :
//   - camera : la représentation de la caméra
//   - floor : la représentation du terrain
//   - character : la représentation du personnage
type Game struct {
	camera    camera.Camera
	floor     floor.Floor
	character character.Character
	ui        ui.UI

	state        int
	zoomLevel    int
	needNewLevel bool
	listPortal []dungeon.Portal
}

/*
Liste des state du jeu
*/
const (
	mainMenu int = iota // commence à 0, les constantes suivantes auront la valeur d'un int + "n° déclaration"
	inGame
	gameOver
	endScreen
)

var sauvegardeCam string
var sauvegarde string
