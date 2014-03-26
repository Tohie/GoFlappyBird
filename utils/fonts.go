package utils

import (
	"github.com/go-gl/gltext"
	"os"
)

func LoadFont(file string, scale int32) *gltext.Font {
	fd, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	defer fd.Close()

	f, err := gltext.LoadTruetype(fd, scale, 32, 127, gltext.LeftToRight)
	if err != nil {
		panic(err)
	}
	return f
}
