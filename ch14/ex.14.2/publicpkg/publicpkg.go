package publicpkg

import "fmt"

const (
	PI = 3.1415   // 공개되는 변수
	pi = 3.141516 // 공개되지 않는 변수
)

func PublicFunc() { // 공개 되는 함수
	const MyConst = 100
	fmt.Println("This is a public function", MyConst)
}

func privateFunc() { // 공개 되지 않는 함수
	fmt.Println("This is a private function")
}

type MyInt int       // 공개되는 별칭 타입
type myString string // 공개되지 않는 별칭 타입

type MyStruct struct {
	Age  int    // 공개되는 구조체 필드
	name string // 공개되지 않는 구조체 필드
}

func (m MyStruct) PublicMethod() { // 공개되는 메서드
	fmt.Println("This is a public method")
}

func (m MyStruct) PrivateMethod() { // 공개되지 않는 메서드
	fmt.Println("This is a private method")
}

type myPrivateStruct struct {
	Age  int    // 공개되지 않는 구조체 필드
	name string // 공개되지 않는 구조체 필드
}

func (m myPrivateStruct) PrivateMethod() { // 공개되지 않는 메서드
	fmt.Println("This is a private method")
}

func NewMyStruct(age int, name string) *MyStruct {
	return &MyStruct{
		Age:  age,
		name: name,
	}
}
