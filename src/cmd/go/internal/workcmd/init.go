// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// go work init

package workcmd

import (
	"cmd/go/internal/base"
	"cmd/go/internal/modload"
	"context"
	"path/filepath"
)

var cmdInit = &base.Command{
	UsageLine: "go work init [moddirs]",
	Short:     "initialize workspace file",
	Long: `Init initializes and writes a new go.work file in the
current directory, in effect creating a new workspace at the current
directory.

go work init optionally accepts paths to the workspace modules as
arguments. If the argument is omitted, an empty workspace with no
modules will be created.

Each argument path is added to a use directive in the go.work file. The
current go version will also be listed in the go.work file.

`,
	Run: runInit,
}

func init() {
	base.AddModCommonFlags(&cmdInit.Flag)
	base.AddWorkfileFlag(&cmdInit.Flag)
}

func runInit(ctx context.Context, cmd *base.Command, args []string) {
	modload.InitWorkfile()

	modload.ForceUseModules = true

	// TODO(matloob): support using the -workfile path
	// To do that properly, we'll have to make the module directories
	// make dirs relative to workFile path before adding the paths to
	// the directory entries

	workFile := modload.WorkFilePath()
	if workFile == "" {
		workFile = filepath.Join(base.Cwd(), "go.work")
	}

	modload.CreateWorkFile(ctx, workFile, args)
}
