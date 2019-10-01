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
	level := 1
	var selectedDir string
	var files []string

	fmt.Printf(" ----------------  %T ||||  %T  +++ %+v \n", config, basic, dirs)
	//helper.ListArr(dirs, "Directories:")

	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	var ptr *[]string
	//arrFolders := []string{"assets", "categories", "items"}
	//arrFiles := []string{"create", "update", "delete", "get"}

	//97: 10, 98: 11, 99: 12, 100: 13, 101: 14, 102: 15
	keyArr := map[int32]int32{
		48: 0, 49: 1, 50: 2, 51: 3, 52: 4, 53: 5, 54: 6, 55: 7, 56: 8, 57: 9,
		65: 10, 66: 11, 67: 12, 68: 13, 69: 14, 70: 15, 71: 16, 72: 17, 73: 18,
		74: 19, 75: 20, 76: 21, 77: 22, 78: 23, 79: 24, 80: 25, 81: 26, 82: 27,
		83: 28, 84: 29, 85: 30, 86: 31, 87: 32, 88: 33, 89: 34, 90: 35,
	}
	var a int32
	a = 107
	fmt.Printf(".......... %v %[1]q \n", a)

	fmt.Println("Press ESC to quit")
	for {
		switch level {
		case 1:
			helper.ListArr(dirs, "Directories: ")
		case 2:
			helper.ListArr(files, "Files: ")

		}

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
			level--
			if level < 1 {
				level = 1
			}
		case char == 100:
			level = 1
		case char == 102:
			if selectedDir == "" {
				fmt.Println("NOTICE: at first choose direcotry")
			}
			level = 2

		case (char > 47 && char < 58) || (char >= 65 && char <= 90):

			fmt.Println("navigation 0-9 a-f has been pressed", ptr)
			selectedNum := int(keyArr[char])
			switch level {
			case 1:
				if selectedNum < len(dirs) {
					selectedDir = dirs[selectedNum]
					files = trace.Files(config, selectedDir)
					level = 2
				}
			case 2:
				if selectedNum < len(files) {
					selectedFile := files[selectedNum]
					fileContent := loaders.Open(config, selectedDir, selectedFile)
					fmt.Printf("%v", fileContent)

					level = 3
				}
			}
			/*
				if selectedNum < len(arrFiles) {

					files = arrFiles[selectedNum]
				}
				fmt.Println("....... ", selectedDir, files)
			*/
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
