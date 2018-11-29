package ext

import "path/filepath"

// GIF - Extension use for GIF file
const GIF = ".gif"

// IsJPEG - Return true file extension is JPEG
func IsJPEG(f string) bool {
	ext := filepath.Ext(f)
	return ext == ".jpeg" || ext == ".jpg" || ext == ".JPEG" || ext == ".JPG"
}

// IsPNG - Return true file extension is PNG
func IsPNG(f string) bool {
	ext := filepath.Ext(f)
	return ext == ".png" || ext == ".PNG"
}
