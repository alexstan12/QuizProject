package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type csvData struct {
	question string
	answer string
}

func main(){
	var score int = 0; 
	data := make([]csvData, 0, 10);
	reader := bufio.NewReader(os.Stdin)
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")

	csvFile, err:= os.OpenFile("src/github.com/alexstan12/QuizProblem/questionaire.csv", os.O_RDWR, 0755)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Successfully opened CSV file")
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err !=nil {
		fmt.Println(err)
	}
	for _,line := range csvLines {
			data = append(data,  csvData{
				question: line[0],
				answer: line[1],
			})
	}
timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
//fmt.Println(data)
problemloop:
	for i, line := range data {
		answerCh := make(chan string)
		go func(){
			fmt.Println("Question " + strconv.Itoa(i) + " " + line.question);
			fmt.Print("Enter your answer: ")
			answer, _ := reader.ReadString('\n')
			answer = strings.TrimSpace(answer);
			answerCh <- answer
		}()

		select{
			case <-timer.C:
				fmt.Println("Stopped on " + strconv.Itoa(i) + " iteration")
				break problemloop 

			case answer:= <-answerCh:
				if answer == line.answer {
					score++
				}
		}

	}

fmt.Println("Your score is: " + strconv.Itoa(score))

}
