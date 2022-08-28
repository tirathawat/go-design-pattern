package main

import (
	"design-pattern/behavioral/state/canvas"
	"design-pattern/behavioral/state/tool"
)

func main() {
	canvas := canvas.New(tool.NewSelection())
	canvas.MouseDown()
	canvas.MouseUp()

	canvas.CurrentTool = tool.NewErase()
	canvas.MouseDown()
	canvas.MouseUp()

	canvas.CurrentTool = tool.NewBrush()
	canvas.MouseDown()
	canvas.MouseUp()
}
