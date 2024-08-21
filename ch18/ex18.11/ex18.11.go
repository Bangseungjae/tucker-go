package main

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct{}

func (f *File) Read() {}

//func (f *File) Close() {}

func ReadFile(reader Reader) {
	// Reader 인터페이스 변수를 Closer 인터페이스로 타입 변환합니다.
	//c, ok := reader.(Closer) // c는 변환된 값 ok는 타입 변환 성공 여부
	//if ok {
	//	c.Close()
	//}
	if c, ok := reader.(Closer); ok { // 한줄로 표현
		c.Close()
	}
}

func main() {
	// File 포인터 인스턴스를 ReadFile() 함수의 인수로 사용합니다.
	file := &File{}
	ReadFile(file)
}
