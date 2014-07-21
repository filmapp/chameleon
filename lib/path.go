package mofu

import(
	"strings"
	"strconv"
)

// @args p [string] image/Bs3iSswCIAE075C.500x500.jpg
func ParsePath(p string) (int, int, string) {
	path := strings.Split(p, "/")
	name := path[len(path) - 1]
	nameArr := strings.Split(name, ".")
	sizeParam := nameArr[1] // "foobar.100x100.jpg" => "100x100"
	width, height := ParseSizeParam(sizeParam) // "100x100" => 100, 100
	fname := nameArr[0] + "." + nameArr[2] // "foobar.100x100.jpg" => "foobar.jpg"
	path[len(path) - 1] = fname
	imgPath := strings.Join(path, "/")

	return width, height, imgPath
}

// @args p [string] 100x100
func ParseSizeParam(p string) (int, int) {
	sizeArr := strings.Split(p, "x")
	width, _ := strconv.Atoi(sizeArr[0])
	height, _ := strconv.Atoi(sizeArr[1])
	return width, height
}