package main

import (
	"errors"
	"fmt"
)

func main() {
	sprintf := fmt.Sprintf("调用函数引擎失败:%v", errors.New("dddddd"))
	fmt.Println(sprintf)
}
