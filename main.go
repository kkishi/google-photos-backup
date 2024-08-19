package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	dryRun = flag.Bool("dry_run", false, "")
	dst    = flag.String("dst", "", "")
	src    = flag.String("src", "", "")
)

func main() {
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	dir := filepath.Join(*dst, "DCIM/GooglePhotosBackup", *src)

	argv := []string{
		"--include=*.JPG", "--include=*/", "--exclude=*",
		"--inplace",
		"--no-perms",
		"--omit-dir-times",
		"--progress",
		"--recursive",
		"--verbose",
	}
	if *dryRun {
		argv = append(argv, "--dry-run")
	}
	cmd := exec.Command("rsync", append(argv, filepath.Clean(*src)+"/", dir+"/")...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if !*dryRun {
		fmt.Printf("Running the command (y/n): %s\n> ", cmd.String())
		var input string
		fmt.Scanf("%s", &input)
		if input != "y" {
			return
		}
	}

	if err := os.MkdirAll(dir, 0750); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
