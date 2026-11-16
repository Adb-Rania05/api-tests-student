package advertisement

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
	adModels "api-tests-template/internal/managers/advertisement/models"
)

func HttpPostAdvertisement(t *testing.T, token string, request adModels.CreateAdvertisementRequest) *http.Response {
	t.Helper()
	builder := apiRunner.GetRunner().Auth(token).Create().Post(path.AdvertisementPath)

	if request.Title != nil {
		builder = builder.MultipartFormData("title", *request.Title)
	}
	if request.Description != nil {
		builder = builder.MultipartFormData("description", *request.Description)
	}
	if request.Price != nil {
		builder = builder.MultipartFormData("price", strconv.FormatFloat(*request.Price, 'f', -1, 64))
	}
	if request.Quantity != nil {
		builder = builder.MultipartFormData("quantity", strconv.Itoa(*request.Quantity))
	}
	for _, photoPath := range request.Photos {
		builder = builder.MultipartFile("photos", photoPath)
	}

	return builder.Expect(t).End().Response
}

func HttpGetAdvertisement(t *testing.T, id string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Create().Get(path.AdvertisementPath).
		Query("id", id).
		Expect(t).
		End().Response
}

func HttpGetAdvertisementsSearch(t *testing.T, search string, limit int) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Create().Get(path.AdvertisementsPath).
		Query("search", search).
		Query("limit", strconv.Itoa(limit)).
		Expect(t).
		End().Response
}

func HttpGetAdvertisementPhotos(t *testing.T, token, id string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Get(fmt.Sprintf(path.AdvertisementPhotosPathTemplate, id)).
		Expect(t).
		End().Response
}

func HttpDeleteAdvertisement(t *testing.T, token, id string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Delete(path.AdvertisementPath).
		Query("id", id).
		Expect(t).
		End().Response
}
