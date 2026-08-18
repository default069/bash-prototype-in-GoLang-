package main

import (
	"fmt"
	"os"
)

func main(){
	var command string

	fmt.Println("Welcome to Bash in GoLang")
	fmt.Println("Please choos operator ls, cd, mkdir, touch ")
	fmt.Scan(&command)

	if command == "ls"{
		ls(".")
	}
}

func ls(path string){
	entries , err := os.ReadDir(path)
	if err != nil{
		fmt.Println("directory read error:", err)
		return
	}

	for _, entry := range entries{
		if entry.IsDir(){
			fmt.Printf("%s\n", entry.Name())
		}else{
			fmt.Printf("%s\n",entry.Name())
		}
	}
}