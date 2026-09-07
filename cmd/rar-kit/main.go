package main

import (
	"fmt"
	"os"
	"os/exec"
	"github.com/rapinsa/rar-project-kit/internal/generator"
)

func generatego(name string) {
	fmt.Println("sedang membuat", name)
	err := os.Mkdir(name, 0775)

	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := generator.Template.ReadFile("templates/go/main.go")

	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Create(name + "/main.go")

	if err != nil {
		fmt.Println(err)
		return
	}

	_,err = file.WriteString(string(data))

	if err != nil {
		fmt.Println(err)
		return
	}

	file.Close()

	initial := exec.Command("go", "mod", "init", name)

	initial.Dir = name

	err = initial.Run()

	if err != nil {
		fmt.Println(err)
		return
	}

}

func main() {
	
		if len(os.Args) < 2 {
			fmt.Println("please select the project type and its name")
			fmt.Println("--go for golang, --cpp for cpp, etc")
			fmt.Println("example : rar-kit --go hello")
			return
		} 

		if len(os.Args) < 3 {
			fmt.Println("please select the project name")
			fmt.Println("word spacing is forbiden, use - / _ instead for space beetwen word")
			fmt.Println("valid : hello_world_project")
			fmt.Println("invalid : hello world project")
			return
		} 

		if len(os.Args) > 3 {
			fmt.Println("please use single project name")
			fmt.Println("valid : hello_world_project")
			fmt.Println("invalid : hello world project")
			return
		} 

		if len(os.Args) == 3 {
			fmt.Println("project flag:", os.Args[1])
			fmt.Println("project name:", os.Args[2])

			switch os.Args[1] {
			case "--go":
				generatego(os.Args[2])
			case "--cpp":
				fmt.Println("aku tau ini bisa di jadikan fungsi lain atau calling calling func")
			case "--java":
				fmt.Println("aku tau ini bisa di jadikan fungsi lain atau calling calling func")
			default:
				fmt.Println("sorry, that language template not avaiable for now")
				fmt.Println("whats available for now is --go --cpp --java")
			}

			return
		}

}