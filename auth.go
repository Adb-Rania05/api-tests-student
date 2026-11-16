package auth

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func GetAuthToken(t *testing.T) string {
	t.Helper()
	_ = godotenv.Load(".env", ".env.override")

	login := os.Getenv("TEST_LOGIN")
	password := os.Getenv("TEST_PASSWORD")
	require.NotEmpty(t, login, "TEST_LOGIN не задан")
	require.NotEmpty(t, password, "TEST_PASSWORD не задан")

	baseURL := os.Getenv("API_BASE_URL")
	if baseURL == "" {
		baseURL = "https://testboard.avito.com"
	}

	payload, _ := json.Marshal(map[string]string{"login": login, "password": password})
	resp, err := http.Post(baseURL+"/auth/login", "application/json", bytes.NewBuffer(payload))
	require.NoError(t, err, "Ошибка запроса авторизации")
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	require.Equal(t, 200, resp.StatusCode, "Авторизация вернула не 200: "+string(body))

	token := gjson.GetBytes(body, "token").String()
	require.NotEmpty(t, token, "Токен не найден: "+string(body))
	return token
}
