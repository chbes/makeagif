package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/chbes/makeagif/generator"
)

const version = "v1.0"

func main() {
	start := time.Now()

	inFolder := flag.String("inFolder", ".", "folder contains input images")
	outFolder := flag.String("outFolder", "", "folder contains output GIF (default input folder)")
	outFile := flag.String("outFile", "", "new GIF name (default input folder name)")
	d := flag.Int("d", 50, "delay into each image, in millisecond")
	w := flag.Uint("w", 0, "GIF width (0 = ratio preserving)")
	h := flag.Uint("h", 0, "GIF height (0 = ratio preserving)")
	v := flag.Bool("v", false, "verbose mode")
	vv := flag.Bool("vv", false, "very verbose mode")
	factory := flag.Bool("factory", false, "factory mode")
	info := flag.Bool("info", false, "get informations")

	flag.Parse()

	defer printDuration(start, *vv)

	if *info {
		printVersion()
		os.Exit(1)
	}

	if *factory {
		factoryMode(*inFolder, *outFolder, *outFile, *w, *h, *d, *v, *vv)
	} else {
		singleMode(*inFolder, *outFolder, *outFile, *w, *h, *d, *v, *vv)
	}

}

func singleMode(inFolder string, outFolder string, outFile string, w uint, h uint, d int, v bool, vv bool) {
	err := generator.New(inFolder, outFolder, outFile, w, h, d, v, vv)
	catchError(err)
}

func factoryMode(inFolder string, outFolder string, outFile string, w uint, h uint, d int, v bool, vv bool) {
	err := generator.Factory(inFolder, outFolder, outFile, w, h, d, v, vv)
	catchError(err)
}

func printDuration(start time.Time, active bool) {
	if active {
		stop := time.Now()
		duration := stop.Sub(start)
		fmt.Println("duration: ", duration)
	}
}

func catchError(err error) {
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
}

func printVersion() {
	fmt.Printf("MakeAGif %s develop with ❤  by chbes\n", version)
}
