package advertisement

import (
	"io"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"

	"api-tests-template/internal/client/http/advertisement"
	adModels "api-tests-template/internal/managers/advertisement/models"
)

// CreateAndAssert создаёт объявление и возвращает его ID.
func CreateAndAssert(t *testing.T, token string, req adModels.CreateAdvertisementRequest) string {
	t.Helper()
	resp := advertisement.HttpPostAdvertisement(t, token, req)
	require.Contains(t, []int{200, 201}, resp.StatusCode, "Ожидаем успешный статус при создании")

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "Не удалось прочитать ответ")
	defer resp.Body.Close()

	adID := gjson.GetBytes(body, "id").String()
	require.NotEmpty(t, adID, "ID объявления не найден в ответе")

	return adID
}

// GetAndAssert получает объявление и сверяет title.
func GetAndAssert(t *testing.T, id string, expectedTitle string) {
	t.Helper()
	resp := advertisement.HttpGetAdvertisement(t, id)
	require.Equal(t, 200, resp.StatusCode, "Ожидаем 200 OK")

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "Не удалось прочитать ответ")
	defer resp.Body.Close()

	require.Equal(t, expectedTitle, gjson.GetBytes(body, "title").String(), "Title не совпадает")
}
