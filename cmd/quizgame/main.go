package main

import (
	"fmt"

	"github.com/mamorski/gophercises"
	"github.com/mamorski/gophercises/internal/quizgame"
)

func main() {
	fmt.Println("testing structure")
	fmt.Println(gophercises.Config())
	fmt.Println(quizgame.RunQuiz())
}
