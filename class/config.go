package class

import (
	"flag"
	"log"
)

var (
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
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
)

type CommandData struct {
	Permissions    int64
	Ratelimit      int64
	BotPerms       int64
	SubCommandData map[string]CommandData
}

func init() {
	TOKEN := "MTA1NTU1MzM1Mzc1NDYyODE5Nw.GbYWaN.lIF0mq9BYSDR4rSUSD8iwd3kR5YyC_6hHrN3qk"
	DBKEY := "mongodb+srv://admin:!Txzje2006@cluster0.axkhjad.mongodb.net/?retryWrites=true&w=majority"
	SERVERID := "1054737473802096671"
	OWNERID := "924351368897106061"

	HOOKURL := "https://discord.com/api/webhooks/1055577948297625600/ci8JcMy8tOfBjI6n3lvz_czL2qepFs4222nOpjQZcr3GlfipNJ_xowy3N-Z1pYBZ6yl7"

	HOOKID := "1055577948297625600"
	HOOKTOKEN := "ci8JcMy8tOfBjI6n3lvz_czL2qepFs4222nOpjQZcr3GlfipNJ_xowy3N-Z1pYBZ6yl7"

	GGAPITOKEN := "AIzaSyDNiNV1ujFuKNCir17Oev7bhDQZgF7givw"
	BINGTOKEN := "Ahz6TU3QuNiBo8gYTTEzJpbTsZpg4TbsOkMcmsTpg6QpopqrS99Qp3BCImIzLFR7"

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

	flag.Parse()

	if Token == "" || ServerID == "" || OwnerID == "" || DBKey == "" {
		log.Fatal("Missing creadentials!")
		return
	}
}
