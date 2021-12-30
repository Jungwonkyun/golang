package main

import (
	"fmt"
	"strings"
)

type person struct {
	name    string
	age     int
	favFood []string
}

//naked return
func lenAndUpper(name string) (lenght int, uppercase string) {
	defer fmt.Println(("I'm done")) // excute after end of function
	lenght = len(name)
	uppercase = strings.ToUpper((name))
	return
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total

	//Same with upper loop
	/*
		for i := 0; i < len(numbers); i++ {
			fmt.Println((nu,bers[i]))
		}
	*/

}

/*func canIDrink(age int) bool {

	//variable expression only if-else문에서만 exclusive하게 사용하겠다.
	if koreanAge := age+2; koreanAge < 18 {
		return false
	}
	return true
}*/

func canIDrink2(age int) bool {

	//variable expression only if-else문에서만 exclusive하게 사용하겠다.
	switch koreanAge := age + 2; {

	case koreanAge < 18:
		return false

	case koreanAge >= 18 && koreanAge <= 50:
		return true

	case koreanAge > 50:
		return false
	}

	return false
}

/*func multyply(a, b int) int {
	return a * b
}*/

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	totalLength, uppername := lenAndUpper("wonkyun")
	fmt.Println(totalLength, uppername)
	repeatMe("A", "B", "C", "D", "E", "F")
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println((result))
	fmt.Println(canIDrink2(15))

	names := []string{"nico", "lynn", "dal"}
	names = append(names, "lala")
	fmt.Println((names))

	nico := map[string]string{"name": "nico", "age": "12"}
	fmt.Println((nico))

	for _, value := range nico {
		fmt.Println((value))
	}

	favFood := []string{"kimchi", "ramen"}

	//make structure
	nico2 := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println((nico2.age))
}
