package project

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	TS = "ts"
	JS = "js"
	GO = "go"
)

type Project struct {
	Name      string              `json:name`
	Type      string              `json:type`
	Framework string              `json:framework`
	Path      string              `json:path`
	Deps      map[string][]string `json:deps`
}

func (p *Project) Create() {
	log.Println("=> [+] Starting project creation")
	path := fmt.Sprintf("%s/%s", p.Path, p.Name)
	dir := os.Mkdir(path, 0755)

	if dir != nil {
		panic(dir)
	}

	log.Printf("=> [+] Project created on path %s", path)
	err := os.Chdir(path)
	if err != nil {
		panic(err)
	}

	log.Printf("=> [+] Open dir on path %s", path)

	switch p.Type {
	case TS:
		p.Deps["dev"] = append(p.Deps["dev"], "typescript", "@types/node", "nodemon")
		log.Printf("=> [+] Starting TS project on path %s", path)

		RunCommand("npm", "init", "-y")

		for _, dep := range p.Deps["dev"] {
			runInstall(dep)
		}

		RunCommand("npx", "tsconfig.json")

	case JS:
		p.Deps["dev"] = append(p.Deps["dev"], "nodemon")
		log.Printf("=> [+] Starting JS project on path %s", path)
		RunCommand("npm", "init", "-y")

		for _, dep := range p.Deps["dev"] {
			runInstall(dep)
		}

	case GO:
		log.Printf("=> [+] Starting GO project on path %s", path)
		RunCommand("go", "mod", "init", p.Name)
	}
}

func runInstall(dep string) {
	log.Printf("==> [+] Installing dependecy %s", dep)
	cmd := exec.Command("npm", "i", "-D", dep)

	stderr, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Printf("===> [+] Dependecy install log: %s\n", m)
	}
}

func RunCommand(command ...string) {
	log.Printf("==> Running command %s", strings.Join(command, " "))
	cmd := exec.Command(strings.Join(command, " "))

	stderr, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Printf("===> [+] Command %s log: %s\n", strings.Join(command, " "), m)
	}
}
