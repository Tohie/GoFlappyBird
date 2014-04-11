package bird

import (
	"github.com/Tohie/GoFlappyBird/utils"
	"github.com/go-gl/glh"
)

type Bird struct {
	X, Y          int
	Width, Height int
	Speed         int
	texture       *glh.Texture
}

func NewBird(x, y int) *Bird {
	texFile := "./head.png"
	t := utils.TextureFromFile(texFile)
	return &Bird{
		x, y,
		40, 40,
		3,
		t,
	}
}

func (b *Bird) MoveUp() {
	b.Y -= b.Speed
}

func (b *Bird) MoveDown() {
	b.Y += b.Speed
}

func (b *Bird) Render() {
	utils.TexturedQuad(b.texture, b.X, b.Y, b.Width, b.Height)
}
