package main

import "fmt"

func main() {
	s := "aaabcccdaa"
	fmt.Println(CompressString(s))
}

func CompressString(s string) string {

	count := 1
	var outString string
	for i := 1; i < len(s); i++ {

		if s[i-1] == s[i] {
			count++
		} else {
			outString = outString + string(s[i-1])
			if count > 1 {
				outString = outString + fmt.Sprint(count)
			}
			count = 1
		}

		//handle last character
		if i == len(s)-1 {
			if count > 1 {
				outString = outString + string(s[i]) + fmt.Sprint(count)
				return outString
			} else {
				outString = outString + string(s[i])
			}
		}
	}
	return outString

}
