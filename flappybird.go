package main

import (
	"github.com/Tohie/GoFlappyBird/bird"
	"github.com/Tohie/GoFlappyBird/pipe"
	"github.com/Tohie/GoFlappyBird/utils"
	"github.com/banthar/Go-SDL/sdl"
	"github.com/banthar/Go-SDL/ttf"
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
	"github.com/go-gl/gltext"
	"strconv"
	"time"
)

func main() {
	sdl.Init(sdl.INIT_VIDEO)
	ttf.Init()
	screen := sdl.SetVideoMode(480, 560, 16, sdl.OPENGL|sdl.RESIZABLE)
	sdl.WM_SetCaption("Flappy Bird", "")
	bg := utils.TextureFromFile("./bg.png")
	font := utils.LoadFont("/usr/share/fonts/truetype/DroidSans.ttf", 32)
	reshape(int(screen.W), int(screen.H))

	renderBackground(screen, bg)
	font.Printf(110, 50, "Click to play")
	sdl.GL_SwapBuffers()
	for {
	OuterLoop:
		for {
			e := sdl.WaitEvent()
			switch e.(type) {
			case *sdl.MouseButtonEvent:
				if e.(*sdl.MouseButtonEvent).Type == sdl.MOUSEBUTTONUP {
					break OuterLoop
				}
			}
		}

		score, quit := playGame(screen, bg, font)
		if quit {
			break
		}

		quit = gameOverScreen(screen, strconv.Itoa(score), bg, font)
		if quit {
			break
		}
	}

	screen.Free()
	ttf.Quit()
	sdl.Quit()
	return
}

func playGame(screen *sdl.Surface, bg *glh.Texture, font *gltext.Font) (int, bool) {
	movingUp := false
	b := bird.NewBird(240, 380)
	score := 0
	var pipes []*pipe.Pipe
	ticker := 0
	for {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		quit := manageEvents(screen, &movingUp)
		if quit {
			return score, true
		}

		renderBackground(screen, bg)
		ticker++
		if ticker > 100 {
			pipes = append(pipes, pipe.NewPipe(int(screen.W), int(screen.H)))
			ticker = 0
		}

		if movingUp {
			b.MoveUp()
		} else {
			b.MoveDown()
		}

		for _, p := range pipes {
			if p.X < 0 {
				p.Destroy()
				pipes = pipes[1:]
			}
			if p.CollidesWith(b.X, b.Y, b.Width, b.Height) {
				return score, false
			}
			if p.GonePast(b.X, b.Y) {
				score++
			}
			p.Tick()
			p.Render()
		}
		b.Render()

		font.Printf(240, 50, strconv.Itoa(score))
		screen.Flip()
		time.Sleep((1 / 30) * time.Second)
		sdl.GL_SwapBuffers()
	}
}

func gameOverScreen(
	screen *sdl.Surface,
	score string,
	bg *glh.Texture,
	font *gltext.Font) bool {

	for {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch e.(type) {
			case *sdl.ResizeEvent:
				resize(screen, e.(*sdl.ResizeEvent))
			case *sdl.QuitEvent:
				return true
			case *sdl.MouseButtonEvent:
				return false
			}
		}
		renderBackground(screen, bg)
		font.Printf(110, 50, "Game Over")
		font.Printf(110, 100, "Your score: "+score)
		font.Printf(110, 150, "Click to play again")
		sdl.GL_SwapBuffers()
		time.Sleep((1 / 30) * time.Second)
	}
	return false
}

func manageEvents(screen *sdl.Surface, movingUp *bool) bool {
	for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
		switch e.(type) {
		case *sdl.ResizeEvent:
			re := e.(*sdl.ResizeEvent)
			resize(screen, re)
			return false

		case *sdl.MouseButtonEvent:
			*movingUp = !*movingUp
			return false
		case *sdl.QuitEvent:
			return true
		}
	}
	return false
}

func resize(screen *sdl.Surface, e *sdl.ResizeEvent) {
	screen = sdl.SetVideoMode(int(e.W), int(e.H), 16,
		sdl.OPENGL|sdl.RESIZABLE)
	if screen != nil {
		reshape(int(screen.W), int(screen.H))
	} else {
		panic("Couldn't set the new video mode??")
	}
}

func reshape(w, h int) {
	// Write to both buffers, prevent flickering
	gl.DrawBuffer(gl.FRONT_AND_BACK)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Viewport(0, 0, w, h)
	gl.Ortho(0, float64(w), float64(h), 0, -1.0, 1.0)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}

func renderBackground(screen *sdl.Surface, bg *glh.Texture) {
	utils.TexturedQuad(bg, 0, 0, int(screen.W), int(screen.H))
}
