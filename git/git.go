package git

import "os/exec"

// WARNING: ONLY USE THESE WITH TRUSTED DATA

// BE CAREFUL WHERE YOU USE THIS, INPUTS ARE NOT SANITIZED
func Clone(url, path string) error {
	cmd := exec.Command("git", "clone", url, path)
	return cmd.Run()
}

// BE CAREFUL WHERE YOU USE THIS, INPUTS ARE NOT SANITIZED
func Status(path string) error {
	err := exec.Command("cd", path).Run()
	if err != nil {
		return err
	}

	return exec.Command("git", "status").Run()
}

// BE CAREFUL WHERE YOU USE THIS, INPUTS ARE NOT SANITIZED
func Pull(path string) error {
	err := exec.Command("cd", path).Run()
	if err != nil {
		return err
	}

	return exec.Command("git", "pull").Run()
}
