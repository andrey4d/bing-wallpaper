/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package cli

import (
	"bing-wallpaper/internal/handlers"
	"bing-wallpaper/internal/loggers"
	"bing-wallpaper/internal/wallpaper"
	"log"
	"os"
	"slices"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get image from bing.com",
	Run:   get,
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.PersistentFlags().String("res", "UHD", "Image resolution is in (\"1920x1200\", \"1920x1080\", \"UHD\").")
	getCmd.PersistentFlags().String("target", "./bing-wallpapers", "Target directory for wallpapers.")
}

func get(cmd *cobra.Command, args []string) {

	infoLogger := loggers.NewInfoLogger(os.Stdout)

	res, err := cmd.Flags().GetString("res")
	handlers.CheckError(err, "images resolution set err")

	target, err := cmd.Flags().GetString("target")
	handlers.CheckError(err, "target directory bad")

	w, err := wallpaper.NewWallpaper(resolution(res), "bing-wallpapers")
	handlers.CheckError(err, "wallpaper.NewWallpaper(")

	if target != "" {
		infoLogger.Printf("Save wallpaper image to %s\n", target)
		w.SetSaveDir(target)
	}

	if err := w.DownloadAndSave(); err != nil {
		if os.IsExist(err) {
			return
		}
		handlers.CheckError(err, "DownloadAndSave()")
	}

}

func resolution(resolution string) string {
	resolutions := []string{"1920x1200", "1920x1080", "UHD"}
	if slices.Contains(resolutions, resolution) {
		return resolution
	}

	log.Printf("use default image resolution UHD. can't set resolution to %s", resolution)
	return "UHD"
}
