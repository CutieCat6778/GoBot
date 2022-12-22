package class

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Token          string
	ServerID       string
	OwnerID        string
	HookID         string
	HookToken      string
	HookURL        string
	DBKey          string
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
		return
	}

	token := os.Getenv("TOKEN")
	serverID := os.Getenv("SERVERID")
	ownerID := os.Getenv("OWNERID")
	dbKey := os.Getenv("DBKEY")
	hookToken := os.Getenv("HOOKID")
	hookUrl := os.Getenv("HOOKURL")
	hookId := os.Getenv("HOOKTOKEN")

	flag.StringVar(&Token, "t", token, "Bot's token")
	flag.StringVar(&ServerID, "s", serverID, "Server's id")
	flag.StringVar(&OwnerID, "o", ownerID, "Owner's id")
	flag.StringVar(&DBKey, "d", dbKey, "")
	flag.StringVar(&HookID, "ht", hookToken, "")
	flag.StringVar(&HookToken, "hi", hookId, "")
	flag.StringVar(&HookURL, "hu", hookUrl, "")

	flag.Parse()

	if Token == "" || ServerID == "" || OwnerID == "" || DBKey == "" {
		log.Fatal("Missing creadentials!")
		return
	}
}
