package slack

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/VirtualWarehouse/jb/internal/client"
	"github.com/VirtualWarehouse/jb/internal/config"
	"github.com/slack-go/slack"
	"github.com/spf13/viper"
)

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	configHome := filepath.Join(home, ".config/jb")
	configName := "config"
	configType := "yml"
	configPath := filepath.Join(configHome, configName+"."+configType)

	viper.AddConfigPath(configHome)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	viper.AutomaticEnv()

	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		if _, err := os.Create(configPath); err != nil {
			panic(err)
		}
	}

	err = config.LoadConfig()
	if err != nil {
		panic(err)
	}

	checkConfig()
	SlackInit()
}

func SlackInit() {
	sendMessage()
	http.HandleFunc("/", actionHandler)
	http.ListenAndServe(":3000", nil)
}

var api *slack.Client
var ch *slack.Channel
var err error
var t string

func sendMessage() {
	api = slack.New(viper.GetString("apptoken"))
	params := &slack.OpenConversationParameters{
		Users: []string{viper.GetString("userid")},
	}
	ch, _, _, err = api.OpenConversation(params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	attachment := slack.Attachment{
		Fallback:   "We don't currently support your client",
		CallbackID: "accept_or_reject",
		Color:      "#FFFFFF",
		Actions: []slack.AttachmentAction{
			{
				Name:  "touch",
				Text:  "Touch",
				Type:  "button",
				Value: "touch",
				Style: "primary",
			},
		},
	}

	message := slack.MsgOptionAttachments(attachment)
	_, t, _ = api.PostMessage(ch.ID, slack.MsgOptionText("Hi", false))
	chID, timestamp, err := api.PostMessage(ch.ID, slack.MsgOptionText("", false), message)
	fmt.Printf("Message with buttons sucessfully sent to channel %s at %s\n", chID, timestamp)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}

func actionHandler(w http.ResponseWriter, r *http.Request) {
	var payload slack.InteractionCallback
	err := json.Unmarshal([]byte(r.FormValue("payload")), &payload)
	if err != nil {
		fmt.Printf("Could not parse action response JSON: %v", err)
	}
	log.Printf("Message button pressed by user %s with value %s", payload.User.Name, payload.ActionCallback.AttachmentActions[0].Value)

	err = shrug([]string{"touch"})
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	update := slack.MsgOptionUpdate(t)
	msg := fmt.Sprintf("Last Updated: %s", time.Now().Format("2006-01-02 15:04:05"))
	_, _, _ = api.PostMessage(ch.ID, slack.MsgOptionText(msg, false), update)
}

func shrug(args []string) error {
	c := client.New(config.Cfg)
	text := strings.Join(args, " ")

	resp, err := c.PostChatCommand(config.Cfg.TouchChannel, text, "shrug")
	if err != nil {
		return err
	}
	if !resp.OK {
		return errors.New(resp.Error)
	}

	fmt.Printf("Successfully touch: \"%s %s\"\n", text, `¯\_(ツ)_/¯`)

	return nil
}
