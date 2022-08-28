package canvas

import (
	"design-pattern/behavioral/state/tool"
)

type canvas struct {
	CurrentTool tool.Tool
}

func New(tool tool.Tool) *canvas {
	return &canvas{
		CurrentTool: tool,
	}
}

func (c *canvas) MouseDown() {
	c.CurrentTool.MouseDown()
}

func (c *canvas) MouseUp() {
	c.CurrentTool.MouseUp()
}
