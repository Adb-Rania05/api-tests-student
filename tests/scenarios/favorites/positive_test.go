package favorites_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/helpers/api"
	adManager "api-tests-template/internal/manager/advertisement"
	adModels "api-tests-template/internal/managers/advertisement/models"
	favManager "api-tests-template/internal/manager/favorites"
)

func TestFavorites_Positive(t *testing.T) {
	t.Log("Шаг 1: Авторизация")
	token := api.GetAuthToken(t)

	t.Log("Шаг 2: Создаём объявление")
	req := adModels.CreateAdvertisementRequest{
		Title:       strPtr("Тест избранного " + time.Now().Format("150405")),
		Description: strPtr("Описание"),
		Price:       floatPtr(1000.0),
		Quantity:    intPtr(1),
	}
	adID := adManager.CreateAndAssert(t, token, req)

	t.Log("Шаг 3: Добавляем в избранное")
	favManager.AddAndAssert(t, token, adID)

	t.Log("Шаг 4: Проверяем наличие в избранном")
	favManager.GetAndAssert(t, token, adID)

	t.Log("Шаг 5: Удаляем из избранного")
	favManager.DeleteAndAssert(t, token, adID)

	t.Log("✅ Тест пройден")
}