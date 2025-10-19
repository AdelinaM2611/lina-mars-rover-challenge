package main

import (
	"fmt"
	"mars/rover"
)


func main(){

	p:=rover.Plateau{
			BndriesX: 5,
			BndriesY: 5,
		}
	rover:= rover.NewRover(1,2,'N',p)

	rover.FollowInstr("LMLMLMLMM")

	fmt.Println(rover)
	
	
}