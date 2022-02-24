package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println("A simple calculator")

	//panic("This is my error message")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter first number:")
	numFirst, _ := reader.ReadString('\n')
	aFloat, err1 := strconv.ParseFloat(strings.TrimSpace(numFirst), 64)

	if err1 != nil {
		fmt.Println(err1)
		panic("You first number is invalud!!")
	}

	fmt.Print("Enter second number:")
	numSecond, _ := reader.ReadString('\n')

	bFloat, err2 := strconv.ParseFloat(strings.TrimSpace(numSecond), 64)

	if err2 != nil {
		fmt.Println(err2)
		panic("You second number is invalud!!")
	}

	sumNum := aFloat + bFloat
	sumNum = math.Round(sumNum*100) / 100

	fmt.Println("A simple calculator: ", sumNum)
}
