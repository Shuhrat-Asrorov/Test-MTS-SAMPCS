package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Profile struct {
	User    string `yaml:"user"`
	Project string `yaml:"project"`
}

func getProfilePath(name string) string {
	return filepath.Join(".", name+".yaml")
}

func loadProfile(name string) (*Profile, error) {
	data, err := os.ReadFile(getProfilePath(name))
	if err != nil {
		return nil, err
	}
	var p Profile
	err = yaml.Unmarshal(data, &p)
	return &p, err
}

func saveProfile(name string, p *Profile) error {
	data, err := yaml.Marshal(p)
	if err != nil {
		return err
	}
	return os.WriteFile(getProfilePath(name), data, 0644)
}

func deleteProfile(name string) error {
	return os.Remove(getProfilePath(name))
}

func listProfiles() ([]string, error) {
	files, err := filepath.Glob("*.yaml")
	if err != nil {
		return nil, err
	}
	names := make([]string, 0, len(files))
	for _, f := range files {
		name := f[:len(f)-5]
		names = append(names, name)
	}
	return names, nil
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	command := os.Args[1]

	switch command {
	case "help":
		printHelp()
	case "profile":
		if len(os.Args) < 3 {
			fmt.Println("Ожидается подкоманда: create, get, list, delete")
			printHelp()
			return
		}
		subcommand := os.Args[2]
		switch subcommand {
		case "create":
			createCmd := flag.NewFlagSet("create", flag.ExitOnError)
			name := createCmd.String("name", "", "Имя профиля")
			user := createCmd.String("user", "", "Пользователь")
			project := createCmd.String("project", "", "Проект")
			createCmd.Parse(os.Args[3:])
			if *name == "" || *user == "" || *project == "" {
				fmt.Println("Необходимы флаги --name, --user, --project")
				return
			}
			err := createProfile(*name, *user, *project)
			if err != nil {
				fmt.Println("Ошибка:", err)
				return
			}
			fmt.Printf("Профиль '%s' создан\n", *name)

		case "get":
			getCmd := flag.NewFlagSet("get", flag.ExitOnError)
			name := getCmd.String("name", "", "Имя профиля")
			getCmd.Parse(os.Args[3:])
			if *name == "" {
				fmt.Println("Необходим флаг --name")
				return
			}
			p, err := loadProfile(*name)
			if err != nil {
				fmt.Println("Ошибка загрузки:", err)
				return
			}
			fmt.Printf("User: %s\nProject: %s\n", p.User, p.Project)

		case "list":
			names, err := listProfiles()
			if err != nil {
				fmt.Println("Ошибка:", err)
				return
			}
			if len(names) == 0 {
				fmt.Println("Нет профилей")
				return
			}
			for _, n := range names {
				fmt.Println(n)
			}

		case "delete":
			deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
			name := deleteCmd.String("name", "", "Имя профиля")
			deleteCmd.Parse(os.Args[3:])
			if *name == "" {
				fmt.Println("Необходим флаг --name")
				return
			}
			err := deleteProfile(*name)
			if err != nil {
				fmt.Println("Ошибка удаления:", err)
				return
			}
			fmt.Printf("Профиль '%s' удалён\n", *name)

		default:
			fmt.Println("Неизвестная подкоманда:", subcommand)
			printHelp()
		}
	default:
		fmt.Println("Неизвестная команда:", command)
		printHelp()
	}
}

func createProfile(name, user, project string) error {
	if _, err := os.Stat(getProfilePath(name)); err == nil {
		return errors.New("профиль с таким именем уже существует")
	}
	p := &Profile{User: user, Project: project}
	return saveProfile(name, p)
}

func printHelp() {
	fmt.Println(`Использование:
  mk profile create --name <имя> --user <пользователь> --project <проект>
  mk profile get --name <имя>
  mk profile list
  mk profile delete --name <имя>
  mk help
`)
}