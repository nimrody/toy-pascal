// Package isa defines the Instruction Set Architecture for the Toy Pascal VM.
// It provides a single source of truth for opcodes, mnemonics, and data
// conversion routines used by both the assembler and disassembler.
package isa

import "encoding/binary"

// Opcodes - These must match the VM's opcodes exactly.
const (
	OpPushConst   byte = 0x01
	OpPop         byte = 0x02
	OpAdd         byte = 0x10
	OpSub         byte = 0x11
	OpMul         byte = 0x12
	OpDiv         byte = 0x13
	OpCmpEq       byte = 0x14
	OpCmpNeq      byte = 0x15
	OpCmpLt       byte = 0x16
	OpCmpGt       byte = 0x17
	OpCmpLe       byte = 0x18
	OpCmpGe       byte = 0x19
	OpLoadGlobal  byte = 0x20
	OpStoreGlobal byte = 0x21
	OpLoadLocal   byte = 0x22
	OpStoreLocal  byte = 0x23
	OpJump        byte = 0x30
	OpJumpIfFalse byte = 0x31
	OpCall        byte = 0x40
	OpRet         byte = 0x41
	OpNew         byte = 0x50
	OpLoadIndirect  byte = 0x51
	OpStoreIndirect byte = 0x52
	OpHalt        byte = 0xFF
)

// MnemonicToOpcode maps the string representation of an instruction to its opcode.
var MnemonicToOpcode = map[string]byte{
	"PUSH_CONST":    OpPushConst,
	"POP":           OpPop,
	"ADD":           OpAdd,
	"SUB":           OpSub,
	"MUL":           OpMul,
	"DIV":           OpDiv,
	"CMP_EQ":        OpCmpEq,
	"CMP_NEQ":       OpCmpNeq,
	"CMP_LT":        OpCmpLt,
	"CMP_GT":        OpCmpGt,
	"CMP_LE":        OpCmpLe,
	"CMP_GE":        OpCmpGe,
	"LOAD_GLOBAL":   OpLoadGlobal,
	"STORE_GLOBAL":  OpStoreGlobal,
	"LOAD_LOCAL":    OpLoadLocal,
	"STORE_LOCAL":   OpStoreLocal,
	"JUMP":          OpJump,
	"JUMP_IF_FALSE": OpJumpIfFalse,
	"CALL":          OpCall,
	"RET":           OpRet,
	"NEW":           OpNew,
	"LOAD_INDIRECT": OpLoadIndirect,
	"STORE_INDIRECT":OpStoreIndirect,
	"HALT":          OpHalt,
}

// OpcodeToMnemonic is the reverse of MnemonicToOpcode, for disassembling.
var OpcodeToMnemonic = make(map[byte]string)

func init() {
	// Programmatically create the reverse map to avoid errors.
	for mnemonic, opcode := range MnemonicToOpcode {
		OpcodeToMnemonic[opcode] = mnemonic
	}
}

// IntToBytes converts an int32 to its 4-byte little-endian representation.
func IntToBytes(n int32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(n))
	return b
}

// BytesToInt converts a 4-byte little-endian slice to an int32.
func BytesToInt(b []byte) int32 {
	return int32(binary.LittleEndian.Uint32(b))
}
