package main

type Message struct {
	Ok     bool `json:"ok"`
	Result struct {
		MessageID int64 `json:"message_id"`
		From      struct {
			ID        int64  `json:"id"`
			IsBot     bool   `json:"is_bot"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
		} `json:"from"`
		Chat struct {
			ID        int64  `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date  int `json:"date"`
		Photo []struct {
			FileID       string `json:"file_id"`
			FileUniqueID string `json:"file_unique_id"`
			FileSize     int64  `json:"file_size"`
			Width        int64  `json:"width"`
			Height       int64  `json:"height"`
		} `json:"photo"`
	} `json:"result"`
}

type Update struct {
	Ok     bool `json:"ok"`
	Result []struct {
		UpdateID        int64 `json:"update_id"`
		MessageReaction struct {
			Chat struct {
				ID        int64  `json:"id"`
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
				Username  string `json:"username"`
				Type      string `json:"type"`
			} `json:"chat"`
			MessageID int64 `json:"message_id"`
			User      struct {
				ID           int64  `json:"id"`
				IsBot        bool   `json:"is_bot"`
				FirstName    string `json:"first_name"`
				LastName     string `json:"last_name"`
				Username     string `json:"username"`
				LanguageCode string `json:"language_code"`
			} `json:"user"`
			Date        int64 `json:"date"`
			OldReaction []any `json:"old_reaction"`
			NewReaction []struct {
				Type  string `json:"type"`
				Emoji string `json:"emoji"`
			} `json:"new_reaction"`
		} `json:"message_reaction"`
	} `json:"result"`
}
