package main

import(
	"fmt"
	"math"
)

func LCP_Merge(str []string) string{
	return Divide_Conquer(str,0,len(str)-1)
}

func Divide_Conquer(str []string, left int, right int) string{
	if(left==right){
		return str[left]
	}
	var mid int
	mid = left+right/2
	leftPrefix := []rune(Divide_Conquer(str,left,mid))
	rightPrefix := []rune(Divide_Conquer(str,mid+1,right))
	commonPrefix := make([]rune,0)
	min := int(math.Min(float64(len(leftPrefix)),float64(len(rightPrefix))))
	for i:=0;i<min;i++{
		if leftPrefix[i]!=rightPrefix[i]{
			break
		}else{
			commonPrefix=append(commonPrefix,leftPrefix[i])
		}
	}
	return string(commonPrefix)
}

func main(){
	strs := []string{
		"Mojo",
		"Mojo",
		"Mojo",
	}
	fmt.Println(LCP_Merge(strs))
}