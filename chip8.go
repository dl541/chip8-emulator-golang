package main

type chip8 struct {
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
}

func (chip8 chip8) initialize() {
	chip8.pc = 0x200
	chip8.opcode = 0
	chip8.I = 0
	chip8.sp = 0
}
