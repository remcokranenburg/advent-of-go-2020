package main

import (
    "fmt"
    "bufio"
    "os"
    "regexp"
    "strconv"
)

type Instruction struct {
    command string
    value int64
    execCount int
}

func run(program []Instruction) (bool, int64) {
    programCounter := 0
    accumulator := int64(0)

    for programCounter < len(program) {
        instruction := &program[programCounter]

        if instruction.execCount > 0 {
            break
        }

        switch {
            case instruction.command == "acc":
            accumulator += instruction.value
            programCounter++
            instruction.execCount++

            case instruction.command == "jmp":
            programCounter += int(instruction.value)
            instruction.execCount++

            case instruction.command == "nop":
            programCounter++
            instruction.execCount++
        }
    }

    return programCounter >= len(program), accumulator
}

func fix(program *[]Instruction, i int) {
    if (*program)[i].command == "jmp" {
        (*program)[i].command = "nop"
    } else if (*program)[i].command == "nop" {
        (*program)[i].command = "jmp"
    }
}

func main() {
    shouldFix := len(os.Args) == 2 && os.Args[1] == "--fix"
    scanner := bufio.NewScanner(os.Stdin)
    program := []Instruction{}

    for scanner.Scan() {
        line := scanner.Text()
        pattern := regexp.MustCompile(`^(\w+) ([\+-]\d+)$`)
        matches := pattern.FindStringSubmatch(line)
        command := matches[1]
        value, _ := strconv.ParseInt(matches[2], 10, 0)
        program = append(program, Instruction{command, value, 0})
    }

    if shouldFix {
        for i := range program {
            fixedProgram := make([]Instruction, len(program))
            copy(fixedProgram, program)
            fix(&fixedProgram, i)
            success, accumulator := run(fixedProgram)

            if success {
                fmt.Println(accumulator)
            }
        }
    } else {
        _, accumulator := run(program)
        fmt.Println(accumulator)
    }
}
