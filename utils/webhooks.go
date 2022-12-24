package utils

import (
	"bytes"
	"cutiecat6778/discordbot/class"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func SendErrorMessage(c string, e string) {

	current_time := time.Now()

	requestBody, _ := json.Marshal(map[string]string{
		"content": "**" + current_time.Format("2006-01-02 15:04:05") + "**\n" + c + "\n```\n" + e + "\n```\n",
	})

	_, err := http.Post(class.HookURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal("Error while trying to send webhook message")
		panic(err)
	}
}

func SendLogMessage(c string) {

	current_time := time.Now()

	requestBody, _ := json.Marshal(map[string]string{
		"content": "**" + current_time.Format("2006-01-02 15:04:05") + "**\n" + c,
	})

	_, err := http.Post(class.HookURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal("Error while trying to send webhook message: ", err)
		panic(err)
	}
}

func SendMessage(c string) {
	requestBody, _ := json.Marshal(map[string]string{
		"content": c,
	})

	_, err := http.Post(class.HookURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal("Error while trying to send webhook message")
		panic(err)
	}
}
