package main

import (
  "image/color"
  "log"
  "fmt"
  "runtime/pprof"
  "os"

  "github.com/hajimehoshi/ebiten/v2"
  _ "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
  TileImage *ebiten.Image
}

func (g *Game) Update() error {
  return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
  screen.Fill(color.RGBA{0x80, 0, 0xff, 0xff})
  options := &ebiten.DrawImageOptions{}
  options.GeoM.Translate(20, 20)
  screen.DrawImage(g.TileImage, options)
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
  return 320, 240
}

func main() {
  f, err := os.Create("main.prof")
  if err != nil {
    log.Fatal(err)
  }
  pprof.StartCPUProfile(f)
  defer pprof.StopCPUProfile()

  fmt.Println("start")
  game := Game{
    TileImage: ebiten.NewImage(64, 64),
  }
  game.TileImage.Fill(color.RGBA{0xff, 0, 0, 0xff})
  fmt.Println("init done")

  ebiten.SetWindowSize(640, 480)
  ebiten.SetWindowTitle("Hello, World!")
  fmt.Println("window init done")

  if err := ebiten.RunGame(&game); err != nil {
    log.Fatal(err)
  }
}
