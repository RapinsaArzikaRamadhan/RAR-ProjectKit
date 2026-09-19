package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/rapinsa/rar-project-kit/internal/generator"
)

func create_folder(path string, permission os.FileMode) error {
	err :=os.Mkdir(path, permission)
	if err != nil {
		return err
	}
	return nil
}

func create_file(nama string, template string) error {
	file, err := os.Create(nama)
	if err != nil {
		return err
	}
	defer file.Close()
	isi, err := generator.Template.ReadFile(template)
	if err != nil {
		return err
	}
	_,err = file.WriteString(string(isi))
	if err != nil {
		return err
	}
	return nil
}

func command_exec(nama string, command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = nama
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func go_sqlite(name string) {
	fmt.Println("✓ generating the folder")
	{
		err := create_folder(name, 0775)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = create_folder(name + "/web", 0775)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = create_folder(name + "/database", 0775)
		if err != nil {
			fmt.Println(err)
			return
		}
	}


	fmt.Println("✓ generating the file")
	{
		err := create_file(name + "/main.go", "templates/go-sqlite/main.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = create_file(name + "/web/index.html", "templates/go-sqlite/index.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = create_file(name + "/web/crud.html", "templates/go-sqlite/crud.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	
	fmt.Println("✓ running the command to..")
	{
		fmt.Println("✓ initializing go mod")
		err := command_exec(name, "go", "mod", "init", name)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("✓ installing sqlite dependency")
		
		err = command_exec(name, "go", "get", "-v", "modernc.org/sqlite")
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("✓ initializing sqlite databases")
		err = command_exec(name, "sqlite3", "database/database.db", ".databases")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	

	fmt.Println("done!")

}

func go_mysql(name string) {
	fmt.Println("✓ generating the folder")
	{
		err := create_folder(name, 0775)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = create_folder(name + "/web", 0775)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("✓ generating the file")
	{
		err := create_file(name + "/main.go", "templates/go-mysql/main.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = create_file(name + "/web/index.html", "templates/go-mysql/index.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = create_file(name + "/web/crud.html", "templates/go-mysql/crud.txt")
	}

	fmt.Println("✓ running the command to..")
	{
		fmt.Println("✓ initializing go mod")
		err := command_exec(name, "go","mod","init", name)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("✓ installing mysql dependency")
		err = command_exec(name, "go","get","-v","github.com/go-sql-driver/mysql")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	
}

func main() {
	
		if len(os.Args) < 2 {
			fmt.Println("please select the project type and its name")
			fmt.Println("--go-sqlite and --go-mysql for golang, etc")
			fmt.Println("example : rar-kit --go-mysql hello")
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
			case "--go-sqlite":
				go_sqlite(os.Args[2])
			case "--go-mysql":
				go_mysql(os.Args[2])
			case "--java":
				fmt.Println("N/A")
			default:
				fmt.Println("sorry, that language template not avaiable for now")
				fmt.Println("whats available for now is --go --cpp --java")
			}

			return
		}

}