package main

type Attacker interface {
	Attack()
}

func main() {
	var att Attacker // 기본값은 nil입니다.
	att.Attack()     // att가 nil이기 때문에 런 타임 에러가 발생합니다.
}
