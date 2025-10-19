package rover


//Notes to Lina
//- Rover needs to understand the boundaries of its environment
//-Rover needs to understand its position and how to turn
//- Rover needs to follow instructions and actually move

/*cartesian plane
move up north = y + 1
move down south = y - 1
move forward east = x + 1
move backword west = x - 1 */

//Platau represents the boundaries of the platau
type Plateau struct{
	BndriesX int
	BndriesY int
}

//Rover represents the rover's current position
type Rover struct{
	X int
	Y int
	Direction byte // 'N' 'W' 'S' 'E'
	Plateau Plateau
}

//2-NewRover represents a rover that understands its current position 
// and the boundaries of it's environment

func NewRover(x,y int, dir byte, p Plateau) *Rover{
	return &Rover{X: x, Y: y,Direction: dir,Plateau: p}
}

//teaching Rover how to move

func (r *Rover) TurnLeft(){
	switch r.Direction{
	case 'N':
		r.Direction= 'W'
	case 'W':
		r.Direction= 'S'
	case 'S':
		r.Direction = 'E'
	case 'E':
		r.Direction= 'N'

	}
}

func (r *Rover) TurnRight(){
	switch r.Direction{
	case 'N':
		r.Direction= 'E'
	case 'E':
		r.Direction= 'S'
	case 'S':
		r.Direction = 'W'
	case 'W':
		r.Direction= 'N'

	}
}

func (r *Rover) Move(){
	switch r.Direction{ //we are checking the condition before making the move, if it satisfies then it will move
	case 'N':
		if r.Y+1<= r.Plateau.BndriesY{ //this prevents a fall off, if there is a fall off move it will skip it
			r.Y++
		}
	case 'S':
		if r.Y-1>=0{ //the lower left corner of the grid is (0,0) by assumption
			r.Y --
		}
	case 'E':
		if r.X+1 <= r.Plateau.BndriesX{
			r.X ++
		}
	case 'W':
		if r.X-1 >= 0 {
			r.X--
		}
	}
}

//FollowInstr represents rover following instructions and moving
func(r *Rover) FollowInstr(instr string){ //LMLMLMLMM
	for i := 0; i< len(instr); i++{
		switch instr[i] {
		case 'L':
			r.TurnLeft()
		case 'R':
			r.TurnRight()
		case 'M':
			r.Move()
		}
			
	}

}