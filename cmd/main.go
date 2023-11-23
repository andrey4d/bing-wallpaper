/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package main

import (
	"bing-walpaper/internal/handlers"
	"bing-walpaper/internal/wallpaper"
	"os"
)

func main() {

	wallpaper, err := wallpaper.NewWallpaper("UHD", "bing-wallpapers")
	handlers.CheckError(err, "wallpaper.NewWallpaper(")

	if err := wallpaper.DownloadAndSave(); err != nil {
		if os.IsExist(err) {
			return
		}
		handlers.CheckError(err, "DownloadAndSave()")
	}

}
