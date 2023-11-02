package problem

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDetails_MarshalJSON_Default(t *testing.T) {
	now := time.Now()

	pb := New(http.StatusForbidden, "Your current balance is 30, but that costs 50.").
		Error(errors.New("test error")). //nolint:goerr113
		Put(fieldTimestamp, now).
		Request(&http.Request{RequestURI: "/transfer"}) //nolint:exhaustruct

	data, err := json.Marshal(pb)
	require.NoError(t, err)

	assert.Equal(
		t,
		`{`+
			`"detail":"Your current balance is 30, but that costs 50.",`+
			`"error":"test error",`+
			`"path":"/transfer",`+
			`"status":403,`+
			`"timestamp":"`+now.Format(time.RFC3339Nano)+`",`+
			`"title":"Forbidden",`+
			`"type":"about:blank"`+
			`}`,
		string(data),
	)
}

func TestDetails_MarshalJSON_Custom(t *testing.T) {
	now := time.Now()

	instanceURL, err := url.Parse("/account/12345/msgs/abc")
	require.NoError(t, err)

	typeURL, err := url.Parse("https://example.com/probs/out-of-credit")
	require.NoError(t, err)

	pb := New(http.StatusForbidden, "Your current balance is 30, but that costs 50.").
		Put(fieldTimestamp, now).
		Put("balance", 30).
		Put("accounts", []string{"/account/12345", "/account/67890"})

	pb.Instance = instanceURL
	pb.Title = "You do not have enough credit."
	pb.Type = typeURL

	data, err := json.Marshal(pb)
	require.NoError(t, err)

	assert.Equal(
		t,
		`{`+
			`"accounts":["/account/12345","/account/67890"],`+
			`"balance":30,`+
			`"detail":"Your current balance is 30, but that costs 50.",`+
			`"instance":"/account/12345/msgs/abc",`+
			`"status":403,`+
			`"timestamp":"`+now.Format(time.RFC3339Nano)+`",`+
			`"title":"You do not have enough credit.",`+
			`"type":"https://example.com/probs/out-of-credit"`+
			`}`,
		string(data),
	)
}

func TestDetails_UnmarshalJSON(t *testing.T) {
	pb := new(Details)

	err := pb.UnmarshalJSON([]byte(`{
    	"accounts": ["/account/12345", "/account/67890"],
    	"balance": 30,
    	"detail": "Your current balance is 30, but that costs 50.",
    	"instance": "/account/12345/msgs/abc",
		"status": 403,
    	"title": "You do not have enough credit.",
    	"type": "https://example.com/probs/out-of-credit"
    }`))
	require.NoError(t, err)

	accounts, found := pb.Get("accounts")
	assert.Equal(t, []any{"/account/12345", "/account/67890"}, accounts)
	assert.True(t, found)

	balance, found := pb.Get("balance")
	assert.Equal(t, float64(30), balance)
	assert.True(t, found)

	assert.Equal(t, "Your current balance is 30, but that costs 50.", pb.Detail)

	assert.NotNil(t, pb.Instance)
	assert.Equal(t, "/account/12345/msgs/abc", pb.Instance.String())

	assert.Equal(t, 403, pb.Status)
	assert.Equal(t, "You do not have enough credit.", pb.Title)

	assert.NotNil(t, pb.Type)
	assert.Equal(t, "https://example.com/probs/out-of-credit", pb.Type.String())
}
