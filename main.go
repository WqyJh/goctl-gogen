package main

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"github.com/wqyjh/goctl-gogen/gogen"
)

const (
	codeFailure = 1
)

var (
	version = "20230306"
	Cmd     = &cobra.Command{
		Use:     "goctl-gogen",
		Short:   "Generate go api related files",
		RunE:    gogen.GoCommand,
		Version: version,
	}
)

func init() {
	// Cmd.Flags().StringVarP(&gogen.VarStringDir, "dir", "d", "", "The target dir")
	// Cmd.Flags().StringVarP(&gogen.VarStringAPI, "api", "a", "", "The api file")
	Cmd.Flags().StringVar(&gogen.VarStringHome, "home", "", "The goctl home path of "+
		"the template, --home and --remote cannot be set at the same time, if they are, --remote "+
		"has higher priority")
	Cmd.Flags().StringVar(&gogen.VarStringRemote, "remote", "", "The remote git repo "+
		"of the template, --home and --remote cannot be set at the same time, if they are, --remote"+
		" has higher priority\nThe git repo directory must be consistent with the "+
		"https://github.com/zeromicro/go-zero-template directory structure")
	Cmd.Flags().StringVar(&gogen.VarStringBranch, "branch", "", "The branch of "+
		"the remote repo, it does work with --remote")
	// Cmd.Flags().StringVarP(&gogen.VarStringStyle, "style", "s", "gozero", "The file naming format,"+
	// " see [https://github.com/zeromicro/go-zero/blob/master/tools/goctl/config/readme.md]")
}

func main() {
	if err := Cmd.Execute(); err != nil {
		fmt.Println(aurora.Red(err.Error()))
		os.Exit(codeFailure)
	}
}
