package main

import(
	"fmt"
	"math"
)

func Check(s string) bool{
	arr := make([]bool,128)
	for _,ch := range s{
		if arr[ch]{
			return false
		}
		arr[ch]=true;
	}
	return true
}

func LenLargest(s string)float64{
	var max float64
	for i:=0;i<len(s);i++{
		for j:=1;j<=len(s)-i;j++{
			if Check(s[i:i+j]){
				max = math.Max(max,float64(len(s[i:i+j])))
			}
		}
	}
	return max
}

func main(){
	str := "abcde"
	fmt.Println(LenLargest(str))
}