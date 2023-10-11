package main

import(
	"fmt"
	"encoding/csv"
	"os"
)


func runGame( records [][]string) ( int) {
	var userGuess string
	var answer string
	var score int

	for _, record := range records {
		answer = record[1]
		fmt.Printf("%s =  ",record[0])
		fmt.Scanln(&userGuess)

		if userGuess == answer{
			score += 1
		}
	}
	return score

} 



func main(){
	var duration int
	fmt.Print("Timer duration: ")
	fmt.Scanln(&duration)

	file, err := os.Open("problems.csv")
	if err != nil{
		fmt.Println("error")
	}

	fileReader := csv.NewReader(file)
	records,err:= fileReader.ReadAll()

	if err != nil {
        fmt.Println("Error:", err)
        return
    }

	score := runGame(records)



	fmt.Printf("Your score is %d \n", score)

}

