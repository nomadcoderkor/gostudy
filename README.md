# GO Lang Study

### go mod init git repository url --> mod.go 파일 생성
### main.go 파일 생성 컴파일이 필요하기 때문에 main.go 파일을 생성한다.
  1. 컴파일하고 싶다면 main 패키지에 있는 main 함수가 프로그램의 시작점이 된다. (아래와 같은 모양으로 생성된다.)
   ```go
    package main

    import "fmt"

    func main() {
	    fmt.Println("Hello, World!")
    }
   ```   
---         
  2. 컴파일러는 패키지 이름이 main인 것부터 찾는다. 즉 main package가 진입점(entry point)이 되며 main package에서 main function 순으로 찾게 된다.   
   
### 변수와 상수
```go
// 아래의 형태로 타입을 지정 하여 상수를 선언 해준다.
const name string = "이승재"
 상수  변수명  타입
 name = "하이" <-- 에러발생 
// 변수는 var을 사용한다.
var name string = "이승재"
name = "하이"

var name string = "이승재" 이렇게 작성하는게 너무 길어서 불편하다면
name := "이승재" 이 처럼 축약하여 작성도 가능하다 이렇게 작성을 하면 타입은 자동으로 맵핑이 된다.

```   
- 변수와 상수의 선언은 javascript, typescript 와 비슷하다.
- [Go Lang Type Link](https://go101.org/article/basic-types-and-value-literals.html)

### Function
```go
// javascript 두수의 합을 리턴을 해주는 펑션
function sum(a,b){
  return a+b;
}

// Go 에서의 함수
func sum(a int,    b int) int {
           변수타입  변수타입  리턴타입
  return a + b
}
// 변수와 리턴타입을 모두 넣어주어야 실행이 된다.

// Go 에서는 리턴을 여러개 줄수도 있다.
func intAndString(name string) (int, string){
  return len(name), strings.toUpper(name);
}
// 위 처럼 작성을 하면 이름의 길이와, 이름을 uppercase한 두개의 값이 return 된다.
// len과 toUpper는 Go에서 제공하는 package 이다.

// 리턴을 줄때 아래와 같이 할수 도 있다.
func intAndString(name string) (nameLen int, upperName string){
  nameLen = len(name);
  upperName = strings.ToUpper(name)
  return
  // 리턴값을 따로 줄 필요없다.
}

// 같은 타입의 여러개의 파라메터를 한번에 처리 하는 방법
// 아래 처럼 작성을 하게 되면 4개의 스트링 파라메터를 넘기기 때문에 컴파일 에러가 발생

repeatFunc("aa", "bb", "cc", "dd")

func repeatFunc(param string){
  fmt.Println(param);
}

// 아래 처럼 작성을 하면 정상 동작을 하게 된다.
func repeatFunc(param ...string){
  fmt.Println(param);
}

// defer
// defer은 펑션종료후에 실행하고픈 기능을 수행 한다.
// 우리가 알던 펑션은 리턴을 하고 난후에 종료가 되지만,
// 아래 코드에서는 리턴을 하고난 후에 defer에 기록해 놓은 출력을 하게 된다.
func intAndString(name string) (nameLen int, upperName string){
  defer fmt.Println("펑션종료")
  nameLen = len(name);
  upperName = strings.ToUpper(name)
  fmt.Println("리턴값 셋팅 완료 리턴 직전");
  return
}

```
## [더많은 Package 정보 링크](https://pkg.go.dev/std)

***

# Loop
  * Go Lang 에서 Loop는 for / for range 두가지가 전부이다.
    흠.. 이게 심플해서 좋은건지.. 아닌지는 아직 잘 모르겠다... ㅎㅎ
```go
func loopSum(numbers ...int) int {
	result := 0
	for _, item := range numbers {
		result += item
	}
	return result
}

func loopSum(numbers ...int) int {
	result := 0
	for i := 0; i < len(numbers); i++ {
    result += numbers[i]
  }
	return result
}
```

# IF ELSE
  * 


