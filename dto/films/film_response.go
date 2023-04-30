package filmsdto

import "dumbflix/models"

type FilmResponse struct {
	Id            int             `json:"id" `
	Title         string          `json:"title" `
	ThumbnailFilm string          `json:"thumbnailFilm" `
	Year          int             `json:"year" gorm:"type:int(4)"`
	Category      models.Category `json:"category"`
	Description   string          `json:"description" `
}
type FilmDeleteResponse struct {
	Id int `json:"id" `
}
