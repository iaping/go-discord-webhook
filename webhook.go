package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Webhook struct {
	Url string
}

func New(url string) *Webhook {
	return &Webhook{
		Url: url,
	}
}

func (hook *Webhook) Say(content string) error {
	return hook.Send(Message{Content: content})
}

func (hook *Webhook) Send(msg Message) error {
	body := &bytes.Buffer{}
	if err := json.NewEncoder(body).Encode(msg); err != nil {
		return err
	}

	resp, err := http.Post(hook.Url, "application/json", body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("error http code (%d)", resp.StatusCode)
	}

	return nil
}
