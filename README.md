# Quadtree 2D Dungeon Explorer

> 2D tile-based game in Go and Ebitengine featuring Quadtree spatial optimization and BSP procedural dungeon generation.

![Context](https://img.shields.io/badge/Context-BUT%201-1f883d?style=flat-square)
![Go 1.21](https://img.shields.io/badge/Go-1.21-00ADD8?style=flat-square&logo=go&logoColor=white)
![Ebitengine](https://img.shields.io/badge/Ebitengine-v2-E6522C?style=flat-square)

## Overview

2D top-down exploration game built in **Go** with **Ebitengine**. It uses a **Quadtree** data structure to load and query large tile maps efficiently, along with **Binary Space Partitioning (BSP)** procedural dungeon generation, teleportation portals, a zoomable minimap camera, and JSON save and load support.

## Features

- **Core Quadtree engine:** Loads tile maps from files, builds a Quadtree (`MakeFromArray`), and queries only visible viewport tiles (`GetContent`).
- **Procedural dungeons (`FloorKind = 3`):** Generates random dungeon rooms and corridors with a BSP tree, interactive tiles, collisions, and teleportation portals.
- **Interface and persistence:** Main menu and JSON save and load system for player position and floor state.
- **Camera zoom and debug overlay:** Zoom out to view the map as a minimap and toggle debug information.
- **Experimental (`FloorKind = 4`):** On-the-fly infinite world generation (work in progress).

## Controls

| Action | Key |
|---|---|
| Move character | Arrow keys (`↑` `↓` `←` `→`) |
| Return to menu | `Escape` |
| Quit current game | `Backspace` |
| Toggle debug overlay | `F1` |
| Camera zoom out / in | `F2` / `F3` |

## Technology Stack

- **Go 1.21**
- **Ebitengine v2** (`github.com/hajimehoshi/ebiten/v2`)
- **Donburi ECS**

## Project Structure

```text
├── assets/                             # Sprites, tile sheets, and font loading
├── camera/                             # Viewport and zoom camera
├── character/                          # Player state, movement, and rendering
├── cmd/                                # Application entry point and configuration
├── configuration/                      # Runtime configuration
├── dungeon/                            # BSP tree procedural dungeon generator
├── floor/                              # Map loading, collisions, and unit tests
├── floor-files/                        # Sample maps and save files
├── game/                               # Main Ebitengine game loop
├── quadtree/                           # Quadtree spatial data structure
└── ui/                                 # Menus and heads-up display
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

- Extreme camera zoom-out levels (`> 6`) on large maps can reduce frame rate or cause an Ebitengine rendering error.
- `FloorKind = 4` (infinite procedural expansion) is still experimental.
- Teleporters may stop working after loading a saved floor, and rare wall-clipping cases can occur right after teleportation.

## License

This project is distributed under the **GNU General Public License v3.0** (inherited from the original starter codebase). See [`LICENSE`](LICENSE) for details.