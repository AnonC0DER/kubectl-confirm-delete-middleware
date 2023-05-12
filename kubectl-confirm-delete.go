package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 && args[0] == "delete" {
		fmt.Print("You're running delete command, are you sure you're not gonna fuck up anything? (y/n) : ")

		var input string
		fmt.Scanln(&input)

		if strings.ToLower(input) == "y" {
			cmd := exec.Command("kubectl", args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else if strings.ToLower(input) == "n" {
			fmt.Println("Well, this time I saved your ass.")
			fmt.Println("Be more careful.")
		}
	} else {
		cmd := exec.Command("kubectl", args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
