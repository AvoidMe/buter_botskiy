package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

const (
	Url = "https://api.telegram.org/bot%s/%s"
)

type Config struct {
	TgToken string `yaml:"tg_token"`
	ChatID  int64  `yaml:"chat_id"`
	PhotoID string `yaml:"photo_id"`
}

func ParseConfig() (*Config, error) {
	content, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func getUpdates(cfg *Config, offset int64) (*Update, error) {
	data, err := json.Marshal(map[string]any{
		"offset":          offset,
		"timeout":         30,
		"allowed_updates": []string{"message_reaction"},
	})
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(Url, cfg.TgToken, "getUpdates")
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := Update{}
	err = json.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(content))
	return &result, nil
}

func postButer(cfg *Config) (int64, error) {
	data, err := json.Marshal(map[string]any{
		"chat_id": cfg.ChatID,
		"photo":   cfg.PhotoID,
	})
	if err != nil {
		return 0, err
	}

	url := fmt.Sprintf(Url, cfg.TgToken, "sendPhoto")
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(data))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	result := Message{}
	err = json.Unmarshal(content, &result)
	if err != nil {
		return 0, err
	}
	fmt.Println(string(content))
	return result.Result.MessageID, nil
}

func bot(cfg *Config) {
	last_buter := int64(-1)
	last_update := int64(0)
	var err error
	count := 0
	for {
		// TODO: if last_buter == -1 -> post new
		if last_buter == -1 {
			last_buter, err = postButer(cfg)
			if err != nil {
				last_buter = -1
				log.Printf("Unable to post new buter: %v", err)
				time.Sleep(10 * time.Second)
				continue
			}
		}
		updates, err := getUpdates(cfg, last_update)
		if err != nil {
			log.Printf("Unable to get updates: %v, sleep for 10", err)
			time.Sleep(10 * time.Second)
			continue
		}
		for _, res := range updates.Result {
			last_update = max(last_update, res.UpdateID+1)
			if res.MessageReaction.Chat.ID != cfg.ChatID {
				continue
			}
			if res.MessageReaction.MessageID != int64(last_buter) {
				continue
			}
			for _, react := range res.MessageReaction.NewReaction {
				if react.Emoji == "👍" {
					count++
				}
			}
		}
		if count >= 2 {
			last_buter = -1
			count = 0
		}
	}
}

func main() {
	cfg, err := ParseConfig()
	if err != nil {
		log.Fatalf("Unable to parse config: %v", err)
	}
	bot(cfg)
}
