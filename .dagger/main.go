package main

import (
	"context"

	"dagger/emacs-ci/internal/dagger"
)

type EmacsCi struct{}

// Validate loads init.el into a containerized Emacs 29 and checks it starts without errors.
func (m *EmacsCi) Validate(ctx context.Context, source *dagger.Directory) (string, error) {
	return dag.Container().
		From("silex/emacs:29").
		WithFile("/root/.emacs.d/init.el", source.File("init.el")).
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
}
