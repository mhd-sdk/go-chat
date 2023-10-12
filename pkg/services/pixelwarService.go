package services

import (
	"github.com/mhd-sdk/go-chat/pkg/models"
)

var PixelMatrix = models.PixelMatrix{}

func InitPixelMatrix() {
	PixelMatrix = make([][]models.Pixel, 5)
	for i := range PixelMatrix {
		PixelMatrix[i] = make([]models.Pixel, 5)
		for j := range PixelMatrix[i] {
			PixelMatrix[i][j].Color = "white"
			PixelMatrix[i][j].X = i
			PixelMatrix[i][j].Y = j
		}
	}
}

func ChangePixelColor(x int, y int, color string) {
	PixelMatrix[x][y].Color = color
	UpdateClients()
}
