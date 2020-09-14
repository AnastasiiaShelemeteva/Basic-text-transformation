package main 

import (
	"fmt"
	"strings"
)
/* /var import []string 
fmt.Scanln(&import)

func reverse(input string) {
	for i := len(input) - 1; i >= 0; i-- {
		fmt.Println(input[i])
	} 
} 

func main() {
	reverse("hey")
}  */

/* func main () {
	var a[] string = [] string {"hallo", "fdsdf", "hgfolpj"}
	for i := len(a) - 1; i >= 0; i-- {
		fmt.Println(a[i])
	}
} // here I am trying to find out, how to get my input reverse. It works with numbers and words (each one of them a string, but how to take one string and make the letters #
	inside to be reversed */

/* func lowerCase(input string) {
		output := strings.ToLower(input)
		fmt.Println(output)	
		}

func main() {
	lowerCase("GZsdfJHSdbsSDSS")
} // here (28-35) I created a function to do a lowerer case. It works */

func upperCase(input string) {
	output := strings.ToUpper(input)
	fmt.Println(output)	
	}

func main() {
upperCase("GZsdfJHSdbsSDSS")
}