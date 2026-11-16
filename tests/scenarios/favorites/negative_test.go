package favorites_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/helpers/api"
	favClient "api-tests-template/internal/client/http/favorites"
)

func TestFavorites_Negative_NoAuth(t *testing.T) {
	t.Log("Шаг 1: Попытка добавить в избранное без токена")
	resp := favClient.HttpPostFavorite(t, "", "fake-id")
	require.Equal(t, http.StatusUnauthorized, resp.StatusCode, "Ожидаем 401 без токена")
	t.Log("✅ Тест пройден")
}

func TestFavorites_Negative_InvalidID(t *testing.T) {
	t.Log("Шаг 1: Авторизация")
	token := api.GetAuthToken(t)

	t.Log("Шаг 2: Попытка добавить несуществующее объявление")
	resp := favClient.HttpPostFavorite(t, token, "non-existent-id")
	require.GreaterOrEqual(t, resp.StatusCode, 400, "Ожидаем ошибку 4xx")
	t.Log("✅ Тест пройден")
}