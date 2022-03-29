package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var name, race, job string
var level int

func basicInfo() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("What is your character's name?")
	input.Scan()
	name = input.Text()

	fmt.Printf("What level is %v?\n", name)
	fmt.Scan(&level)

	input = bufio.NewScanner(os.Stdin)
	fmt.Printf("What species is %v?\n", name)
	input.Scan()
	race = input.Text()

	input = bufio.NewScanner(os.Stdin)
	fmt.Println("Which class will you be playing?")
	input.Scan()
	job = input.Text()

	time.Sleep(time.Second)

	fmt.Printf("%v is a level %v %v %v.\n\n", name, level, race, job)

	time.Sleep(time.Second)

}
