package server

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestCheckSecurity(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	var err error
	//threshold := make(chan bool, 1)
	input := [][]int{
		{1,3,3},
		{2,3,1},
	}
	go func() {
		err = CheckSecurity(input)
		wg.Done()
	}()
	wg.Wait()
	assert.NoError(t, err)
	//<- threshold
}

func TestAccelerometer(t *testing.T) {
	s := [][]int{{1,5,1}, {5,1,1}, {1,1,5}}
	for _, j := range s{

			b := Accelerometer(j[0], j[1], j[2])
			fmt.Println(
				b)
			assert.Equal(t, true, b, fmt.Sprintf("%v %v %v",j[0], j[1], j[2]))
	}
	s1 := [][]int{{1,3,1}}
	for _, j := range s1{
		//for h,j := range v{
		//	fmt.Println(h, j)
		b := Accelerometer(j[0], j[1], j[2])
		fmt.Println(
			b)
		assert.Equal(t, false, b, fmt.Sprintf("%v %v %v",j[0], j[1], j[2]))
	}

}

func TestGyroscope(t *testing.T) {
	s := [][]int{{1,16,1}, {16,1,1}, {1,1,16}}
	for _, j := range s{

		b := Gyroscope(j[0], j[1], j[2])
		assert.Equal(t, true, b, fmt.Sprintf("%v %v %v",j[0], j[1], j[2]))
	}
	s1 := [][]int{{1,3,1}}
	for _, j := range s1{
		//for h,j := range v{
		//	fmt.Println(h, j)
		b := Gyroscope(j[0], j[1], j[2])
		assert.Equal(t, false, b, fmt.Sprintf("%v %v %v",j[0], j[1], j[2]))
	}

}