package favorites_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

<<<<<<< HEAD
	"api-tests-template/internal/helpers/auth"
	adManager "api-tests-template/internal/managers/advertisement"
	adModels "api-tests-template/internal/managers/advertisement/models"
	favManager "api-tests-template/internal/managers/favorites"
=======
	"api-tests-template/internal/helpers/api"
	adManager "api-tests-template/internal/manager/advertisement"
	adModels "api-tests-template/internal/managers/advertisement/models"
	favManager "api-tests-template/internal/manager/favorites"
>>>>>>> b13fc60 (final)
)

func TestFavorites_Positive(t *testing.T) {
	t.Log("Шаг 1: Авторизация")
<<<<<<< HEAD
	token := auth.GetAuthToken(t)
	uniqueTitle := "Тест для избранного " + time.Now().Format("20060102150405")
	price := 1000.0
	qty := 1
	req := adModels.CreateAdvertisementRequest{
		Title:       strPtr(uniqueTitle),
		Description: strPtr("Тестовое описание"),
		Price:       &price,
		Quantity:    &qty,
	}
	adID := adManager.CreateAndAssert(t, token, req)
	require.NotEmpty(t, adID, "ID не должен быть пустым")
	t.Log("Шаг 2: Добавление в избранное")
	favManager.AddAndAssert(t, token, adID)
	t.Log("Шаг 3: Проверка наличия")
	favManager.GetAndAssert(t, token, adID)
	t.Log("Шаг 4: Удаление из избранного")
	favManager.DeleteAndAssert(t, token, adID)
	t.Log("✅ Тест пройден")
}
=======
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
>>>>>>> b13fc60 (final)
