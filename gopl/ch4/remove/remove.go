package main

import "fmt"

//删除后保证原来顺序
func remove(slice []int, i int) []int {
	fmt.Println(slice[i:], slice[i+1:])
	copy(slice[i:], slice[i+1:]) // 56899
	return slice[:len(slice)-1]  //5689
}

//删除后不保证原来顺序
func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2)) // "[5 6 8 9]"

	s2 := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s2, 2)) // "[5 6 9 8]
}
