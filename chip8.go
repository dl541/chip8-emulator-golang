package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	CLOCK_FREQUENCY = 60
)

type Chip8 struct {
	// CPU registers: V0 to VE
	// I: Index register
	// pc: Program counter
	// graphics dimension = 64*32 = 2048
	// sp: stack pointer
	opcode     uint16
	memory     [4096]byte
	V          [16]byte
	I          uint16
	pc         uint16
	gfx        [64 * 32]byte
	delayTimer uint8
	soundTimer uint8
	stack      [16]uint16
	sp         uint16
	keyboard   [16]byte

	drawFlag bool
}

func (chip8 *Chip8) initialize() {
	chip8.pc = 0x200
	chip8.opcode = 0
	chip8.I = 0
	chip8.sp = 0

	chip8.clearDisplay()
	chip8.clearStack()
	chip8.clearRegisters()
	//chip8.loadFontset()

	chip8.delayTimer = chip8.getInitialClockCount()
	chip8.soundTimer = chip8.getInitialClockCount()
}

func (chip8 *Chip8) clearDisplay() {
	var emptyScreen [64 * 32]byte
	chip8.gfx = emptyScreen
}

func (chip8 *Chip8) clearStack() {
	var emptyStack [16]uint16
	chip8.stack = emptyStack
}

func (chip8 *Chip8) clearRegisters() {
	var emptyRegisters [16]byte
	chip8.V = emptyRegisters
}

func (chip8 *Chip8) getInitialClockCount() uint8 {
	return CLOCK_FREQUENCY
}

/**
func (chip8 *Chip8) loadFontset() {
	for i := 0; i < 80; i++ {
		chip8.memory[i] = chip8_fontset[i]
	}
}
**/
func (chip8 *Chip8) loadROM(programPath string) {
	bytes, err := retrieveROM(programPath)

	if err != nil {
		log.Fatal(err)
	}

	chip8.copyROMToMemory(bytes)
}

func (chip8 *Chip8) copyROMToMemory(ROMBytes []byte) {
	for i, oneByte := range ROMBytes {
		chip8.memory[i+512] = oneByte
	}
}

func retrieveROM(filename string) ([]byte, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, statsErr
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)

	return bytes, err
}

func (chip8 *Chip8) emulateCycle() {
	opcode := chip8.fetchOpcode()
	chip8.decodeOpcode(opcode)
	chip8.updateDelayTimer()
	chip8.updateSoundTimer()
}

func (chip8 *Chip8) fetchOpcode() uint16 {
	firstByte := uint16(chip8.memory[chip8.pc])
	secondByte := uint16(chip8.memory[chip8.pc+1])
	return firstByte<<8 | secondByte
}

func (chip8 *Chip8) decodeOpcode(opcode uint16) {

	// TODO: more opcodes
	switch opcode & 0xF000 {
	case 0xA000:
		chip8.I = opcode & 0x0FFF
		chip8.pc += 2
		break

	default:
		log.Fatal(fmt.Sprintf("Unknown opcode %x", opcode))
	}

}

func (chip8 *Chip8) updateDelayTimer() {
	if chip8.delayTimer > 0 {
		chip8.delayTimer--
	}
}

func (chip8 *Chip8) updateSoundTimer() {
	if chip8.soundTimer > 0 {
		if chip8.soundTimer == 1 {
			chip8.makeSound()
		}
		chip8.soundTimer--
	}
}

func (chip8 *Chip8) makeSound() {
	fmt.Println("BEEP!")
}
