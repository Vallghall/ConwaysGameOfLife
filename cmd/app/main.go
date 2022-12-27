package main

import (
	"github.com/Vallghall/golmie/cmd"
	"github.com/Vallghall/golmie/pkg/grid"
	"github.com/Vallghall/golmie/pkg/mode"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"log"
)

type Game struct {
	mode     mode.Mode
	cd       int
	cellSize int
	grid     *grid.Grid

	deadImg  *ebiten.Image
	aliveImg *ebiten.Image
}

func (g *Game) Start() {
	g.mode = mode.Started
}

func (g *Game) Pause() {
	g.mode = mode.Paused
}

func (g *Game) Img(cellAlive bool) *ebiten.Image {
	if cellAlive {
		return g.aliveImg
	} else {
		return g.deadImg
	}
}

func (g *Game) Update() error {
	switch g.mode {
	case mode.Initial:
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			mx, my := ebiten.CursorPosition()
			posX, posY := mx/g.cellSize, my/g.cellSize
			g.grid.Rows[posX][posY].Alive = !g.grid.Rows[posX][posY].Alive
		}

		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.Start()
		}
	case mode.Started:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.Pause()
			return nil
		}

		if g.cd == 0 {
			g.grid.NextGeneration()
			g.cd = 10
		} else {
			g.cd--
		}
		return nil
	case mode.Paused:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.Start()
			return nil
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i, row := range g.grid.Rows {
		for j, cell := range row {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(cell.Size*i), float64(cell.Size*j))
			screen.DrawImage(g.Img(cell.Alive), op)
		}
	}
}

func (g *Game) Layout(w, h int) (rw, rh int) {
	return ebiten.WindowSize()
}

func main() {
	configs := cmd.NewConfigs()
	
	ebiten.SetTPS(ebiten.SyncWithFPS)
	ebiten.SetWindowSize(configs.Window.Width, configs.Window.Height)
	ebiten.SetWindowTitle(configs.Window.Title)
	log.Fatalln(ebiten.RunGame(&Game{
		cellSize: configs.Cell.Size,
		grid:     grid.New(configs.Cell.Size),
		deadImg:  configs.DeadSprite(),
		aliveImg: configs.AliveSprite(),
	}))
}
