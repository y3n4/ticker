package storage

import (
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/mattn/go-mastodon"
)

type Ticker struct {
	ID           int       `storm:"id,increment"`
	CreationDate time.Time `storm:"index"`
	Domain       string    `storm:"unique"`
	Title        string
	Description  string
	Active       bool
	Information  Information
	Twitter      Twitter
	Telegram     Telegram
	Mastodon     Mastodon
	Location     Location
}

func NewTicker() Ticker {
	return Ticker{
		CreationDate: time.Now(),
	}
}

func (t *Ticker) Reset() {
	t.Active = false
	t.Description = ""
	t.Information = Information{}
	t.Location = Location{}

	t.Twitter.Reset()
	t.Telegram.Reset()
	t.Mastodon.Reset()
}

type Information struct {
	Author   string
	URL      string
	Email    string
	Twitter  string
	Facebook string
	Telegram string
}

type Twitter struct {
	Active bool   `json:"active"`
	Token  string `json:"token"`
	Secret string `json:"secret"`
	User   twitter.User
}

func (tw *Twitter) Reset() {
	tw.Active = false
	tw.Token = ""
	tw.Secret = ""
	tw.User = twitter.User{}
}

func (tw *Twitter) Connected() bool {
	return tw.Token != "" && tw.Secret != ""
}

type Telegram struct {
	Active      bool   `json:"active"`
	ChannelName string `json:"channel_name"`
}

func (tg *Telegram) Reset() {
	tg.Active = false
	tg.ChannelName = ""
}

func (tg *Telegram) Connected() bool {
	return tg.ChannelName != ""
}

type Mastodon struct {
	Active      bool   `json:"active"`
	Server      string `json:"server"`
	Token       string `json:"token"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
	User        mastodon.Account
}

func (m *Mastodon) Connected() bool {
	return m.Token != "" && m.Secret != "" && m.AccessToken != ""
}

func (m *Mastodon) Reset() {
	m.Active = false
	m.Server = ""
	m.Token = ""
	m.Secret = ""
	m.AccessToken = ""
	m.User = mastodon.Account{}
}

type Location struct {
	Lat float64
	Lon float64
}
