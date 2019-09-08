package main

func main() {
	chip8 := Chip8{}
	chip8.initialize()

	for {
		chip8.emulateCycle()

		if chip8.drawFlag {
			chip8.drawGraphics()
		}

		chip8.setKeys()
	}

}
