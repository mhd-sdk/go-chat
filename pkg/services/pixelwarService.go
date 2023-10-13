package services

import (
	"github.com/mhd-sdk/go-chat/pkg/models"
)

var PixelMatrix = models.PixelMatrix{}
var paneSize = 50

func InitPixelMatrix() {
	PixelMatrix = make([][]models.Pixel, paneSize)
	for i := range PixelMatrix {
		PixelMatrix[i] = make([]models.Pixel, paneSize)
		for j := range PixelMatrix[i] {
			PixelMatrix[i][j].Color = "white"
			PixelMatrix[i][j].X = i
			PixelMatrix[i][j].Y = j
		}
	}
}

func ChangePixelColor(x int, y int, color string) {
	Mu.Lock()
	PixelMatrix[x][y].Color = color
	Mu.Unlock()
	UpdateClients()
}
