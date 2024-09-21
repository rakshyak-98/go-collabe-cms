package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"

	"os"

	"net/http"

	"github.com/joho/godotenv"
)

// [[ Structs ]]

type JumbleWords struct {
	// keep it this way ! Don't change it :(
	Words []string `json:"body"`
}

// [[ Main function ]]
func main() {
	setEnvVariables()
	startGame()
}

// [[ Game function ]]
func setEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("[Error] Not able to load .env file")
	}
}

func startGame() {
	todayWord := getTodayWord()
	fmt.Printf("%s <- Today's wordn", todayWord)
}

func getEnvVariable(key string) string {
	return os.Getenv(key)
}

// [[ Helper function ]]

func getTodayWord() string {
	words := getWord()
	randomIndex := rand.Intn(len(words))
	jumbleWord := jumbleWord(words[randomIndex])
	return jumbleWord
}

func getWord() []string {
	// TODO: Find a way to cache this request
	request, err := http.NewRequest("GET", "https://word-generator2.p.rapidapi.com/?count=5&length=10", nil)

	if err != nil {
		fmt.Println(err)
	}

	request.Header.Add("x-rapidapi-key", getEnvVariable("RAPID_API_KEY"))
	request.Header.Add("x-rapidapi-host", getEnvVariable("RAPID_API_HOST"))

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode != http.StatusOK {
		panic("[Error] Not able to fetch words")
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	var words JumbleWords
	// new thing for me so I did it this way
	json.Unmarshal(body, &words)

	return words.Words

}

func jumbleWord(word string) string {
	runes := []rune(word)
	for i := range word {
		j := rand.Intn(i + 1)
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
