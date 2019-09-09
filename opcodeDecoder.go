package main

import (
	"fmt"
	"log"
	"math/rand"
)

func (chip8 *Chip8) decodeOpcode(opcode uint16) {

	// TODO: more opcodes
	switch opcode & 0xF000 {
	case 0x0000:
		switch opcode & 0x000F {
		case 0x0000:
			chip8.clearDisplay()
			break
		case 0x000E:
			chip8.stackPointer--
			chip8.programCounter = chip8.stack[chip8.stackPointer]
			break
		}
		break

	case 0x1000:
		chip8.programCounter = getNibblesFromTwoBytes(opcode, 1, 4)
		break

	case 0x2000:
		chip8.stack[chip8.stackPointer] = chip8.programCounter
		chip8.stackPointer++
		chip8.programCounter = opcode & 0x0FFF
		break

	case 0x3000:
		registerIndexInTwoBytes := (opcode & 0x0F00) >> 8
		registerIndex := byte(registerIndexInTwoBytes)
		registerValue := chip8.V[registerIndex]
		targetValue := byte(opcode & 0x00FF)

		if registerValue == targetValue {
			chip8.programCounter += 4
		} else {
			chip8.programCounter += 2
		}
		break

	case 0x4000:
		registerIndexInTwoBytes := (opcode & 0x0F00) >> 8
		registerIndex := byte(registerIndexInTwoBytes)
		registerValue := chip8.V[registerIndex]
		targetValue := byte(opcode & 0x00FF)

		if registerValue != targetValue {
			chip8.programCounter += 4
		} else {
			chip8.programCounter += 2
		}
		break

	case 0x5000:
		registerIndex1 := (opcode & 0x0F00) >> 8
		registerIndex2 := (opcode & 0x00F0) >> 4
		if chip8.V[registerIndex1] == chip8.V[registerIndex2] {
			chip8.programCounter += 4
		} else {
			chip8.programCounter += 2
		}
		break

	case 0x6000:
		registerIndex := (opcode & 0x0F00) >> 8
		chip8.V[registerIndex] = byte(opcode & 0x00FF)
		chip8.programCounter += 2
		break

	case 0x7000:
		break

	case 0x8000:
		chip8.logicOperations(opcode)
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
		chip8.programCounter = getNibblesFromTwoBytes(opcode, 1, 4) + uint16(chip8.V[0])
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

func (chip8 *Chip8) logicOperations(opcode uint16) {
	registerIndex1 := getNibblesFromTwoBytes(opcode, 1, 2)
	registerIndex2 := getNibblesFromTwoBytes(opcode, 2, 3)

	switch opcode & 0x000F {
	case 0x0000:
		chip8.V[registerIndex1] = chip8.V[registerIndex2]
		break

	case 0x0001:
		chip8.V[registerIndex1] = chip8.V[registerIndex1] | chip8.V[registerIndex2]
		break

	case 0x0002:
		chip8.V[registerIndex1] = chip8.V[registerIndex1] & chip8.V[registerIndex2]
		break

	case 0x0003:
		chip8.V[registerIndex1] = chip8.V[registerIndex1] ^ chip8.V[registerIndex2]
		break

	case 0x0004:
		if chip8.V[registerIndex1]+chip8.V[registerIndex2] <= 255 {
			chip8.V[15] = 1
		} else {
			chip8.V[15] = 0
		}
		chip8.V[registerIndex1] += chip8.V[registerIndex2]
		break

	case 0x0005:
		if chip8.V[registerIndex1] < chip8.V[registerIndex2] {
			chip8.V[15] = 0
		} else {
			chip8.V[15] = 1
		}
		chip8.V[registerIndex1] -= chip8.V[registerIndex2]
		break

	case 0x0006:
		chip8.V[15] = chip8.V[registerIndex1] % 2
		chip8.V[registerIndex1] = chip8.V[registerIndex1] >> 1
		break

	case 0x0007:
		if chip8.V[registerIndex1] > chip8.V[registerIndex2] {
			chip8.V[15] = 0
		} else {
			chip8.V[15] = 1
		}
		chip8.V[registerIndex1] = chip8.V[registerIndex2] - chip8.V[registerIndex1]
		break

	case 0x0008:
		if chip8.V[registerIndex1] >= 128 {
			chip8.V[15] = 1
		} else {
			chip8.V[15] = 0
		}
		chip8.V[registerIndex1] = chip8.V[registerIndex1] << 1
		break
	}
}

func getNibblesFromTwoBytes(twoBytes uint16, firstNibbleIndex int, lastNibbleIndex int) uint16 {
	var mask uint16 = 0xFFFF
	for i := 0; i < firstNibbleIndex; i++ {
		mask = mask >> 4
	}
	return (twoBytes & mask) >> (4 * (4 - lastNibbleIndex))
}
