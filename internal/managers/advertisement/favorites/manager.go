package favorites

import (
	"io"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"

	"api-tests-template/internal/client/http/favorites"
)

// AddAndAssert добавляет объявление в избранное и проверяет статус.
func AddAndAssert(t *testing.T, token string, adID string) {
	t.Helper()
	resp := favorites.HttpPostFavorite(t, token, adID)
	require.Contains(t, []int{200, 201, 204}, resp.StatusCode, "Ожидаем успешное добавление в избранное")
}

// GetAndAssert проверяет, что объявление есть в избранном.
func GetAndAssert(t *testing.T, token string, adID string) {
	t.Helper()
	resp := favorites.HttpGetFavorites(t, token)
	require.Equal(t, 200, resp.StatusCode, "Ожидаем 200 OK при получении избранного")

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "Не удалось прочитать ответ")
	defer resp.Body.Close()

	require.True(t, gjson.GetBytes(body, "data").Exists() || gjson.GetBytes(body, "items").Exists(),
		"Ответ должен содержать список избранного")
}

// DeleteAndAssert удаляет объявление из избранного.
func DeleteAndAssert(t *testing.T, token string, adID string) {
	t.Helper()
	resp := favorites.HttpDeleteFavorite(t, token, adID)
	require.Contains(t, []int{200, 204}, resp.StatusCode, "Ожидаем успешное удаление из избранного")
}
