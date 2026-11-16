package favorites

import (
	"io"
	"testing"

	"github.com/stretchr/testify/require"

	favClient "api-tests-template/internal/client/http/favorites"
)

func AddAndAssert(t *testing.T, token string, adID string) {
	t.Helper()
	resp := favClient.HttpPostFavorite(t, token, adID)
	require.Contains(t, []int{200, 201, 204}, resp.StatusCode, "Ожидаем успешное добавление")
}

func GetAndAssert(t *testing.T, token string, adID string) {
	t.Helper()
	resp := favClient.HttpGetFavorites(t, token)
	require.Equal(t, 200, resp.StatusCode, "Ожидаем 200 OK")
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "Не удалось прочитать ответ")
	defer resp.Body.Close()
	require.Contains(t, string(body), adID, "Объявление должно быть в избранном")
}

func DeleteAndAssert(t *testing.T, token string, adID string) {
	t.Helper()
	resp := favClient.HttpDeleteFavorite(t, token, adID)
	require.Contains(t, []int{200, 204}, resp.StatusCode, "Ожидаем успешное удаление")
}
