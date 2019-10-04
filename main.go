package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"go_crud/helper"
	"go_crud/loaders"
	//"go_crud/models"
	"go_crud/engine"
	. "go_crud/global"
	"go_crud/trace"
	//"log"
	//"os"
)

func main() {

	engine := new(engine.Engine)
	engine.Config = loaders.LoadConfig()

	basic := loaders.LoadBasic(engine.Config)

	engine.Dirs = trace.Directory(engine.Config)

	engine.Env = helper.ParseEnv(basic.Env)
	engine.InitEnvArr(DIRECTORIES)

	fmt.Printf(" ----------------  %T ||||  %+v  +++ %+v \n\n ======ENV %+v \n", engine.Config, basic, engine.Dirs, engine.Env)

	//fC := loaders.Open(engine.Config, "assets", "k9.json")

	//for k, v := range basic.Payload {
	//switch vv := v.(type) {
	/*
		k := 22
		switch vv := basic.Payload.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)

				u2 := u.(map[string]interface{})
				fmt.Println("................   @@@@@@  ", u2)
				for i3, u3 := range u2 {
					fmt.Println("~~~~~~~~~~~~ ", i3, u3)
				}
			}

		case interface{}:
			fmt.Println(k, "is an interface:")
			ff := vv.(map[string]interface{})
			fmt.Println(ff, "is an interface: parsed")
			for zk, zv := range ff {
				switch zz := zv.(type) {
				case string:
					fmt.Println("_______ inside zz", zz, zk, zv)
				}
			}

		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
		//}
	*/

	//engine.Env := basic.Env.([]map[string]interface{})
	//fmt.Printf("&&&&&&&&&&&&&&&&&&   %+v -----  %T \n", rr, rr)

	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	var ptr *[]string

	keyArr := map[int32]int32{
		48: 0, 49: 1, 50: 2, 51: 3, 52: 4, 53: 5, 54: 6, 55: 7, 56: 8, 57: 9,
		65: 10, 66: 11, 67: 12, 68: 13, 69: 14, 70: 15, 71: 16, 72: 17, 73: 18,
		74: 19, 75: 20, 76: 21, 77: 22, 78: 23, 79: 24, 80: 25, 81: 26, 82: 27,
		83: 28, 84: 29, 85: 30, 86: 31, 87: 32, 88: 33, 89: 34, 90: 35,
	}

	fmt.Println("Press ESC to quit")
	engine.ShowPage(engine.Dirs, "Directories: ")
	for {
		/*
			switch engine.Page {
			case DIRECTORIES:
				engine.ShowPage(engine.Dirs, "Directories: ")
			case FILES:
				engine.ShowPage(engine.Files, "Files: ")
			case ENVIRONMENTS:
				engine.ShowPage(engine.EnvArr, "Environments: ")

			}
		*/

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
			engine.Page--
			if engine.Page < 1 {
				engine.Page = 1
			}
		case char == 100:
			engine.Page = DIRECTORIES
			engine.ShowPage(engine.Dirs, "Directories: ")
		case char == 102:
			if engine.SelectedDir == "" {
				fmt.Println("NOTICE: at first choose direcotry")
			}
			engine.Page = FILES
			engine.ShowPage(engine.Files, "Files: ")

		case (char > 47 && char < 58) || (char >= 65 && char <= 90):

			fmt.Println("navigation 0-9 a-f has been pressed", ptr)
			selectedNum := int(keyArr[char])
			switch engine.Page {
			case DIRECTORIES:
				if selectedNum < len(engine.Dirs) {
					engine.PrePage = engine.Page
					engine.SelectedDir = engine.Dirs[selectedNum]
					engine.Files = trace.Files(engine.Config, engine.SelectedDir)
					engine.Page = FILES
					engine.ShowPage(engine.Files, "Files: ")
				}
			case FILES:
				if selectedNum < len(engine.Files) {
					engine.PrePage = engine.Page
					selectedFile := engine.Files[selectedNum]
					fileContent := loaders.Open(engine.Config, engine.SelectedDir, selectedFile)
					fmt.Printf("%v", fileContent)

					engine.Page = CONTENT
				}
			case ENVIRONMENTS:
				if selectedNum < len(engine.EnvArr) {
					engine.Page = engine.PrePage
					engine.SelectedEnv = engine.EnvArr[selectedNum]

				}
			}
			/*
				if selectedNum < len(arrFiles) {

					engine.Files = arrFiles[selectedNum]
				}
				fmt.Println("....... ", engine.SelectedDir, engine.Files)
			*/
		case char == 101:
			engine.Page = ENVIRONMENTS
			engine.ShowPage(engine.EnvArr, "Environments: ")
		}

		fmt.Printf("You pressed:%[1]T %[1]q   %[1]v\r\n", char)
		fmt.Println(keyArr[char], engine.PrePage, engine.SelectedEnv)
	}

	/*
		rr := basic.Payload.(map[string]interface{})
		fmt.Println("first basicFile", rr["category_id"])
		kk := helper.KeysString(rr)
		log.Printf(".........########$$$$$$$$      %T ------  %T \n::::::: %+v", rr, kk, kk)
	*/

	//helper.SimplePost()
}
