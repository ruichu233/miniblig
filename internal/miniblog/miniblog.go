// Copyright 2023 Ruichu Zhang <ruichu233@outlook.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
// The original repo for this file is https://github.com/ruichu233/miniblog.

package miniblog

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewMiniBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		// 指定命令的名字，该名字会出现在帮助信息中
		Use: "miniblog",
		// 命令的简短描述
		Short: "A good Go practical project",
		// 命令的详细描述
		Long: `A good Go practical project, used to create user with basic information.

Find more miniblog information at:
		https://github.com/ruichu233/miniblog#readme`,
		// 命令出错时，不打印帮助信息。不需要打印帮助信息，设置为 true 可以保持命令出错时一眼就能看到错误信息
		SilenceUsage: true,
		// 指定调用 cmd.Execute() 时，执行的 Run 函数，函数执行失败会返回错误信息
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
		// 这里设置命令运行时，不需要指定命令行参数
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(args) > 0 {
					return fmt.Errorf("%q does not take any argument, got %q", cmd.CommandPath(), arg)
				}
			}
			return nil
		},
	}
	return cmd
}

func run() error {
	fmt.Println("Hello MiniBlog!")
	return nil
}
