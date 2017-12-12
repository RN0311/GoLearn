package main

import "fmt"
func main(){
	var arr [5]int
	fmt.Println("emp:",arr)

	arr[4] = 3
	fmt.Println("set:",arr)
	fmt.Println("get:",arr[4])
	fmt.Println("len of array:",len(arr))

	brr := [5]int{2,4,6,8,10}
	fmt.Println("dcl:",brr)

	var twoD [3][3]int
	for i :=0;i<3;i++{
		for j :=0;j<3;j++{
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D: ",twoD)

}