package main

import "fmt"

func selection_sort(arr[] int)[]int{
	var temp int
	var min int
	for i := 0 ; i<len(arr)-1; i++{
		min = i
		for j:=i+1 ; j<len(arr); j++{
			if arr[j] < arr[min]{
				min = j
			}
		}
		temp = arr[i]
		arr[i] = arr[min]
		arr[min]= temp

	}
	return arr

}

func main(){
	arr := []int{3,4,9,1,5,9}
	fmt.Println(selection_sort(arr))
}
