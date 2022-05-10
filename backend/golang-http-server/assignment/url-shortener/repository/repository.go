package repository

import (
	"sync"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/entity"
)

type URLRepository struct {
	mu   sync.Mutex
	Data map[string]string
}

func NewMapRepository() URLRepository {
	data := make(map[string]string, 0)
	return URLRepository{
		Data: data,
	}
}

func (r *URLRepository) Get(path string) (*entity.URL, error) {

	url, ok := r.Data[path]
	if ok {
		return &entity.URL{
			LongURL:  url,
			ShortURL: path,
		}, nil
	}
	return &entity.URL{}, entity.ErrURLNotFound

}

func (r *URLRepository) Create(longURL string) (*entity.URL, error) {
	shortURL := entity.GetRandomShortURL(longURL)
	r.Data[shortURL] = longURL
	return &entity.URL{
		LongURL:  longURL,
		ShortURL: shortURL,
	}, nil
}

func (r *URLRepository) CreateCustom(longURL, customPath string) (*entity.URL, error) {
	customUrl, ok := r.Data[customPath]
	if ok {
		return &entity.URL{
			LongURL:  customUrl,
			ShortURL: customPath,
		}, entity.ErrCustomURLIsExists
	}

	r.Data[customPath] = longURL
	return &entity.URL{
		LongURL:  longURL,
		ShortURL: customPath,
	}, nil
}
