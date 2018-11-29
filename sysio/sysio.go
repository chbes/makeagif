package sysio

import (
	"errors"
	"image/gif"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"

	"../ext"
)

const mode = 0666

// GetImages - Return all JPEG or PNG file name found in folder
func GetImages(in string) ([]string, error) {
	imgs := []string{}
	fis, err := getFiles(in)
	if err != nil {
		return nil, err
	}
	for _, f := range fis {
		if ext.IsJPEG(f) || ext.IsPNG(f) {
			imgs = append(imgs, f)
		}
	}
	if len(imgs) == 0 {
		return nil, errors.New("No image in input")
	}
	return imgs, nil
}

// GetFolders - Return all folders found in folder
func GetFolders(in string) ([]string, error) {
	fos := []string{}

	contents, err := ioutil.ReadDir(in)
	if err != nil {
		return nil, err
	}
	for _, c := range contents {
		if c.IsDir() {
			fos = append(fos, c.Name())
		}
	}
	if len(fos) == 0 {
		return nil, errors.New("No folder in input")
	}
	return fos, nil
}

// FolderName - Return the last folder name of the path
func FolderName(p string) (string, error) {
	abs, err := filepath.Abs(p)
	_, folderName := path.Split(abs)
	return folderName, err
}

// SaveGIF - Save output GIF file
func SaveGIF(p string, data *gif.GIF) error {
	fo, _ := path.Split(p)
	err := createFolder(fo)
	if err != nil {
		return err
	}
	fi, err := os.OpenFile((p + ext.GIF), os.O_WRONLY|os.O_CREATE, 0664)
	defer fi.Close()
	if err != nil {
		return err
	}
	err = gif.EncodeAll(fi, data)
	return err
}

// FixPathTidle - Return the absolu path if the path contain "~" (user directory reference on linux)
func FixPathTidle(p string) (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	home := usr.HomeDir

	if p == "~" {
		p = home
	} else if strings.HasPrefix(p, "~/") {
		p = filepath.Join(home, p[2:])
	}

	return p, nil
}

func createFolder(p string) error {
	_, err := os.Stat(p)
	if validFolderName(p) && os.IsNotExist(err) {
		err = os.MkdirAll(p, 0775)
		if err != nil {
			return err
		}
	}
	return nil
}

func getFiles(in string) ([]string, error) {
	fis := []string{}
	contents, err := ioutil.ReadDir(in)
	if err != nil {
		return nil, err
	}
	for _, c := range contents {
		if !c.IsDir() {
			fis = append(fis, path.Join(in, c.Name()))
		}
	}
	return fis, nil
}

func validFolderName(p string) bool {
	return p != "." && p != ".." && p != ""
}
