package fileenv

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/fileenv", new(Fileenv))

	load_env_file()
}

type Fileenv struct{}

const (
	env_var = "K6_FILE_ENV"
)

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func get_file_content(path string) []string {
	log.Printf(":))))))))))")
	pwd, _ := os.Getwd()
	log.Printf(pwd)
	log.Printf(path)

	bytesRead, _ := ioutil.ReadFile(pwd + path)
	file_content := string(bytesRead)
	log.Printf(file_content)

	lines := strings.Split(file_content, "\n")
	return lines
}

func set_variables(path string) {
	lines := get_file_content(path)
	log.Printf(lines[0])
	for i := 1; i <= len(lines); i++ {
		line := lines[i]
		log.Printf(line)

		parts := strings.SplitN(line, "=", 2)
		log.Printf(parts[0])
		log.Printf(parts[1])

		err := os.Setenv(parts[0], parts[1])

		if err != nil {
			log.Fatalf("An error occured while setting environment variables %v. Check that each line is in the KEY=VALUE convention.", err)
		}
	}
}

func load_env_file() {
	env_file_path := os.Getenv(env_var)

	// if env file specified
	if env_file_path != "" {
		if_exists, err := exists(env_file_path)
		if err != nil {
			log.Fatalf("An error occurred while checking for a file with environment variables %v. Remember that the path MUST be relative.", err)
			return
		}

		if !if_exists {
			log.Fatalf("The specified path to the environment variables file does not exist. The name of the environment variables file: %v. Remember that the path MUST be relative.", env_file_path)
			return
		}

		if if_exists {
			log.Printf("The name of the environment variables file path: %v.", env_file_path)
			set_variables(env_file_path)
			return
		}
	}
}
