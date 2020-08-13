package main

import "fmt"

func main() {
	var d int
	var h int
	var m int
	const degreesInCircle int = 360
	const hoursInCircle int = 12
	const minutesInHour int = 60
	const degreesInHour int = degreesInCircle / hoursInCircle
	const minutesInDegree = minutesInHour / degreesInHour
	fmt.Scan(&d)
	h = d / degreesInHour
	d %= degreesInHour
	m = d * minutesInDegree
	fmt.Println("It is", h, "hours", m, "minutes.")
}

