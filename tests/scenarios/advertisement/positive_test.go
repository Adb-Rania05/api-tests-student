package advertisement_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/helpers/auth"
	adManager "api-tests-template/internal/managers/advertisement"
	adModels "api-tests-template/internal/managers/advertisement/models"
)

// TestCreateAdvertisement_Positive — позитивный сценарий создания объявления.
func TestCreateAdvertisement_Positive(t *testing.T) {
	t.Log("Шаг 1: Авторизация")
	token := auth.GetAuthToken(t)

	uniqueTitle := "Тестовый товар " + time.Now().Format("20060102150405")
	price := 1000.0
	qty := 1

	req := adModels.CreateAdvertisementRequest{
		Title:       strPtr(uniqueTitle),
		Description: strPtr("Тестовое описание"),
		Price:       &price,
		Quantity:    &qty,
	}

	t.Log("Шаг 2: Создание объявления")
	adID := adManager.CreateAndAssert(t, token, req)
	require.NotEmpty(t, adID, "ID не должен быть пустым")

	t.Log("Шаг 3: Проверка через GET")
	adManager.GetAndAssert(t, adID, uniqueTitle)

	t.Log("✅ Тест пройден")
}