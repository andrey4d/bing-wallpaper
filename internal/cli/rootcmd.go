/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package cli

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Short: "Downloading wallpaper from bing.com.",
	}
)

func Execute() error {

	return rootCmd.Execute()

}
