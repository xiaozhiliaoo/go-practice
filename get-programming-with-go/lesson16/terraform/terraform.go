package main

import "fmt"

// terraform accomplishes nothing
func terraform(planets [8]string) {
	for i := range planets {
		planets[i] = "New " + planets[i]
	}
}

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	//如果改成切片，会改变值
	terraform(planets)
	fmt.Println(planets)
}
