package main

import (
	"fmt"

	"github.com/mamorski/gophercises/internal/quizgame"
)

func main() {
	a := quizgame.InitArgs()
	fmt.Println(quizgame.RunQuiz(*a))
}
