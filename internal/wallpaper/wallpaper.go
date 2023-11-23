/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package wallpaper

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
)

const bingUrl = "https://www.bing.com"

type Wallpaper struct {
	url        string
	resolution string // ("1920x1200", "1920x1080", "UHD")
	saveDir    string
}

func NewWallpaper(resolutions string, saveDir string) *Wallpaper {
	return &Wallpaper{
		resolution: resolutions,
		saveDir:    saveDir,
	}
}

func (w Wallpaper) String() string {
	return fmt.Sprintf("url : %s, resolution: %s, save: %s", w.url, w.resolution, w.saveDir)
}

func (w *Wallpaper) SetUrl(url string) {
	w.url = url
}

func (w *Wallpaper) SetResolutions(resolutions string) {
	w.resolution = resolutions
}

func (w *Wallpaper) SetSaveDir(saveDir string) {
	w.saveDir = saveDir
}

func (w Wallpaper) GetUrl() string {
	return w.url
}

func (w Wallpaper) GetResolutions() string {
	return w.resolution
}

func (w Wallpaper) GetSaveDir() string {
	return w.saveDir
}

func (w *Wallpaper) GetWallpaperImageUrl() (string, error) {
	resp, err := http.Get(bingUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	log.Println("Response status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile(`th\\?id=(.*?_tmb.jpg)`)
	urlGroup := re.FindStringSubmatch(string(body))
	url := urlGroup[len(urlGroup)-1]

	re = regexp.MustCompile("tmb")
	w.url = re.ReplaceAllLiteralString(url, w.resolution)

	return w.url, nil
}

func (w Wallpaper) Download() ([]byte, error) {
	resp, err := http.Get(w.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	log.Println("Response status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (w Wallpaper) Save() error {

}
