package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"go_crud/helper"
	"go_crud/loaders"
	"go_crud/trace"
	//"go_crud/models"
	//"log"
	//"os"
)

func main() {

	config := loaders.LoadConfig()

	basic := loaders.LoadBasic(config)
	dirs := trace.Directory(config)

	fmt.Printf(" ----------------  %T ||||  %T  +++ %+v \n", config, basic, dirs)
	helper.ListArr(dirs, "Directories:")

	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	//level := 1
	var ptr *[]string
	arrFolders := []string{"assets", "categories", "items"}
	arrFiles := []string{"create", "update", "delete", "get"}

	keyArr := map[int32]int32{48: 0, 49: 1, 50: 2, 51: 3, 52: 4, 53: 5, 54: 6, 55: 7, 56: 8, 57: 9, 97: 10,
		98: 11, 99: 12, 100: 13, 101: 14, 102: 15}
	var a int32
	a = 107
	fmt.Printf(".......... %T %[1]q \n", a)

	fmt.Println("Press ESC to quit")
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		} else if key == keyboard.KeyEsc {
			break
		} else if key == keyboard.KeyCtrlC {
			break
		}

		switch {
		case char == 120:
			fmt.Println("X pressed, it means go back")
			ptr = &arrFolders
		case (char > 47 && char < 58) || (char > 96 && char < 103):

			fmt.Println("navigation 0-9 a-f has been pressed", ptr)
			selectedNum := int(keyArr[char])
			var folders, files string
			if selectedNum < len(arrFolders) {
				folders = arrFolders[selectedNum]
			}
			if selectedNum < len(arrFiles) {
				files = arrFiles[selectedNum]
			}
			fmt.Println("....... ", folders, files)
		}

		fmt.Printf("You pressed:%[1]T %[1]q   %[1]v\r\n", char)
		fmt.Println(keyArr[char])
	}

	/*
		rr := basic.Payload.(map[string]interface{})
		fmt.Println("first basicFile", rr["category_id"])
		kk := helper.KeysString(rr)
		log.Printf(".........########$$$$$$$$      %T ------  %T \n::::::: %+v", rr, kk, kk)
	*/

	//helper.SimplePost()
}
