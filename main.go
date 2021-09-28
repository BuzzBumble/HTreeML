package main

import "fmt"

func test(c rune) bool {
	return c != ' '
}

func main() {
	var p Parser
	p.input = "The quick brown fox"

	//for i := 0; i < 10; i++ {
	//	fmt.Printf("Pos %d: ", p.pos)
	//	fmt.Printf("%c\n", p.consumeChar())
	//}

	fmt.Println(p.parseWord())
}
