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
    cmd := exec.Command("git", "status")
    cmd.Dir = path

	log.Printf("[Git] Getting status of repository in %s", path)
	return cmd.Run()
}

// BE CAREFUL WHERE YOU USE THIS, INPUTS ARE NOT SANITIZED
func Pull(path string) error {
    cmd := exec.Command("git", "pull")
    cmd.Dir = path

	log.Printf("[Git] Pulling repository in %s", path)
	return cmd.Run()
}
