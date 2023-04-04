package main

import (
	"bufio"
	"fmt"
	"googleBard/bard"
	"log"
	"os"
)

func main() {
	sessionID := ""

	b := bard.NewBard(sessionID)
	bardOptions := bard.Options{
		ConversationID: "",
		ResponseID:     "",
		ChoiceID:       "",
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("You: ")
		scanner.Scan()
		message := scanner.Text()

		response, err := b.SendMessage(message, bardOptions)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("Bard: %s\n\n\n", response.Choices[0].Answer)

		bardOptions.ConversationID = response.ConversationID
		bardOptions.ResponseID = response.ResponseID
		bardOptions.ChoiceID = response.Choices[0].ChoiceID
	}
}
