package models

import (
	"github.com/Kamva/mgm/v3"
)

type Urls struct {
	mgm.DefaultModel `bson:",inline"`
	ShortUrl         string `json:"shortUrl" bson:"shortUrl"`
	LongUrl          string `json:"longUrl" bson:"longUrl"`
	Count            int    `json:"count" bson:"count"`
	UrlCode          string `json:"urlCode" bson:"urlCode"`
}

func CreateUrl(longUrl string, shortCode string, baseUrl string) *Urls {
	return &Urls{
		ShortUrl: baseUrl + shortCode,
		LongUrl:  longUrl,
		Count:    0,
		UrlCode:  shortCode,
	}
}
