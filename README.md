# Quadtree 2D Dungeon Explorer

> 2D tile-based game in Go (Ebitengine) featuring Quadtree spatial optimization and BSP procedural dungeon generation.

![Context](https://img.shields.io/badge/Context-BUT%201%20(S1)-1f883d?style=flat-square)
![Go 1.21](https://img.shields.io/badge/Go-1.21-00ADD8?style=flat-square&logo=go&logoColor=white)
![Ebitengine](https://img.shields.io/badge/Ebitengine-v2-E6522C?style=flat-square)

## Overview

Developed in **Go** with **Ebitengine** during BUT 1 (S1 — SAE 1.01), this 2D top-down exploration game uses a **Quadtree** data structure to efficiently load and query large tile maps, extended with **Binary Space Partitioning (BSP)** procedural dungeon generation, teleportation portals, a zoomable minimap camera, and JSON save/load support.

## Features

- **Core Quadtree engine**: tile map loading from files (`readFromFile`, `updateFromFile`), Quadtree construction (`MakeFromArray`), and viewport querying (`GetContent`).
- **Procedural dungeons (`FloorKind = 3`)**: random dungeon rooms & corridors generated via a **BSP tree**, interactive tiles/collisions, and **teleportation portals**.
- **UI & persistence**: main menu, JSON save/load system storing player position and floor state.
- **Camera zoom & debug overlay**: zoom out to use the camera as a minimap and toggle debug info.
- **Experimental (`FloorKind = 4`)**: on-the-fly infinite world generation *(work in progress)*.

## Controls

| Action | Key |
|---|---|
| Move character | Arrow keys (`↑` `↓` `←` `→`) |
| Return to menu | `Escape` |
| Quit current game | `Backspace` |
| Toggle debug overlay | `F1` |
| Camera zoom out / in | `F2` / `F3` |

## Tech Stack

- **Go 1.21**
- **Ebitengine v2** (`github.com/hajimehoshi/ebiten/v2`)
- **Donburi ECS**

## Project Structure

```text
├── assets/                             # Sprites, tile sheets & font loading
├── camera/                             # Viewport & zoom camera
├── character/                          # Player state, movement & rendering
├── cmd/                                # Application entry point & config.json
├── configuration/                      # Runtime configuration
├── dungeon/                            # BSP tree procedural dungeon generator
├── floor/                              # Map loading, collisions & unit tests
├── floor-files/                        # Sample maps & save files
├── game/                               # Main Ebitengine game loop
├── quadtree/                           # Quadtree spatial data structure
└── ui/                                 # Menus & HUD
```

## Getting Started

### Linux
Install X11 development libraries if needed:
```bash
sudo apt-get install libx11-dev xserver-xorg-dev xorg-dev
```
Build and run:
```bash
cd cmd && go build . && ./cmd
```

### Windows (PowerShell)
```powershell
cd cmd; go build .; .\cmd.exe
```

## Known Limitations

- Extreme camera zoom-out levels (`> 6`) on large maps can reduce FPS or trigger an Ebitengine rendering error.
- `FloorKind = 4` (infinite procedural expansion) is still experimental.
- Teleporters may stop working when resuming a level loaded from a save file, and rare wall-clipping edge cases can occur right after teleportation.

## Contributors

- Léa Cortet
- Sacha Chauvel