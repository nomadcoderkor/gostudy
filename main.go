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

// func isLogin(userName string) bool {
// 	if len(userName) > 0 {
// 		return true
// 	}
// 	return false
// }

// func handyScore(over int) string {
// 	if score := over + 72; score < 82 {
// 		return "Single 골퍼"
// 	}
// 	return "휴.."
// }

func handyScore(over int) string {
	// 조건절에 변수설정을 할 수 있다.
	// score 변수 생성후 세미콜론 다음 조건을 줄 수 있다.
	// score := over + 72
	switch score := over + 72; {
	case score == 72:
		return "이븐"
	case score < 72:
		return "와.. 언더.."
	case score < 82:
		return "싱글"
	case score > 81 && score < 90:
		return "80대 고수"
	case score == 90:
		return "보기플레이어"
	case score > 90 && score < 100:
		return "흠.. 애매하네.."
	}
	return "제발 연습좀 .."
}

func main() {
	// len, upper := intAndString("고랭 좋은거 같은데")
	// fmt.Println(len, upper)
	// result := loopSum(2, 44, 5, 43, 8, 18)
	// fmt.Println(result)
	// fmt.Println(isLogin("test"))
	fmt.Println(handyScore(0))
}
