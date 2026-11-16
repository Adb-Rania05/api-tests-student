package favorites

import (
	"fmt"
	"net/http"
	"testing"

	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
)

// HttpPostFavorite добавляет объявление в избранное
func HttpPostFavorite(t *testing.T, token string, adID string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Post(path.FavoritesPath).
		MultipartFormData("ad_id", adID).
		Expect(t).
		End().Response
}

// HttpGetFavorites получает список избранного
func HttpGetFavorites(t *testing.T, token string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Get(path.FavoritesPath).
		Expect(t).
		End().Response
}

// HttpDeleteFavorite удаляет объявление из избранного
func HttpDeleteFavorite(t *testing.T, token string, adID string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Delete(fmt.Sprintf("%s/%s", path.FavoritesPath, adID)).
		Expect(t).
		End().Response
}