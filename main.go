package main

import "fmt"

type triangle struct{
	size int
}

func (t triangle) perimeter() int{
	return t.size * 3
}
// Method arguement is too large or update a variable
func ( t *triangle) doubleSize(){
	t.size*=2
}
func main(){
	t := triangle{
		size: 3,
	}
	fmt.Println("Perimeter:", t.perimeter())

}