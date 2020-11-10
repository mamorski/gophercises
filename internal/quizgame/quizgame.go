package quizgame

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Args - Arguments structure
type Args struct {
	path    string
	timeout int
	shuffle bool
}

// RunQuiz ...
func RunQuiz(a Args) string {

	f, err := os.Open(a.path)

	if err != nil {
		log.Fatalln("An error occurred when trying to open a CSV file,", err)
	}

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	correctAnswers := 0

	if a.shuffle {
		records = shuffle(records)
	}

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

// InitArgs - initiate arguments
func InitArgs() *Args {

	filename := flag.String("filename", "problems.csv", "Full path to the csv file.")
	shuffle := flag.Bool("shuffle", false, "Shuffle the quiz.")
	limit := flag.Int("timeout", 30, "Time limit per question.")
	flag.Parse()
	a := Args{path: *filename, timeout: *limit, shuffle: *shuffle}

	return &a
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

func shuffle(records [][]string) [][]string {
	fmt.Println("Shuffle called!")
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range records {
		np := r.Intn(len(records) - 1)
		records[i], records[np] = records[np], records[i]
	}

	return records
}
