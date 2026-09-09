package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/rapinsa/rar-project-kit/internal/generator"
)

func generatego(name string) {
	fmt.Println("sedang membuat", name)
	
	//creating folder and checking the if the error present
	err := os.Mkdir(name, 0775)
	if err != nil {
		fmt.Println(err)
		return
	}
	//making file named main.go and checking if the error present
	file, err := os.Create(name + "/main.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	//reading the template and checking if the error present
	data, err := generator.Template.ReadFile("templates/go/main.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//insert the template main.go into the freshly made /main.go inside the folder
	_,err = file.WriteString(string(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	// closed after inserting the template
	file.Close()

	// execute the command "go mod init {name}"
	cmd1 := exec.Command("go", "mod", "init", name)
	// execute it in the folder {name}
	cmd1.Dir = name
	//error checking
	err = cmd1.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
	// go initialization command end here


	// same as above, but this one is downloading the sqlite dependency
	cmd2 := exec.Command("go", "get", "modernc.org/sqlite")
	cmd2.Dir = name
	err = cmd2.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
	// downloading sqlite dependency end here

	//this one make the index.html
	file1, err := os.Create(name + "/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	data1, err := generator.Template.ReadFile("templates/go/index.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = file1.WriteString(string(data1))
	if err != nil {
		fmt.Println(err)
		return
	}
	file1.Close()
	//index creating ends here

	//this one create crud.html
	file2, err := os.Create(name + "/crud.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	data2, err := generator.Template.ReadFile("templates/go/crud.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = file2.WriteString(string(data2))
	if err != nil {
		fmt.Println(err)
		return
	}
	file2.Close()
	//crud creating ends here

	cmd3 := exec.Command("sqlite3", "database.db", ".databases")
	cmd3.Dir = name
	err = cmd3.Run()
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