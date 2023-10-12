package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)


func runGame( records [][]string, duration int) ( int) {
	var userGuess string
	var answer string
	var score int

	timer1 := time.NewTimer(time.Duration(duration) * time.Second)
	answer_channel := make(chan string)


	for _, record := range records {

		answer = record[1]
		fmt.Printf("%s =  ",record[0])

		// using a Goroutine to handle user input, which is a good approach to avoid blocking the program while waiting for user responses.
		go func(){
			fmt.Scanln(&userGuess)
			answer_channel <- userGuess
		}()

		select {
	
		case recievedAnswer := <-answer_channel:
			if recievedAnswer == answer{
				score+=1
			}

		case <-timer1.C:
			fmt.Println("Out of time!")
			
			return score

		}	
	}
	return score
} 

func main(){

	limit := flag.Int("limit", 100000, "Sets a time limit for the program to run")
	csv_file := flag.String("csv","problems.csv","a csv file in the format of 'question,answer")
	flag.Parse()


	file, err := os.Open(*csv_file)
	if err != nil{
		fmt.Printf("error, %s is not found in directory \n",*csv_file)
	}

	fileReader := csv.NewReader(file)
	records,err:= fileReader.ReadAll()

	if err != nil {
        fmt.Println("Error no file records:", err)
        return
    }

	score := runGame(records, *limit)
	fmt.Printf("Your score is %d \n", score)

}

