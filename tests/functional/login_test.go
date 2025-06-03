package functional

import (
	"encoding/json"
	"fmt"
	"github.com/RomanPlyazhnic/todolist/internal/app/contracts"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"

	"github.com/RomanPlyazhnic/todolist/tests/suite"
)

func TestLogin_LoginHappyPath(t *testing.T) {
	s := suite.New(t)

	url := fmt.Sprintf("http://%s:%d/Register", s.App.Config.Domain, s.App.Config.Port)
	s.App.Logger.Info(url)

	r := contracts.RegisterRequest{
		Username: "testuser",
		Password: "testpassword",
	}
	jsonStr, err := json.Marshal(r)
	body := strings.NewReader(string(jsonStr))
	resp, err := http.Post(url, "application/json", body)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	assert.Equal(t, 200, resp.StatusCode)
}
