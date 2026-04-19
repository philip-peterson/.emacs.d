package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Emacs config validated successfully.")
}

func run() error {
	ctx := context.Background()

	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return err
	}
	defer client.Close()

	initEl := client.Host().File("../init.el")

	// silex/emacs:29 has a recent Emacs with use-package built in
	_, err = client.Container().
		From("silex/emacs:29").
		WithFile("/root/.emacs.d/init.el", initEl).
		WithExec([]string{
			"emacs", "--batch",
			"--eval", `(setq use-package-always-ensure t)`,
			"--eval", `(require 'package)`,
			"--eval", `(add-to-list 'package-archives '("melpa" . "https://melpa.org/packages/") t)`,
			"--eval", `(package-initialize)`,
			"--eval", `(package-refresh-contents)`,
			"--load", "/root/.emacs.d/init.el",
		}).
		Stdout(ctx)

	return err
}
