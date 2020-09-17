package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"math/rand"
	"time"
)

var choice int
var sentence string
var input = bufio.NewScanner(os.Stdin)



func main() {
	
	fmt.Print("\n", `Welcome to the text transformator: 
	Press (1) to transform your sentence reverse, 
	Press (2) to transform your sentence in random sequence,
	Press (3) to transform your sentence in upper case,
	Press (4) to transform your sentence in lower case,
	Press (5) to make every first letter in upper case,
	Press (6) to go out of the program.
	Please choose one of the options: `) 
	fmt.Scanln(&choice)	
	
	switch choice {
	case 1, 2, 3, 4, 5:
		writeSentence()
	case 6:
		fmt.Println("You have been exited the program") 
		break 
	default: 
		fmt.Print("This is not a command. Try again")
		fmt.Scanln(&choice)	
		main()
	}

	}

func writeSentence() {
	fmt.Print("Write your sentence: ")
	input.Scan()
	sentence = input.Text()
	
	menu()
}


func continueOrExit() {
	fmt.Print(`
If you want to continue with the same sentence, press (c) and then choose the option, 
if you want to start again with a new sentence, press (n), 
if you want to exit the programm, press (x).
Choose the option: `)
	var nextStep string
	fmt.Scanln(&nextStep)
		switch nextStep {
		case "c":
			fmt.Print("Please choose one another option: ")
			fmt.Scanln(&choice)	
			menu()
		case "n": 
			main()
		case "x":
			fmt.Println("You have been exited the program")
			break
		default:
			fmt.Print("This is not a command. Try again")
			continueOrExit()
				
		}

}

func menu() {
	
	switch choice {
	case 1: 
		reverse()
	case 2:
		random()
	case 3:
		upperCase()
	case 4:
		lowerCase()
	case 5:
		firstLettersUpper()
	case 6:		 
		fmt.Println("You have been exited the program") 
		break 
	default:
		fmt.Print("This is not a command. Try again ")
		fmt.Scanln(&choice)
		menu()
	}
} 

func reverse() {
	sentenceReverse := strings.Split(sentence, "")
	fmt.Print("Here is your reversed sentence: ")
	for i := len(sentenceReverse) - 1; i >= 0; i-- {
		fmt.Print(sentenceReverse[i])
	} 

	continueOrExit()
 }

func random() {
	sentenceInSlice := strings.Fields(sentence)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(sentenceInSlice), func(i, j int) {
		sentenceInSlice[i], sentenceInSlice[j] = sentenceInSlice[j], sentenceInSlice[i]
	})
	fmt.Print("Here is your random ordered sentence: ", strings.Join(sentenceInSlice, " "))

	continueOrExit()
} 


func upperCase() {
	sentenceUpperCase := strings.ToUpper(sentence)
	fmt.Print("Here is your sentence in upper case: ", sentenceUpperCase)
	
	continueOrExit()
}

func lowerCase() {
	sentenceInLowerCase := strings.ToLower(sentence)
	fmt.Print("Here is your sentence in lower case: ", sentenceInLowerCase)
	
	continueOrExit()
}

 
func firstLettersUpper() {
	sentenceFirstLetterUpper := strings.Title(sentence)
	fmt.Print("Here is your sentence with the first capital letter in every word: ", sentenceFirstLetterUpper)
	
	continueOrExit()
}
