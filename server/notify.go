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
func CheckSecurity(input [][]int, threshold chan string ) error {
	//A := make(chan int, 3)
	//B := make(chan int, 3)

	//loop:
	for i := 0; i < 10; i++ {

		t := true
		fmt.Println("**************")

		fmt.Println("Reading Sensor...")
		//(input[1][0], input[1][1], input[1][2])
		if !Accelerometer(input[1][0], input[1][1], input[1][2]) {
			//threshold <- true
			t = false
			threshold <-ActivateStolen("Displacement")
			//fmt.Println(ActivateStolen("Displacement"))

		}
		//s1 := [][]int{{1,3,1}}

		if !Gyroscope(input[0][0], input[0][1], input[0][2])  {
			t = false
			threshold <- ActivateStolen("Angular Velocity")
			//fmt.Println(ActivateStolen("Angular Velocity"))
		}
		if !t {
			break
		}
		threshold <- ActivateStolen("system secured")
	}
	return nil
}
func ActivateStolen(input string) string {
	return SendMessage(input)
}

func SendMessage(input string) string {
	//Notify Owner
	s := fmt.Sprintf("Notify  owner of theft, System detected change in %v\n", input)
	return s
}