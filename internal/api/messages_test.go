package api_test

import (
	"testing"
	"encoding/json"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"

	"git.codecoop.org/systemli/ticker/internal/api"
	"git.codecoop.org/systemli/ticker/internal/model"
	"git.codecoop.org/systemli/ticker/internal/storage"
	"strings"
)

func TestGetMessages(t *testing.T) {
	r := setup()

	ticker := new(model.Ticker)
	ticker.ID = 1

	storage.DB.Save(ticker)

	r.GET("/v1/admin/tickers/1/messages").
		Run(api.API(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
		assert.Equal(t, 200, r.Code)
		assert.Equal(t, `{"data":{"messages":[]},"status":"success","error":null}`, strings.TrimSpace(r.Body.String()))
	})
}

func TestGetMessage(t *testing.T) {
	r := setup()

	ticker := new(model.Ticker)
	ticker.ID = 1

	storage.DB.Save(ticker)

	r.GET("/v1/admin/tickers/1/messages/1").
		Run(api.API(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
		assert.Equal(t, 404, r.Code)
		assert.Equal(t, `{"data":{},"status":"error","error":{"code":1001,"message":"not found"}}`, strings.TrimSpace(r.Body.String()))
	})
}

func TestPostMessage(t *testing.T) {
	r := setup()

	ticker := new(model.Ticker)
	ticker.ID = 1

	storage.DB.Save(ticker)

	body := `{
		"text": "message"
	}`

	r.POST("/v1/admin/tickers/1/messages").
		SetBody(body).
		Run(api.API(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
		assert.Equal(t, 200, r.Code)

		type jsonResp struct {
			Data   map[string]model.Message `json:"data"`
			Status string                   `json:"status"`
			Error  interface{}              `json:"error"`
		}

		var jres jsonResp

		err := json.Unmarshal(r.Body.Bytes(), &jres)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, model.ResponseSuccess, jres.Status)
		assert.Equal(t, nil, jres.Error)
		assert.Equal(t, 1, len(jres.Data))

		message := jres.Data["message"]

		assert.Equal(t, "message", message.Text)
		assert.Equal(t, 1, message.Ticker)
	})
}

func TestDeleteMessage(t *testing.T) {
	r := setup()

	ticker := new(model.Ticker)
	ticker.ID = 1

	storage.DB.Save(ticker)

	message := model.NewMessage()
	message.Text = "Text"
	message.Ticker = 1

	storage.DB.Save(&message)

	r.DELETE("/v1/admin/tickers/1/messages/2").
		Run(api.API(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
		assert.Equal(t, 404, r.Code)
	})

	r.DELETE("/v1/admin/tickers/1/messages/1").
		Run(api.API(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
		assert.Equal(t, 200, r.Code)

		type jsonResp struct {
			Data   map[string]model.Message `json:"data"`
			Status string                   `json:"status"`
			Error  interface{}              `json:"error"`
		}

		var jres jsonResp

		err := json.Unmarshal(r.Body.Bytes(), &jres)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, model.ResponseSuccess, jres.Status)
		assert.Nil(t, jres.Data)
		assert.Nil(t, jres.Error)
	})
}
