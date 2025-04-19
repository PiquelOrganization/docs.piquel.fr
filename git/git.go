package git

import (
	"log"
	"os/exec"
)

// WARNING: ONLY USE THESE WITH TRUSTED DATA

// BE CAREFUL WHERE YOU USE THIS, INPUTS ARE NOT SANITIZED
func Clone(url, path string) error {
	log.Printf("[Git] Cloning %s into %s", url, path)
	return exec.Command("git", "clone", url, path).Run()
}

// BE CAREFUL WHERE YOU USE THIS, INPUTS ARE NOT SANITIZED
func Status(path string) error {
	err := exec.Command("cd", path).Run()
	if err != nil {
		return err
	}

	log.Printf("[Git] Getting git status in %s", path)
	return exec.Command("git", "status").Run()
}

// BE CAREFUL WHERE YOU USE THIS, INPUTS ARE NOT SANITIZED
func Pull(path string) error {
	err := exec.Command("cd", path).Run()
	if err != nil {
		return err
	}

	log.Printf("[Git] Pulling repository in %s", path)
	return exec.Command("git", "pull").Run()
}
