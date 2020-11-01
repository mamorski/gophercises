package quizgame

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

// RunQuiz ...
func RunQuiz() string {

	fileName := "problems.csv"
	f, err := os.Open(fileName)

	if err != nil {
		log.Fatalln("An error occurred when trying to open a CSV file,", err)
	}

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	correctAnswers := 0

	for _, record := range records {
		userAnswer := askQuestion(record[0])
		if userAnswer == record[1] {
			correctAnswers++
		}
	}

	fmt.Println("Correct answers -", correctAnswers)
	fmt.Println("Total number of questions -", len(records))
	return "Done!"
}

func askQuestion(q string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(q, ": ")
	answer, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Something went wrong when trying to read the input...")
	}
	answer = strings.Replace(answer, "\r\n", "", -1)
	return answer
}
