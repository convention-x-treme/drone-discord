package discord

import (
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func setup() {
	// Load env-file if it exists first
	if filename, found := os.LookupEnv("PLUGIN_ENV_FILE"); found {
		godotenv.Load(filename)
	}

	if _, err := os.Stat("/run/drone/env"); err == nil {
		godotenv.Overload("/run/drone/env")
	}
}

func TestMissingConfig(t *testing.T) {
	setup()
	var plugin Plugin

	err := plugin.Exec()

	assert.NotNil(t, err)
}

func TestTemplate(t *testing.T) {
	setup()
	plugin := Plugin{
		Repo: Repo{
			Name:      "go-hello",
			Namespace: "convention-x-treme",
		},
		Commit: Commit{
			Author:  "convention-x-treme",
			Branch:  "master",
			Message: "update by drone discord plugin. \r\n update by drone discord plugin.",
			Avatar:  "https://avatars0.githubusercontent.com/u/21979?v=3&s=100",
		},
		Build: Build{
			Number: 101,
			Status: "success",
			Link:   "https://github.com/convention-x-treme/go-hello",
			Event:  "tag",
		},
		Config: Config{
			WebhookID:    os.Getenv("WEBHOOK_ID"),
			WebhookToken: os.Getenv("WEBHOOK_TOKEN"),
			Message:      []string{"test one message from drone testing", "test two message from drone testing"},
			File:         []string{"../../images/discord-logo.png"},
			Drone:        true,
		},
		Payload: Payload{
			Username: "drone",
			TTS:      false,
			Wait:     false,
		},
	}

	err := plugin.Exec()
	assert.Nil(t, err)

	plugin.Clear()
	plugin.Config.Message = []string{"We are Convention-X-Treme"}
	plugin.Config.File = nil
	plugin.Payload.TTS = false
	plugin.Payload.Wait = true
	err = plugin.Exec()
	assert.Nil(t, err)

	// send success embed message
	plugin.Config.Message = []string{}
	plugin.Config.File = nil
	plugin.Payload.TTS = false
	plugin.Payload.Wait = false
	plugin.Clear()
	err = plugin.Exec()
	assert.Nil(t, err)

	// send success embed message
	plugin.Build.Status = "failure"
	plugin.Commit.Message = "send failure embed message"
	plugin.Clear()
	err = plugin.Exec()
	assert.Nil(t, err)
	time.Sleep(1 * time.Second)

	// send default embed message
	plugin.Build.Status = "test"
	plugin.Commit.Message = "send default embed message"
	plugin.Clear()
	err = plugin.Exec()
	assert.Nil(t, err)

	// change color for embed message
	plugin.Config.Color = "#4842f4"
	plugin.Commit.Message = "Change embed color to #4842f4"
	plugin.Clear()
	err = plugin.Exec()
	assert.Nil(t, err)
}

func TestDefaultTemplate(t *testing.T) {
	plugin := Plugin{
		Config: Config{
			WebhookID:    os.Getenv("WEBHOOK_ID"),
			WebhookToken: os.Getenv("WEBHOOK_TOKEN"),
			Message:      []string{"default message 1", "default message 2"},
			Color:        "#48f442",
		},

		Payload: Payload{
			Username: "drone-ci",
			TTS:      false,
			Wait:     false,
		},
	}

	time.Sleep(1 * time.Second)
	plugin.Clear()
	err := plugin.Exec()
	assert.Nil(t, err)

	plugin.Config.Color = "#f4be41"
	time.Sleep(1 * time.Second)
	plugin.Clear()
	err = plugin.Exec()
	assert.Nil(t, err)
}
