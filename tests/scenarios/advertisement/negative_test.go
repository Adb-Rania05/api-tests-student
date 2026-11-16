package advertisement_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	adClient "api-tests-template/internal/client/http/advertisement"
	"api-tests-template/internal/helpers/auth"
	adModels "api-tests-template/internal/managers/advertisement/models"
)

// TestCreateAdvertisement_Negative_NoAuth — проверка 401 без токена.
func TestCreateAdvertisement_Negative_NoAuth(t *testing.T) {
	t.Log("Шаг 1: Попытка создать объявление без токена")
	price := 100.0
	qty := 1
	req := adModels.CreateAdvertisementRequest{
		Title:    strPtr("Объявление без токена"),
		Price:    &price,
		Quantity: &qty,
	}
	resp := adClient.HttpPostAdvertisement(t, "", req)
	require.Equal(t, http.StatusUnauthorized, resp.StatusCode, "Без токена ожидаем 401")
	t.Log("✅ Тест пройден")
}

// TestCreateAdvertisement_Negative_InvalidPrice — проверка отклонения отрицательной цены.
func TestCreateAdvertisement_Negative_InvalidPrice(t *testing.T) {
	t.Log("Шаг 1: Авторизация")
	token := auth.GetAuthToken(t)

	t.Log("Шаг 2: Создание объявления с отрицательной ценой")
	badPrice := -500.0
	qty := 1
	req := adModels.CreateAdvertisementRequest{
		Title:    strPtr("Объявление с плохой ценой"),
		Price:    &badPrice,
		Quantity: &qty,
	}
	resp := adClient.HttpPostAdvertisement(t, token, req)
	require.GreaterOrEqual(t, resp.StatusCode, 400, "Ожидаем ошибку 4xx")
	require.Less(t, resp.StatusCode, 500, "Не ожидаем 5xx")
	t.Log("✅ Тест пройден")
}