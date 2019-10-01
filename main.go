package main

import (
	"encoding/json"
	"fmt"
	"github.com/eiannone/keyboard"
	"go_crud/helper"
	"log"
	"os"
)

func main() {
	basicFile, err := os.Open("basic.json")
	if err != nil {
		log.Fatal("can't open tokens.json: ", err)
	}
	defer basicFile.Close()
	decoder := json.NewDecoder(basicFile)

	basic := helper.Basics{}
	err = decoder.Decode(&basic)
	if err != nil {
		log.Fatal("can't decode basic JSON: ", err)
	}

	err = keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	level := 1
	arrFolders := []string{"assets", "categories", "items"}

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
