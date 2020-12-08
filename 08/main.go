package zeroeight

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

type bootCode struct {
	accumulator          int
	instructions         []instruction
	sp                   int
	executedInstructions map[int]struct{}
}

type instruction struct {
	op  string
	arg int
}

const (
	opNOP = "nop"
	opACC = "acc"
	opJMP = "jmp"
)

var errInfLoop = errors.New("inf loop")

func (bc *bootCode) run() error {
	for {
		if bc.sp >= len(bc.instructions) {
			break
		}
		if _, exists := bc.executedInstructions[bc.sp]; exists {
			return errInfLoop
		}

		instruction := bc.instructions[bc.sp]
		bc.executedInstructions[bc.sp] = struct{}{}

		switch instruction.op {
		case opNOP:
			bc.sp++
		case opACC:
			bc.accumulator += instruction.arg
			bc.sp++
		case opJMP:
			bc.sp += instruction.arg
		}
	}

	return nil
}

func (bc *bootCode) fixInfLoop() {
	for i, inst := range bc.instructions {
		if inst.op == opJMP {
			bc.instructions[i].op = opNOP
		} else if inst.op == opNOP {
			bc.instructions[i].op = opJMP
		}

		if err := bc.run(); err == nil {
			return
		}
		bc.reset()
		bc.instructions[i].op = inst.op
	}
}

func (bc *bootCode) reset() {
	bc.sp = 0
	bc.accumulator = 0
	bc.executedInstructions = make(map[int]struct{})
}

func newBootCode(reader io.Reader) (*bootCode, error) {
	instructions := make([]instruction, 0)
	s := bufio.NewScanner(reader)
	for s.Scan() {
		instructionParts := strings.Split(s.Text(), " ")
		arg, err := strconv.Atoi(instructionParts[1])
		if err != nil {
			return nil, err
		}

		instructions = append(instructions, instruction{
			op:  instructionParts[0],
			arg: arg,
		})
	}

	return &bootCode{instructions: instructions, executedInstructions: make(map[int]struct{})}, nil
}
