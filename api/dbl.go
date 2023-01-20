package api

import (
	"bytes"
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/database"
	"cutiecat6778/discordbot/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
		CheckRedirect: redirectPolicyFunc,
	}

	return DBL{
		HttpClient: &client,
	}
}

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	req.Header.Add("Authorization", class.DBLKey)
	return nil
}

func ListenVotes() {
	http.HandleFunc("/dbl", WebhookHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to new server!")
	})

	class.Ignore = false

	utils.Debug.Println(http.ListenAndServe(":3000", nil))
	utils.Debug.Println("Serving dbl")
}

func (handler DBL) PostStats(ServerCount int) error {

	reqBody, err := json.Marshal(map[string]int{
		"server_count": ServerCount,
	})
	if err != nil {
		print(err)
	}

	url := fmt.Sprintf("https://top.gg/api/bots/%v/stats", class.BotID)
	utils.Debug.Println(url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	req.Header.Add("Authorization", class.DBLKey)
	req.Header.Add("Content-Type", "application/json")
	resp, err := handler.HttpClient.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		utils.HandleServerError(err)
	}

	resp.Body.Close()

	utils.Debug.Println(resp.StatusCode, body)

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

	utils.Debug.Println("User voted ", t.User)

	f := database.UserVoted(t.User)
	if !f {
		utils.HandleServerError(errors.New("user failed to update their tokens counts"))
	}
}
