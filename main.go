package main

import (
  "image/color"
  "log"

  "github.com/hajimehoshi/ebiten/v2"
  _ "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {}

func (g *Game) Update() error {
  return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
  screen.Fill(color.RGBA{0x80, 0, 0xff, 0xff})
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
  return 320, 240
}

func main() {
  ebiten.SetWindowSize(640, 480)
  ebiten.SetWindowTitle("Hello, World!")
  if err := ebiten.RunGame(&Game{}); err != nil {
    log.Fatal(err)
  }
}
