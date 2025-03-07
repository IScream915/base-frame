package cmd

import (
	"base_frame/internal"
	"base_frame/pkg/config"
	"context"
	"github.com/spf13/cobra"
)

type RootCmd struct {
	command cobra.Command
	ctx     context.Context // 设置基本的上下文
	config  config.Config   // 配置项
}

func NewApiCmd() *RootCmd {
	var ret RootCmd
	ret.ctx = context.WithValue(context.Background(), "version", "test_version")
	ret.command.RunE = func(cmd *cobra.Command, args []string) error {
		return ret.runE()
	}
	return &ret
}

func (a *RootCmd) runE() error {
	return internal.Start(a.ctx, &a.config)
}

func (a *RootCmd) Exec() error {
	return a.command.Execute()
}
