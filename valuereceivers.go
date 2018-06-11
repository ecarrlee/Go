package main

import "fmt"
const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
gas_pedal uint16
brake_pedal uint16
steering_wheel int16
top_speed_kmh float64
}

func (c car) kmh() float64 {
	return float64(c.gas_pedal) *
		(c.top_speed_kmh/usixteenbitmax)
}

// Based on tutorial @ https://youtu.be/3fsvqo9pQyg
func main() {
	a_car := car{gas_pedal:221,
		brake_pedal:4,
		steering_wheel:234,
		top_speed_kmh:88888}

	b_car := car{221,4,234,8}

	fmt.Println(a_car.gas_pedal)
	fmt.Println(b_car.top_speed_kmh)
	car_pointer := &b_car

	/*Pointers can be accessed with 
	"." instead of "->" in go!*/
	fmt.Println(car_pointer.top_speed_kmh)
	fmt.Println(b_car.kmh())
}
