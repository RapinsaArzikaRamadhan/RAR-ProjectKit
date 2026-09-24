package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/rapinsa/rar-project-kit/internal/generator"

)

func create_folder(path string, permission os.FileMode) error {
	err :=os.Mkdir(path, permission)
	if err != nil {
		return err
	}
	return nil
}

func create_file(nama string, template string, module string) error {
	file, err := os.Create(nama)
	if err != nil {
		return err
	}
	defer file.Close()
	isi, err := generator.Template.ReadFile(template)
	if err != nil {
		return err
	}
	isi_string := strings.ReplaceAll(
		string(isi),
		"{{MODULE}}",
		module,
	)
	_,err = file.WriteString(isi_string)
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
		err := create_file(name + "/main.go", "templates/go-sqlite/main.txt", name)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = create_file(name + "/web/index.html", "templates/go-sqlite/index.txt", name)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = create_file(name + "/web/crud.html", "templates/go-sqlite/crud.txt", name)
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

func go_sqlite_restful(name string) {

	fmt.Println("✓ generating the folder")
	{
		err := create_folder(name, 0775)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = create_folder(name + "/database", 0775)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = create_folder(name + "/handler", 0775)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = create_folder(name + "/js", 0775)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = create_folder(name + "/middleware", 0775)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = create_folder(name + "/model", 0775)
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
		err := create_file(name + "/main.go", "templates/go-sqlite-restful/main.txt", name)
		if err != nil {
			fmt.Println(err)
			return
		}
		// /web
		err = create_file(name + "/web/index.html", "templates/go-sqlite-restful/web/index.txt", name)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = create_file(name + "/web/crud.html", "templates/go-sqlite-restful/web/crud.txt", name)
		if err != nil {
			fmt.Println(err)
			return
		}
		// /js
		err = create_file(name + "/js/create.js", "templates/go-sqlite-restful/js/create.txt", name)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = create_file(name + "/js/delete.js", "templates/go-sqlite-restful/js/delete.txt", name)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = create_file(name + "/js/read.js", "templates/go-sqlite-restful/js/read.txt", name)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = create_file(name + "/js/update.js", "templates/go-sqlite-restful/js/update.txt", name)
		if err != nil {
			fmt.Println(err)
			return
		}
		// /middleware
		err = create_file(name + "/middleware/middleware.go", "templates/go-sqlite-restful/middleware/middleware.txt", name)
		if err != nil {
			fmt.Println(err)
			return
		}
		// /model
		err = create_file(name + "/model/model.go", "templates/go-sqlite-restful/model/model.txt", name)
		if err != nil {
			fmt.Println(err)
			return
		}
		// /database
		err = create_file(name + "/database/database.go", "templates/go-sqlite-restful/database/database.txt", name)
		if err != nil {
			fmt.Println(err)
			return
		}
		// /handler
		err = create_file(name + "/handler/handler.go", "templates/go-sqlite-restful/handler/handler.txt", name)
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
		err = command_exec(name, "sqlite3", "database.db", ".databases")
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
		err := create_file(name + "/main.go", "templates/go-mysql/main.txt", name)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = create_file(name + "/web/index.html", "templates/go-mysql/index.txt", name)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = create_file(name + "/web/crud.html", "templates/go-mysql/crud.txt", name)
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
			case "--go-sqlite-restful":
				go_sqlite_restful(os.Args[2])
			default:
				fmt.Println("sorry, that language template not avaiable for now")
				fmt.Println("whats available for now is --go --cpp --java")
			}

			return
		}

}