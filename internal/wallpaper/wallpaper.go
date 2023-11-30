/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package wallpaper

import (
	"bing-wallpaper/internal/handlers"
	"bing-wallpaper/internal/loggers"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"regexp"
)

const bingUrl = "https://www.bing.com"

var infoLogger = loggers.NewInfoLogger(os.Stdout)

type Wallpaper struct {
	url        string
	resolution string // ("1920x1200", "1920x1080", "UHD")
	saveDir    string
	fileName   string
}

func NewWallpaper(resolutions string, saveDir string) (*Wallpaper, error) {

	dir, err := handlers.GetAbsPath(saveDir)

	if err != nil {
		return nil, err
	}

	return &Wallpaper{
		resolution: resolutions,
		saveDir:    dir,
	}, nil
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

	if resp.StatusCode != 200 {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile("th\\?id=(.*?_tmb.jpg)")
	urlGroup := re.FindStringSubmatch(string(body))

	re = regexp.MustCompile("tmb")
	w.url = fmt.Sprintf("http://bing.com/%s", re.ReplaceAllLiteralString(urlGroup[0], w.resolution))
	w.fileName = fmt.Sprintf("%s/%s", w.saveDir, re.ReplaceAllLiteralString(urlGroup[len(urlGroup)-1], w.resolution))

	if _, err := os.Stat(w.fileName); err == nil {
		infoLogger.Printf("File %s exists.\n", w.fileName)
		return w.url, fs.ErrExist
	}

	return w.url, nil
}

func (w Wallpaper) Download() ([]byte, error) {
	infoLogger.Println(w.url)
	resp, err := http.Get(w.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (w Wallpaper) Save(image []byte) error {

	if err := handlers.MakeDirAllIfNotExists(w.saveDir, 0755); err != nil {
		return err
	}

	if err := os.WriteFile(w.fileName, image, 0644); err != nil {
		return err
	}
	infoLogger.Printf("Save wallpaper %s\n.", w.fileName)

	return nil
}

func (w Wallpaper) DownloadAndSave() error {

	if _, err := w.GetWallpaperImageUrl(); err != nil {
		return err
	}

	tmp, err := w.Download()
	if err != nil {
		return err
	}

	if err := w.Save(tmp); err != nil {
		return err
	}
	return nil
}
