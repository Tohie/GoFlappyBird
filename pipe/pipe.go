package pipe

import (
	"github.com/Tohie/GoFlappyBird/utils"
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
	"math/rand"
)

type Pipe struct {
	// Y is where the pipe is 'open'
	X, Y int
	// height of pipe opening
	Height       int
	Width        int
	ScreenHeight int
	Speed        int
	gonePast     bool
	pipe         *glh.Texture
	rim          *glh.Texture
}

func (p *Pipe) Tick() {
	p.X -= p.Speed
}

func NewPipe(screenWidth, screenHeight int) *Pipe {
	y := rand.Intn(screenHeight-300) + 100
	height := 100
	width := 30
	speed := 2
	pipe := utils.TextureFromFile("./pipe.png")
	rim := utils.TextureFromFile("./rim.png")

	return &Pipe{
		screenWidth, y,
		height,
		width,
		screenHeight,
		speed,
		false,
		pipe,
		rim,
	}
}

func (p *Pipe) Render() {
	gl.Color4f(0.0, 255.0, 0.0, 1.0)
	rimHeight := 20
	utils.TexturedQuad(p.pipe, p.X, 0, p.Width, p.Y-rimHeight)
	utils.TexturedQuad(p.rim, p.X, p.Y-rimHeight, p.Width, rimHeight)
	utils.TexturedQuad(p.rim, p.X, p.Y+p.Height, p.Width, rimHeight)
	utils.TexturedQuad(p.pipe, p.X, p.Y+p.Height+rimHeight, p.Width, p.ScreenHeight-p.Height)
}

func between(x, loc, locWidth int) bool {
	return x > loc && x < (loc+locWidth)
}

func (p *Pipe) CollidesWith(x, y, w, h int) bool {
	if between(x, p.X, p.Width) || between(x+w, p.X, p.Width) {
		if y < p.Y {
			return true
		}
		if y+h > p.Y+p.Height {
			return true
		}
	}
	return false
}

func (p *Pipe) GonePast(x, y int) bool {
	if !p.gonePast && p.X < x {
		p.gonePast = true
		return true
	} else {
		return false
	}
}

func (p *Pipe) Destroy() {
	p.rim.Delete()
	p.pipe.Delete()
}
