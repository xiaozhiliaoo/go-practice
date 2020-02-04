package main

type MyFloat float64

func (m MyFloat) Abs() float64 {
	f := float64(m)
	if f < 0 {
		return -f
	}
	return f
}

func main() {
	f := MyFloat(-42)
	f.Abs()
}
