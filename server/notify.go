package server

import (
	"fmt"
	"math/rand"
	"sync"
)

func gyroscope(x,y,z int) bool {
	var threshold = 15	//degree
	if x > threshold || y > threshold || z > threshold{
		return true
	}
	return false
}

func accelerometer(x, y, z int) bool {
	var threshold = 4	//meter/sec
	if x > threshold || y > threshold || z > threshold{
		return true
	}
	return false
}
//
func NotifyU(threshold chan bool, wg *sync.WaitGroup)  {
	A := make(chan int, 3)
	B := make(chan int, 3)

//loop:
	for {
		var i int
		i++
		fmt.Printf("********%v\n******", i)
		A <-rand.Intn(19)
		A <-rand.Intn(19)
		A <-rand.Intn(19)
		B <-rand.Intn(5)
		B <-rand.Intn(5)
		B <-rand.Intn(5)
		wg.Add(1)
		go Notify(threshold, A, B, wg)
		//if !(<-threshold) {
		//	wg.Done()
		//}

		//if <- threshold{
		//	goto loop
		//}

	}

	//<- threshold
}

func Notify( threshold chan bool, A chan int, B chan int, wg *sync.WaitGroup )  {
	//threshold := make(chan bool)
	//var wg sync.WaitGroup
	//wg.Add(1)
	//go func() {
	fmt.Println("Reading Sensor...")

	//fmt.Printf("Acceleromter x : %v \n", <-A)
	//fmt.Printf("Acceleromter y : %v \n", <-A)
	//fmt.Printf("Acceleromter z : %v \n", <-A)
	//fmt.Printf("Gyroscope x : %v \n", <-B)
	//fmt.Printf("Gyroscope y : %v \n", <-B)
	//fmt.Printf("Gyroscope z : %v \n",<-B)
	if accelerometer(<-B, <-B, <-B){
		threshold <- true
		fmt.Println("Notify owner of error")
	}
	if gyroscope(<-A, <-A, <-A){
		threshold <- true
		fmt.Println("Notify owner of error")
	}
	//time.Sleep(time.Second)

	wg.Done()
}