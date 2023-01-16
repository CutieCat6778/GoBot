package api

import (
	"bytes"
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/database"
	"cutiecat6778/discordbot/utils"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type DBL struct {
	HttpClient *http.Client
}

var (
	DBLURL string = "https://top.gg/api/%v"
)

func NewDBL() DBL {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			r.Header.Add("Authorization", class.DBLKey)
			return nil
		},
	}

	http.HandleFunc("/dbl", WebhookHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to new server!")
	})

	log.Println(http.ListenAndServe(":3000", nil))
	log.Println("Serving dbl")

	return DBL{
		HttpClient: &client,
	}
}

func (handler DBL) PostStats(ServerCount int) error {

	reqBody, err := json.Marshal(map[string]string{
		"server_count": fmt.Sprintf("%v", ServerCount),
	})
	if err != nil {
		print(err)
	}

	url := fmt.Sprintf(DBLURL, class.BotID+"/stats")
	_, err = handler.HttpClient.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	return nil
}

type WebhookResp struct {
	Bot       string "json:bot"
	User      string "json:user"
	Type      string "json:type"
	isWeekend bool   "json:isWeekend"
	query     string "json:query"
}

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" || r.Header.Get("authorization") != "thinyisgood" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var t WebhookResp
	err := decoder.Decode(&t)
	if err != nil {
		utils.HandleServerError(err)
	}

	log.Println("User voted ", t.User)

	f := database.UserVoted(t.User)
	if !f {
		utils.HandleServerError(errors.New("user failed to update their tokens counts"))
	}
}
