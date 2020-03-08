package main

import "fmt"
import "bufio"
import "os"
import "strings"

//99算术表
func ArithmeticTable() {
	for y := 1; y <= 9; y++ {
		for x := 1; x <= 9; x++ {
			fmt.Printf("%d*%d=%d ", x, y, x*y)
		}
		fmt.Println()
	}
}

//聊天机器人
func TalkRoboot() {
	// get data from stdInput
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input you name")
	// get data until get \n
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("An error occured: %s \n", err)
		os.Exit(1)
	} else {
		name := input[:len(input)-2]
		fmt.Printf("Hello, %s! What can I do for you?\n", name)
	}

	for {
		input, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occured: %s \n", err)
			os.Exit(1)
		}
		input = input[:len(input)-2]
		input = strings.ToLower(input)
		switch input {
		case "":
			continue
		case "nothing", "bye":
			fmt.Println("Bye!")
			os.Exit(0)
		default:
			fmt.Println("Sorry, I didn't catch you")
		}
	}
}
