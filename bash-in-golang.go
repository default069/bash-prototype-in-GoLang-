package main

import (
	"fmt"
	"os"
)

func main(){
	var command string

	fmt.Println("Welcome to Bash in GoLang")
	for{

	
	fmt.Println("Please choos operator ls, cd, mkdir, touch ")
	fmt.Scan(&command)

	if command == "ls"{
		ls(".")
	}else if command == "cd"{
		var path string
		fmt.Println("Enter the path to the folder:")
		fmt.Scan(&path)
		cd(path)
	}
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
func cd(path string){
	err := os.Chdir(path)
	if err != nil {
		fmt.Println("Directory change error:")
		return
	}

	currentDir, err := os.Getwd()
	if err == nil {
		fmt.Println("We have successfully transitioned to:", currentDir)
	} 
}