package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("sh", "-c", "zip -r files.zip files && mv files.zip backups/")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	} else {
		println("Directory successfully zipped and moved.")
	}

}
