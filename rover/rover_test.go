package rover

import (
	
	"testing"
)


func TestRoverPosition(t *testing.T){

	p:=Plateau{
		BndriesX: 5,
		BndriesY: 5,
		}

		rover:= NewRover(1,2,'N',p)

	if rover.X!= 1 || rover.Y!= 2 || rover.Direction!= 'N' {
			t.Errorf(" got (%d,%d,%q) and want 1,2,'N' ",rover.X,rover.Y,rover.Direction)
	}
	

}

func TestTurnLeft(t *testing.T){

	p:=Plateau{
			BndriesX: 5,
			BndriesY: 5,
		}

		rover:= NewRover(0,0,'N',p)
		rover.TurnLeft()

		if rover.Direction!= 'W'{
			t.Errorf("turn left is %q want 'W'",rover.Direction) 

		}

		
}

func TestTurnRight(t *testing.T){

	p:=Plateau{
			BndriesX: 5,
			BndriesY: 5,
		}

		rover:= NewRover(0,0,'N',p)
		rover.TurnRight()


		if rover.Direction!= 'E'{
			t.Errorf("turn left is %q want 'E'",rover.Direction) 

		}
		
}

func TestMove(t *testing.T){
	p:=Plateau{
			BndriesX: 5,
			BndriesY: 5,
		}

		rover:= NewRover(1,2,'S',p)
		rover.Move()

		if rover.X!=1 || rover.Y!=1 || rover.Direction!='S'{  

			t.Errorf("move is (%d,%d,%q) want 1,1,'S'", rover.X, rover.Y, rover.Direction)

		}
}


func TestFollowInstr(t *testing.T){
	p:=Plateau{
			BndriesX: 5,
			BndriesY: 5,
		}

		rover:= NewRover(1,2,'N',p)
		rover.FollowInstr("LMLMLMLMM")

		if rover.X!=1 || rover.Y!=3 || rover.Direction!='N'{      

			t.Errorf("move is (%d,%d,%q) want 1,3,'N'", rover.X, rover.Y, rover.Direction)

		}

}