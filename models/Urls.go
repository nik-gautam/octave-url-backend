package models

import (
	"github.com/Kamva/mgm/v3"
	"github.com/teris-io/shortid"
)

type Urls struct {
	mgm.DefaultModel `bson:",inline"`
	ShortUrl         string `json:"shortUrl" bson:"shortUrl"`
	LongUrl          string `json:"longUrl" bson:"longUrl"`
	Count            int    `json:"count" bson:"count"`
	UrlCode          string `json:"urlCode" bson:"urlCode"`
}

func generateCode() (string, error) {
	return shortid.Generate()
}

func CreateUrl(longUrl string, customCode string, baseUrl string) *Urls {
	shortCode := ""

	if customCode != "" {
		shortCode = customCode
	} else {
		shortCode, _ = generateCode()
	}

	return &Urls{
		ShortUrl: baseUrl + shortCode,
		LongUrl:  longUrl,
		Count:    0,
		UrlCode:  shortCode,
	}
}
