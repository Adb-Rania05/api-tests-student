package path

const (
<<<<<<< HEAD
	AuthLoginPath                   = "/auth/login"
	MyAdvertisementsPath            = "/my/advertisements"
	AdvertisementPath               = "/advertisement"
	AdvertisementsPath              = "/advertisements"
	AdvertisementPhotosPathTemplate = "/advertisements/%s/photos"
	FavoritePath                    = "/favorites"
	FavoritesPath                   = "/favorites"
	FavoritePathTemplate            = "/favorites/%s"
=======
	// AuthLoginPath — эндпоинт авторизации
	AuthLoginPath = "/auth/login"

	// MyAdvertisementsPath — список моих объявлений
	MyAdvertisementsPath = "/my/advertisements"

	// AdvertisementPath — создание/получение/удаление объявления
	AdvertisementPath = "/advertisement"

	// AdvertisementsPath — поиск объявлений
	AdvertisementsPath = "/advertisements"

	// AdvertisementPhotosPathTemplate — фотографии объявления
	AdvertisementPhotosPathTemplate = "/advertisements/%s/photos"

	// FavoritePath — добавление/удаление из избранного
	FavoritePath = "/favorites"

	// FavoritesPath — список избранного
	FavoritesPath = "/favorites"

	// FavoritePathTemplate — удаление конкретного объявления из избранного
	FavoritePathTemplate = "/favorites/%s"
>>>>>>> b13fc60 (final)
)
