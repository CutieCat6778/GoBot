package class

import (
	"errors"
	"flag"
	"log"
)

var (
	LOCAL          bool
	Token          string
	ServerID       string
	OwnerID        string
	HookID         string
	HookToken      string
	HookURL        string
	DBKey          string
	GGAPIKey       string
	BINGKey        string
	WeatherKey     string
	NasaKey        string
	DBLKey         string
	BotID          string
	Ignore         bool
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
)

type CommandData struct {
	Permissions    int64
	Ratelimit      int64
	BotPerms       int64
	SubCommandData map[string]CommandData
}

func init() {
	LOCAL = false
	var TOKEN string
	Ignore = true
	if !LOCAL {
		TOKEN = "MTA1NTU1MzM1Mzc1NDYyODE5Nw.GbYWaN.lIF0mq9BYSDR4rSUSD8iwd3kR5YyC_6hHrN3qk"
	} else {
		TOKEN = "MTA1NzY0NjA1OTk2MDE2MDI4Ng.G1qp-u.orbBv-9teDg-Aqg4-MVx9JKx8uNcki3wpHaFLM"
	}
	DBLTOKEN := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEwNTU1NTMzNTM3NTQ2MjgxOTciLCJib3QiOnRydWUsImlhdCI6MTY3Mzg5NzkyMH0.YFmtd2FjaBpv1P2_M0tdg9U8mu4IGxuLK1yezG1ztGg"
	DBKEY := "mongodb+srv://admin:!Txzje2006@cluster0.axkhjad.mongodb.net/?retryWrites=true&w=majority"
	SERVERID := "1056177394093936760"
	OWNERID := "924351368897106061"

	HOOKURL := "https://discord.com/api/webhooks/1057740580861509702/n9pV73DsR6yL3fr5qIbske_gcHsfYTiBpkgEKCVbMmJrE82dG44GkB-1ZNpw-pwW0U3t"

	HOOKID := "1055577948297625600"
	HOOKTOKEN := "ci8JcMy8tOfBjI6n3lvz_czL2qepFs4222nOpjQZcr3GlfipNJ_xowy3N-Z1pYBZ6yl7"

	GGAPITOKEN := "AIzaSyDNiNV1ujFuKNCir17Oev7bhDQZgF7givw"
	BINGTOKEN := "Ahz6TU3QuNiBo8gYTTEzJpbTsZpg4TbsOkMcmsTpg6QpopqrS99Qp3BCImIzLFR7"
	NASATOKEN := "lgVZOMJlINVlio85EHjb3Z14BkbJlcwgcmKprZQS"
	BOTID := "1055553353754628197"

	token := TOKEN
	serverID := SERVERID
	ownerID := OWNERID
	dbKey := DBKEY
	hookToken := HOOKID
	hookUrl := HOOKURL
	hookId := HOOKTOKEN
	ggApi := GGAPITOKEN
	bingApi := BINGTOKEN

	flag.StringVar(&Token, "t", token, "Bot's token")
	flag.StringVar(&ServerID, "s", serverID, "Server's id")
	flag.StringVar(&OwnerID, "o", ownerID, "Owner's id")
	flag.StringVar(&DBKey, "d", dbKey, "")
	flag.StringVar(&HookID, "ht", hookToken, "")
	flag.StringVar(&HookToken, "hi", hookId, "")
	flag.StringVar(&HookURL, "hu", hookUrl, "")
	flag.StringVar(&GGAPIKey, "ga", ggApi, "")
	flag.StringVar(&BINGKey, "ba", bingApi, "")
	flag.StringVar(&WeatherKey, "wa", "e5bafc3618ad164189d42eece89d6bc2", "")
	flag.StringVar(&NasaKey, "na", NASATOKEN, "")
	flag.StringVar(&DBLKey, "dbl", DBLTOKEN, "")
	flag.StringVar(&BotID, "bi", BOTID, "")

	flag.Parse()

	if Token == "" || ServerID == "" || OwnerID == "" || DBKey == "" {
		log.Fatal(errors.New("missing credentials"))
	}
}
