package services

import (
	"errors"
	"go-shortener/models"
	"go-shortener/repositories"
	"go-shortener/utils"
	"net/url"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type URLService struct {
	urlRepo *repositories.UrlRepository
}

func NewURLService(urlRepo *repositories.UrlRepository) *URLService {
	return &URLService{
		urlRepo: urlRepo,
	}
}

func (s *URLService) CreateNewUrl(originalUrl string, userID bson.ObjectID) (*models.Url, error) {
	originalUrl = strings.TrimSpace(originalUrl)

	parsed, err := url.ParseRequestURI(originalUrl)

	if err != nil || parsed.Host == "" || (parsed.Scheme != "http" && parsed.Scheme != "https") {
		return nil, errors.New("Invalid url")
	}

	//loop to avoid collions

	for i := 0; i < 3; i++ {
		code, err := utils.GenerateShortCode()

		if err != nil {
			return nil, err
		}

		record := &models.Url{
			UserID:      userID,
			OriginalUrl: originalUrl,
			ShortCode:   code,
			CreatedAt:   time.Now().UTC(),
		}

		err = s.urlRepo.Create(record)
		if err == nil {
			return record, nil
		}

		if !errors.Is(err, utils.ErrUrlExists) {
			return nil, err
		}
	}

	return nil, errors.New("Could not generate a unique code")
}

func (s *URLService) RedirectURL(code string) (*models.Url, error) {
	code = strings.TrimSpace(code)
	if code == "" {
		return nil, mongo.ErrNoDocuments
	}

	record, err := s.urlRepo.FindByShortCode(code)
	if err != nil {
		return nil, err
	}

	if err := s.urlRepo.IncrementClickCount(code); err != nil {
		return nil, err
	}

	record.Clicks++
	return record, nil
}
