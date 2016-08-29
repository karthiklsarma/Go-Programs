package main

import (
	"sort"
)

type sortRunes []rune

func (s sortRunes)Less(i,j int) bool{
	return s[i]<s[j]
}

func (s sortRunes)Swap(i,j int){
	s[i],s[j]=s[j],s[i]
}

func (s sortRunes)Len()int{
	return len(s)
}

func SortString(s string)string{
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
func merge(ins []rune, c rune) (result []string) {
	for i := 0; i <= len(ins); i++ {
		result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
	}
	return
}

func permutations(input string) []string {
	if len(input) == 1 {
		return []string{input}
	}

	runes := []rune(input)
	subPermutations := permutations(string(runes[0:len(input) - 1]))

	result := []string{}
	for _, s := range subPermutations {
		result = append(result, merge([]rune(s), runes[len(input)-1])...)
	}

	return result
}


func permutationsTailRecursive(input string) []string {
	var helper func(input []rune, previousResult []string) []string
	helper = func(input []rune, previousResult []string) []string{
		if len(input) == 0 {
			return previousResult
		}else {
			newResult := []string{}
			for _, element := range previousResult {
				newResult = append(newResult, merge([]rune(element), input[0])...)
			}
			return helper(input[1:], newResult)
		}
	}

	runes := []rune(input)
	return helper(runes[1:], []string{string(runes[0])})
}

/*
func main(){
	str := "CAT"
	fmt.Println(permutationsTailRecursive(str))
}*/
