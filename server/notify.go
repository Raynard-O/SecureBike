package server

import (
	"fmt"
)

func Gyroscope(x,y,z int) bool {
	var threshold = 15	//degree
	if x > threshold || y > threshold || z > threshold{
		return true
	}
	return false
}

func Accelerometer(x, y, z int) bool {
	var threshold = 4	//meter/sec
	if x > threshold || y > threshold || z > threshold{
		return true
	}
	return false
}
//
func CheckSecurity(input [][]int) error {
	//A := make(chan int, 3)
	//B := make(chan int, 3)

	//loop:
	for {

		t := true
		fmt.Println("**************")
		//A <-rand.Intn(16)
		//A <-rand.Intn(16)
		//A <-rand.Intn(16)
		//B <-rand.Intn(5)
		//B <-rand.Intn(5)
		//B <-rand.Intn(5)
		//wg.Add(1)
		//go Notify(threshold, A, B, wg)
		fmt.Println("Reading Sensor...")
		//(input[1][0], input[1][1], input[1][2])
		if !Accelerometer(input[1][0], input[1][1], input[1][2]) {
			//threshold <- true
			t = false
			ActivateStolen("Displacement")
		}
		//s1 := [][]int{{1,3,1}}

		if !Gyroscope(input[0][0], input[0][1], input[0][2])  {
			t = false
			ActivateStolen("Angular Velocity")
		}
		if !t {
			break
		}
		fmt.Println("system secure")
	}
	return nil
}
func ActivateStolen(input string)  {
	SendMessage(input)
}

func SendMessage(input string)  {
	//Notify Owner
	fmt.Printf("Notify  owner of theft, System detected change in %v\n", input)
}