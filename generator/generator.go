package generator

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"path/filepath"
	"runtime"

	"github.com/chbes/makeagif/frame"
	"github.com/chbes/makeagif/sysio"
)

// New - Create and save a new GIF
func New(inFo string, outFo string, outFi string, w uint, h uint, d int, v bool, vv bool) error {
	var err error

	inFo, err = sysio.FixPathTidle(inFo)
	if err != nil {
		return err
	}

	outFo, err = sysio.FixPathTidle(outFo)
	if err != nil {
		return err
	}

	outFi, err = sysio.FixPathTidle(outFi)
	if err != nil {
		return err
	}

	if outFi == "" {
		outFi, err = sysio.FolderName(inFo)
		if err != nil {
			return err
		}
		if outFi == "" {
			outFi = "image"
		}
	}

	if outFo == "" {
		outFo = inFo
	}

	imgs, err := sysio.GetImages(inFo)
	if err != nil {
		return err
	}

	if v || vv {
		fmt.Println("==============================================")
		fmt.Printf("Folder load: %s\n", inFo)
		for i, j := range imgs {
			fmt.Printf("frame[%d] %s\n", i+1, j)
		}
		fmt.Println()
	}

	nbFrame := len(imgs)
	outGif := &gif.GIF{}
	outGif.Image = make([]*image.Paletted, nbFrame)
	outGif.Delay = make([]int, nbFrame)

	todo := make(chan frame.Frame, nbFrame)
	done := make(chan frame.Frame, nbFrame)

	for w := 0; w < runtime.NumCPU(); w++ {
		go worker(todo, done)
	}

	for i, j := range imgs {
		f := frame.New(j, i, w, h)
		todo <- f
		if vv {
			fmt.Printf("frame[%d]: WIP\n", f.Order+1)
		}
	}
	close(todo)

	for i := 0; i < nbFrame; i++ {
		f := <-done
		if f.NotGenerated() {
			return errors.New("frame[" + string(f.Order+1) + "]: failed")
		}
		if vv {
			fmt.Printf("frame[%d]: DONE\n", f.Order+1)
		}
		pushFrame(outGif, f, d)
	}

	if vv {
		fmt.Println()
		fmt.Printf("%s.gif: WIP\n", outFi)
	}

	pathFi := filepath.Join(outFo, outFi)
	err = sysio.SaveGIF(pathFi, outGif)
	if err != nil {
		return err
	}

	if v || vv {
		fmt.Printf("%s.gif: DONE !\n", pathFi)
		fmt.Println("==============================================")

	}

	return nil
}

// Factory - Launch GIF creation for each folder found in the input folder
func Factory(inFo string, outFo string, outFi string, w uint, h uint, d int, v bool, vv bool) error {
	fos, err := sysio.GetFolders(inFo)
	if err != nil {
		return err
	}

	for _, fo := range fos {
		err = New(inFo+"/"+fo, outFo, outFi, w, h, d, v, vv)
		if err != nil {
			return err
		}
	}
	return nil
}

func worker(todo <-chan frame.Frame, done chan<- frame.Frame) {
	for f := range todo {
		f.Generate()
		done <- f
	}
}

func pushFrame(dataGif *gif.GIF, f frame.Frame, delay int) {
	dataGif.Image[f.Order] = f.Data.(*image.Paletted)
	dataGif.Delay[f.Order] = delay
}
