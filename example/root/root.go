package main

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/x-mod/cmd"
)

func main() {
	cmd.Add(
		cmd.Name("root"),
		cmd.Main(RootMain),
	)
	cmd.PersistentFlags().StringP("parameter", "p", "test", "flags usage")
	cmd.Execute()
}

func RootMain(c *cmd.Command, args []string) error {
	fmt.Println("my root command running ...parameter:", viper.GetString("parameter"))
	return nil
}
