// Copyright 2023 Ruichu Zhang <ruichu233@outlook.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
// The original repo for this file is https://github.com/ruichu233/miniblog.

package miniblog

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// recommendedHomeDir 指定了 miniblog 服务的默认配置目录
	recommendedHomeDir = ".miniblog"
	// defaultConfigName 指定了 miniblog 服务的默认配置文件名
	defaultConfigName = "miniblog"
)

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// 获取用户主目录
		home, err := os.UserHomeDir()
		// 如果获取用户主目录失败，打印 `Error:xxx` 错误并退出程序（code=1）
		cobra.CheckErr(err)
		// 添加 `$HOME/<recommendedHomeDir>` 目录到配置文件搜索路径
		viper.AddConfigPath(filepath.Join(home, recommendedHomeDir))
		// 添加当前目录到配置文件搜索路径
		viper.AddConfigPath(filepath.Join("."))
		// 设置配置文件拓展名
		viper.SetConfigType("yaml")
		// 设置配置文件名
		viper.SetConfigName(defaultConfigName)

	}
	// 读取匹配的环境变量
	viper.AutomaticEnv()
	// 读取前缀为 `MINIBLOG_`的环境变量
	viper.SetEnvPrefix("miniblog")
	// 将 viper.Get(key) key 字符串中 '.' 和 '-' 替换为 '_'
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	// 读取配置文件。如果指定配置文件名，则使用指定的配置文件，否则在路径中搜索
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	// 打印 viper 当前使用的配置文件
	fmt.Fprintln(os.Stdout, "Using config file:", viper.ConfigFileUsed())
}
