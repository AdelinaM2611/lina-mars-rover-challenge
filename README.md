# Mars Rover Challenge 

Technical test solution by **Lina Izidro**

---

## Overview

This Go program simulates a rover exploring a rectangular plateau on Mars.

The rover can:
- Turn left (`L`)
- Turn right (`R`)
- Move forward (`M`)

The plateau bounds are defined by:
- `Plateau.BndriesX` → maximum allowed X
- `Plateau.BndriesY` → maximum allowed Y  
The lower-left corner is assumed to be `(0,0)`.

The rover tracks its state as `(X, Y, Direction)` and never moves outside these bounds.

---

##  Requirements

- Go 1.21+ (or a recent Go version)

---

##  Run & Test

Clone and run tests:
```bash
git clone https://github.com/AdelinaM2611/lina-mars-rover-challenge.git
cd lina-mars-rover-challenge
go test ./...

##  Project Structure

├── go.mod
├── main.go                 # Example: builds a rover, runs instructions, prints result
└── rover/
    ├── rover.go            # Plateau, Rover, NewRover, TurnLeft/Right, Move, FollowInstr
    └── rover_test.go       # Unit tests for position, turning, move, and sequences

##  Example

**Plateau:** `BndriesX=5`, `BndriesY=5`  
**Start:** `X=1`, `Y=2`, `Direction='N'`  
**Instructions:** `LMLMLMLMM`

**Expected final state:** `X=1`, `Y=3`, `Direction='N'`

`main.go` runs this scenario.

##  Tests Included

- **Position setup** (`TestRoverPosition`)
- **Turning left/right** (`TestTurnLeft`, `TestTurnRight`)
- **Move with bounds** (`TestMove`)
- **Instruction sequence** (`TestFollowInstr`)
