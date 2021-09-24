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
```go
func isLogin(userName string) bool {
	if len(userName) > 0 {
		return true
	}
	return false
}

func handyScore(over int) string {
  // 조건절에 변수설정을 할 수 있다.
  // score 변수 생성후 세미콜론 다음 조건을 줄 수 있다.
	if score := over + 72; score < 82 {
		return "Single 골퍼"
	}
	return "휴.."
}
```
---

# Switch
```go
func handyScore(over int) string {
	// 조건절에 변수설정을 할 수 있다.
	// score 변수 생성후 세미콜론 다음 조건을 줄 수 있다.
	// score := over + 72
	switch score := over + 72; {
	case score == 72:
		return "이븐"
	case score < 72:
		return "절대 고수"
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
```

# Pointer
```go
a := 2
b := a
fmt.Println(a, b)
// 결과 => 2 2

a := 2
b := a
a = 10
fmt.Println(a, b)
// 결과 => 10 2

a := 2
b := a
a = 10
fmt.Println(&a, &b)
// 결과 => 0xc0000b2008 0xc0000b2010
// 변수 앞에 &를 붙이면 메모리 주소값을 출력해 준다.

a := 2
b := &a
fmt.Println(a, b)
// 결과 => 2 0xc0000b2008

a := 2
b := &a
a = 10
fmt.Println(a, b)
// 결과 => 10 0xc000014098

a := 2
b := &a
a = 10
fmt.Println(a, b)
// 결과 => 10 0xc000014098

a := 2
b := &a
fmt.Println(a, *b)
// 결과 => 2 2
// 변수 앞에 *을 붙여주면 주소가 참조하고 있는 값을 찾는다

a := 2
b := &a
a = 11
fmt.Println(a, *b)
// 결과 => 11 11

a := 2
b := &a
*b = 222
fmt.Println(a, *b)
// 결과 => 222 222
```

# Array

```go
// 배열 선언
arr := [5]string{"user01","user02","user03", "user04", "user05"}
// 크기가 없는 배열선언
arr := []string{"user01", "user02", "user03"}

// Error
arr := []string{"user01", "user02", "user03"}
arr[3] = "user04"
fmt.Println(arr)
// 최초 선언된 배열의 크기가 3인데 4번째에 아이템을 넣으려고 하면 Error
// panic: runtime error: index out of range [3] with length 3

// append
// 기존에 알던 문법으로는.. push 등을 이용하는데.. 
// Go에서는 append를 이용하는데.. 컨셉이.. 배열과, 추가할 인자를 주면 새로운 배열이 리턴되고
// 리턴값을 다시 선언된 어레이에 넣어줘야 한다.
arr := []string{"user01", "user02", "user03"}
arr = append(arr,"user04")
```

# Struct

```go
// 일반적으로 알고 있는 Object와 비슷하다.
// struct 선언
type userObj struct {
	name  string
	age   int
	hobby []string
}
hobby01 := []string{"golf", "baseball"}
// struct에 선언한 순서 그대로 데이터를 입력하여 만들수 있다.
userInfo := userObj{"Lee", 42, hobby01}

// 좀더 명확하게 키와 함께 써도 된다. 
userInfo := userObj{"name": "Lee", "age": 42, "hobby": hobby01}
```

# Map

```go
// 기본 선언
user := map[string]string{"name":"이승재","id":"test","age":"33"}
          key 데이터타입 value 데이터타입
// 키, 밸류 모두 스트링 타입으로 만들었기 때문에 age를 숫자로 넣을 수가 없다.           
```