package main

import (
	"fmt"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	//"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fmt.Println("Let the games begin!")

	a := app.New()
	w := a.NewWindow("Dice Bag")
	l := [6]fyne.Canvas{
		w.Canvas(),
		w.Canvas(),
		w.Canvas(),
		w.Canvas(),
		w.Canvas(),
		w.Canvas(),
	}
	initCanvas(l)
	tabs := container.NewAppTabs(
		container.NewTabItem("D4", l[0].Content()),
		container.NewTabItem("D6", l[1].Content()),
		container.NewTabItem("D6", l[2].Content()),
		container.NewTabItem("D10", l[3].Content()),
		container.NewTabItem("D12", l[4].Content()),
		container.NewTabItem("D20", l[5].Content()),
	)

	tabs.SetTabLocation(container.TabLocationLeading)
	w.SetContent(tabs)

	w.Resize(fyne.NewSize(400, 400))
	w.ShowAndRun()

}

func initCanvas(l [6]fyne.Canvas) {
	for index, canvas := range l {
		switch num := index; num {
		case 0:
			setCanvasContent(canvas, 4)
		case 1:
			setCanvasContent(canvas, 6)
		case 2:
			setCanvasContent(canvas, 8)
		case 3:
			setCanvasContent(canvas, 10)
		case 4:
			setCanvasContent(canvas, 12)
		case 5:
			setCanvasContent(canvas, 20)
		default:
			setCanvasContent(canvas, 6)
		}
	}
}

func setCanvasContent(c fyne.Canvas, i int) {
	img := canvas.NewImageFromFile(fmt.Sprintf("imgs/d%d/1.png", i))
	img.FillMode = canvas.ImageFillOriginal
	btn := widget.NewButton("Roll", func() {
		rand := rand.Intn(i) + 1
		img.File = fmt.Sprintf("imgs/d%d/%d.png", i, rand)
		img.Refresh()
	})

	c.SetContent(
		container.NewVBox(
			img,
			btn,
		),
	)

}
