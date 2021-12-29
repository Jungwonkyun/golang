package something

import "fmt"

//얘는 다른 패키지에서 불러올 수 없다 -> 소문자 시작 private function
func sayBye() {
	fmt.Println("Bye")
}

//얘는 불러올 수 있다. -> 대문자로 시작
func SayHello() {
	fmt.Println("Hello")
}
