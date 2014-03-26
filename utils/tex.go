package utils

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
	"image"
	_ "image/png"
	"os"
)

func TextureFromFile(fname string) *glh.Texture {
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}

	// All images will be valid
	img, _, _ := image.Decode(f)
	width := img.Bounds().Size().X
	height := img.Bounds().Size().Y
	t := glh.NewTexture(width, height)
	t.FromImage(img, 0)
	return t
}

func TexturedQuad(t *glh.Texture, x, y, w, h int) {
	glh.With(t, func() {
		gl.Enable(gl.BLEND)
		gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
		gl.Color4f(255.0, 255.0, 255.0, 1.0)
		gl.Begin(gl.TRIANGLE_FAN)

		gl.TexCoord2f(0, 0)
		gl.Vertex2i(x, y)
		gl.TexCoord2f(1, 0)
		gl.Vertex2i(x+w, y)
		gl.TexCoord2f(1, 1)
		gl.Vertex2i(x+w, y+h)
		gl.TexCoord2f(0, 1)
		gl.Vertex2i(x, y+h)

		gl.End()
	})
}
