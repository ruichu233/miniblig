// Copyright 2023 Ruichu Zhang <ruichu233@outlook.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
// The original repo for this file is https://github.com/ruichu233/miniblog.

package main

import (
	"os"

	"github.com/ruichu233/miniblog/internal/miniblog"
)

func main() {
	cmd := miniblog.NewMiniBlogCommand()

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}

}
