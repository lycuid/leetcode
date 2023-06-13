// https://leetcode.com/problems/design-parking-system/
package main

type ParkingSystem struct{ inner [4]int }

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{[4]int{0, big, medium, small}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.inner[carType] == 0 {
		return false
	}
	this.inner[carType]--
	return true
}

func main() {}
