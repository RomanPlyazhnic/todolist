package functional

import (
	"encoding/json"
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"

	"github.com/RomanPlyazhnic/todolist/tests/suite"
)

type RegisterRequest struct {
	Username string `faker:"username"`
	Password string `faker:"password"`
}

type LoginRequest struct {
	Username string `faker:"username"`
	Password string `faker:"password"`
}

func TestLogin_RegisterLoginHappyPath(t *testing.T) {
	s := suite.New(t)

	// Register user
	r := RegisterRequest{}
	err := faker.FakeData(&r)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	jsonStr, err := json.Marshal(r)
	reqBody := strings.NewReader(string(jsonStr))

	url := fmt.Sprintf("http://%s:%d/Register", s.App.Config.Domain, s.App.Config.Port)
	resp, err := http.Post(url, "application/json", reqBody)
	closeResp(t, resp)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	assert.Equal(t, 200, resp.StatusCode)

	// Login registered user
	l := LoginRequest{
		Username: r.Username,
		Password: r.Password,
	}
	jsonStr, err = json.Marshal(l)
	reqBody = strings.NewReader(string(jsonStr))

	url = fmt.Sprintf("http://%s:%d/Login", s.App.Config.Domain, s.App.Config.Port)
	resp, err = http.Post(url, "application/json", reqBody)
	closeResp(t, resp)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	assert.Equal(t, 200, resp.StatusCode)

	// Access service by logined user
	client := &http.Client{}
	url = fmt.Sprintf("http://%s:%d", s.App.Config.Domain, s.App.Config.Port)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	cookie := resp.Cookies()[0]
	req.AddCookie(cookie)

	resp, err = client.Do(req)
	defer closeResp(t, resp)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	assert.Equal(t, 200, resp.StatusCode)
}

func closeResp(t *testing.T, resp *http.Response) {
	if resp == nil {
		return
	}

	err := resp.Body.Close()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
