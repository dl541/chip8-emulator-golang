package main

import (
	"fmt"
	"log"
	"math/rand"
)

func (chip8 *Chip8) decodeOpcode(opcode uint16) {

	// TODO: more opcodes
	switch opcode & 0xF000 {
	case 0x2000:
		chip8.stack[chip8.sp] = chip8.pc
		chip8.sp++
		chip8.pc = opcode & 0x0FFF
		break

	case 0x3000:
		registerIndexInTwoBytes := (opcode & 0x0F00) >> 8
		registerIndex := byte(registerIndexInTwoBytes)
		registerValue := chip8.V[registerIndex]
		targetValue := byte(opcode & 0x00FF)

		if registerValue == targetValue {
			chip8.pc += 4
		} else {
			chip8.pc += 2
		}
		break

	case 0x4000:
		registerIndexInTwoBytes := (opcode & 0x0F00) >> 8
		registerIndex := byte(registerIndexInTwoBytes)
		registerValue := chip8.V[registerIndex]
		targetValue := byte(opcode & 0x00FF)

		if registerValue != targetValue {
			chip8.pc += 4
		} else {
			chip8.pc += 2
		}
		break

	case 0x5000:
		registerIndex1 := (opcode & 0x0F00) >> 8
		registerIndex2 := (opcode & 0x00F0) >> 4
		if chip8.V[registerIndex1] == chip8.V[registerIndex2] {
			chip8.pc += 4
		} else {
			chip8.pc += 2
		}
		break

	case 0x6000:
		registerIndex := (opcode & 0x0F00) >> 8
		chip8.V[registerIndex] = byte(opcode & 0x00FF)
		chip8.pc += 2
		break

	case 0x7000:
		break

	case 0x8000:
		break

	case 0x9000:
		registerIndex1 := getNibblesFromTwoBytes(opcode, 1, 2)
		registerIndex2 := getNibblesFromTwoBytes(opcode, 2, 3)
		if chip8.V[registerIndex1] != chip8.V[registerIndex2] {
			chip8.skipNextInstruction()
		} else {
			chip8.moveOnToNextInstruction()
		}
		break

	case 0xA000:
		chip8.I = opcode & 0x0FFF
		chip8.moveOnToNextInstruction()
		break

	case 0xB000:
		chip8.pc = getNibblesFromTwoBytes(opcode, 1, 4) + uint16(chip8.V[0])
		break

	case 0xC000:
		registerIndex := getNibblesFromTwoBytes(opcode, 1, 2)
		randomNumber := uint16(rand.Intn(256))
		newValue := randomNumber & getNibblesFromTwoBytes(opcode, 2, 4)
		chip8.V[registerIndex] = byte(newValue)
		break

	case 0xD000:
		break

	default:
		log.Fatal(fmt.Sprintf("Unknown opcode %x", opcode))
	}

}

func getNibblesFromTwoBytes(twoBytes uint16, firstNibbleIndex int, lastNibbleIndex int) uint16 {
	var mask uint16 = 0xFFFF
	for i := 0; i < firstNibbleIndex; i++ {
		mask = mask >> 4
	}

	return (twoBytes & mask) >> (4 * (4 - lastNibbleIndex))
}
