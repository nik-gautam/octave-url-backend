package handlers

import (
	"github.com/Kamva/mgm/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/nik-gautam/octave-url-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
)

type RequestUrl struct {
	LongUrl       string `json:"longUrl"`
	CustomUrlCode string `json:"urlCode"`
}

type PatchUrl struct {
	Id      primitive.ObjectID `json:"id"`
	UrlCode string             `json:"urlCode"`
	LongUrl string             `json:"longUrl"`
}

func GetAllUrls(c *fiber.Ctx) error {
	var allUrl []models.Urls

	if err := mgm.Coll(&models.Urls{}).SimpleFind(&allUrl, bson.M{}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"msg":     "Error while retrieving Data from DB",
			"err":     err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"urls": allUrl,
	})
}

func GetLongUrl(c *fiber.Ctx) error {
	shortCode := c.Params("shortCode")

	urlColl := mgm.CollectionByName("urls")

	url := &models.Urls{}

	err := urlColl.First(bson.M{"urlCode": shortCode}, url)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"msg":     "Error while retrieving Data from DB",
			"err":     err.Error(),
		})
	}

	url.Count++

	if err := urlColl.Update(url); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"msg":     "Unable to update count to DB",
			"err":     err.Error(),
		})
	}

	return c.Redirect(url.LongUrl)
}

func PostAddUrl(c *fiber.Ctx) error {
	reqUrl := new(RequestUrl)

	if err := c.BodyParser(reqUrl); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"msg":     "Internal Body Parser Error",
			"err":     err.Error(),
		})
	}

	newUrl := models.CreateUrl(reqUrl.LongUrl, reqUrl.CustomUrlCode, os.Getenv("BASE_URL"))

	urlColl := mgm.Coll(newUrl)

	if err := urlColl.Create(newUrl); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"msg":     "Unable to save in DB",
			"err":     err.Error(),
		})
	}

	return c.JSON(newUrl)
}

func PatchEditUrl(c *fiber.Ctx) error {

	urlColl := mgm.CollectionByName("urls")
	existingUrl := &models.Urls{}

	reqUrl := new(PatchUrl)

	if err := c.BodyParser(reqUrl); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"msg":     "Internal Body Parser Error",
			"err":     err.Error(),
		})
	}

	if err := urlColl.FindByID(reqUrl.Id, existingUrl); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"msg":     "Url not found in DB",
			"err":     err.Error(),
		})
	}

	if reqUrl.LongUrl != "" {
		existingUrl.LongUrl = reqUrl.LongUrl
	}

	if reqUrl.UrlCode != "" {
		existingUrl.UrlCode = reqUrl.UrlCode
		existingUrl.ShortUrl = os.Getenv("BASE_URL") + reqUrl.UrlCode
	}

	if err := urlColl.Update(existingUrl); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"msg":     "Unable to save in DB",
			"err":     err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"updatedUrl": existingUrl,
	})
}

func DeleteUrl(c *fiber.Ctx) error {

	urlColl := mgm.CollectionByName("urls")
	existingUrl := &models.Urls{}

	urlId := c.Params("id")

	if err := urlColl.FindByID(urlId, existingUrl); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"msg":     "Url not found in DB",
			"err":     err.Error(),
		})
	}

	if err := urlColl.Delete(existingUrl); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"msg":     "Unable to delete in DB",
			"err":     err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"deletedUrl": existingUrl,
	})
}
