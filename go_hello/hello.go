package main // package name

import (
	"fmt"
	_ "net/http"
)

func main() {
	var name = "LL" // 不想給值的話 var name string
	//i := 100        //省略 var 的宣告方法
	const i2 string = "const 宣告的變數其值不可變動"
	//name = name + "a"
	//slice := [...]string{1: "RTA"} // 不用宣告長度的 array
	//array := [10]int {0}// 長度為10的 array

	age := make(map[string]int) //就是 HASH
	age["RTA"] = 20
	age["A4"] = 10
	age["idk"] = 0
	delete(age, "idk") // map內建的刪除方式

	fmt.Println(name)
	fmt.Println("AYAYA")
	var c int = 0
	fmt.Println(aaa(c))
}

func aaa(b int) (RTA string) { //func aaa(傳入參數)(returm參數型態)
	var a int = 5
	if a > 1 {
		RTA := "RTA"
		return RTA
	}
	return "none"
}

func sum(a, b int) int {
	return a + b
}
