package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"nimrody.com/toypascal/v2/assembler/isa"
)

// Disassemble reads bytecode and prints the assembly instructions.
func Disassemble(inputFile string) error {
	bytecode, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("could not read input file: %w", err)
	}

	fmt.Printf("; Disassembly of %s\n", inputFile)
	fmt.Printf("; File size: %d bytes\n\n", len(bytecode))

	pc := 0 // Program Counter
	for pc < len(bytecode) {
		opcode := bytecode[pc]
		mnemonic, exists := isa.OpcodeToMnemonic[opcode]
		if !exists {
			return fmt.Errorf("address %d: unknown opcode 0x%02X", pc, opcode)
		}

		// Print the current address and mnemonic
		fmt.Printf("%04d: %s", pc, mnemonic)

		// Advance pc past the opcode
		pc++

		// Handle arguments based on the instruction
		switch opcode {
		case isa.OpPushConst, isa.OpLoadGlobal, isa.OpStoreGlobal, isa.OpLoadLocal, isa.OpStoreLocal, isa.OpJump, isa.OpJumpIfFalse, isa.OpNew:
			if pc+4 > len(bytecode) {
				return fmt.Errorf("address %d: unexpected end of file for %s argument", pc, mnemonic)
			}
			arg := isa.BytesToInt(bytecode[pc : pc+4])
			fmt.Printf(" %d\n", arg)
			pc += 4 // Move past the 4-byte argument

		case isa.OpCall:
			if pc+8 > len(bytecode) {
				return fmt.Errorf("address %d: unexpected end of file for %s arguments", pc, mnemonic)
			}
			// First argument: address or built-in ID
			addr := isa.BytesToInt(bytecode[pc : pc+4])
			pc += 4
			// Second argument: number of arguments
			numArgs := isa.BytesToInt(bytecode[pc : pc+4])
			pc += 4
			fmt.Printf(" %d %d\n", addr, numArgs)

		default:
			// No arguments for this opcode
			fmt.Println()
		}
	}

	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run disasm.go <input_file.bin>")
		os.Exit(1)
	}

	inputFile := os.Args[1]

	err := Disassemble(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
