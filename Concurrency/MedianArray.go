package main

import(
	"fmt"
	"sort"
)

func FindMedianSortedArrays(arr1,arr2 []int) float64{
	arr1 = append(arr1,arr2...)
	sort.Ints(arr1)
	check:=len(arr1)%2
	len := len(arr1)-1
	if check==0{
		return float64((float64(arr1[len/2])+float64(arr1[(len/2)+1]))/2)
	}
	return float64(arr1[len/2])
}

func main(){
	nums1 := []int{1,3}
	nums2 := []int{2}
	fmt.Println(FindMedianSortedArrays(nums1,nums2))
}
