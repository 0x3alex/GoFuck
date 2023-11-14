package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	mem    [30000]int
	ptrPos int
)

func reset() {
	mem = [30000]int{}
}

func fuckFromFile(path string) {
	f, err := os.ReadFile(path)
	if err == os.ErrNotExist {
		panic("File not found")
	}
	fuck(string(f))
}

func fuck(f string) {
	f = strings.ReplaceAll(f, "\n", "")
	f = strings.ReplaceAll(f, "\r", "")
	ops := strings.Split(f, "")
	var (
		loops       []int
		lastClosing int
	)
	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case ">":
			if ptrPos+1 < len(mem) {
				ptrPos++
			}
			break
		case "<":
			if ptrPos-1 >= 0 {
				ptrPos--
			}
			break
		case "+":
			mem[ptrPos]++
			break
		case "-":
			mem[ptrPos]--
			break
		case "[":
			//add loop start to references
			if len(loops) == 0 || loops[len(loops)-1] != i {
				loops = append(loops, i)
			}
			if mem[ptrPos] > 0 {
				continue
				//exec while
			} else {
				//jump to the end
				i = lastClosing
				loops = loops[:len(loops)-1]
			}
			break
		case "]":
			lastClosing = i
			if len(loops) == 0 {
				panic("No loop was started")
			}
			i = loops[len(loops)-1] - 1
			break
		case ".":
			fmt.Printf("%c", mem[ptrPos])
			break
		case ",":
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			mem[ptrPos] = int(input[0])
			break
		}
	}
	println("")
}
