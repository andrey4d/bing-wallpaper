/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package cli

import (
	"bing-wallpaper/internal/handlers"
	"bing-wallpaper/internal/wallpaper"
	"os"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get image from bing.com",
	Run:   get,
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func get(cmd *cobra.Command, args []string) {

}

func d() {
	wallpaper, err := wallpaper.NewWallpaper("UHD", "bing-wallpapers")
	handlers.CheckError(err, "wallpaper.NewWallpaper(")

	if err := wallpaper.DownloadAndSave(); err != nil {
		if os.IsExist(err) {
			return
		}
		handlers.CheckError(err, "DownloadAndSave()")
	}

}
