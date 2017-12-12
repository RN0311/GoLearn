package main
import "fmt"

func plus(a int, b int) int{
	return a+b
}

func doubleplus(a, b, c int) int{
	return a+b+c
}

func main(){
	res := plus(2,4)
	fmt.Println("2+4=",res)

	res = doubleplus(2,4,6)
	fmt.Println("2+4+6=",res)
}
