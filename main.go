package main

import (
	"fmt"
)

// func intAndString(name string) (nameLen int, upperName string) {
// 	defer fmt.Println("펑션종료")
// 	nameLen = len(name)
// 	upperName = strings.ToUpper(name)
// 	fmt.Println("리턴값 셋팅 완료 리턴 직전")
// 	return
// }

// func loopSum(numbers ...int) int {
// 	result := 0
// 	for _, item := range numbers {
// 		result += item
// 	}
// for i := 0; i < count; i++ {
//   result += item
// }
//   return result
// }

func isLogin(userName string) bool {
	if len(userName) > 0 {
		return true
	}
	return false
}

func main() {
	// len, upper := intAndString("고랭 좋은거 같은데")
	// fmt.Println(len, upper)
	// result := loopSum(2, 44, 5, 43, 8, 18)
	// fmt.Println(result)
	fmt.Println(isLogin(""))
}
