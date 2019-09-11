package main

import (
	"fmt"
	"strings"
)

type graphicsDrawer struct {
	width, height int
	defaultFiller string
	emptySpace    string
}

func (graphicsDrawer *graphicsDrawer) clearAndDraw(graphics []byte) {
	graphicsDrawer.clearScreen()
	graphicsDrawer.setCursorPosition(0, 0)
	graphicsDrawer.draw(graphics)
}

func (graphicsDrawer graphicsDrawer) clearScreen() {
	fmt.Print("\033[2J") //
}

func (graphicsDrawer graphicsDrawer) setCursorPosition(row, col int) {
	fmt.Printf("\033[%d;%dH", row, col)
}

func (graphicsDrawer *graphicsDrawer) draw(graphics []byte) {
	stringBuilder := strings.Builder{}
	for rowIndex := 0; rowIndex < graphicsDrawer.height; rowIndex++ {
		for colIndex := 0; rowIndex < graphicsDrawer.width; colIndex++ {
			if graphics[rowIndex*graphicsDrawer.width+colIndex] == 0 {
				stringBuilder.WriteString(graphicsDrawer.emptySpace)
			} else {
				stringBuilder.WriteString(graphicsDrawer.defaultFiller)
			}
		}
		stringBuilder.WriteString("\n")
	}
	fmt.Print(stringBuilder.String())
}
