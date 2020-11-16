package database

import (
	"github.com/Kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(baseURL string) error {

	err := mgm.SetDefaultConfig(nil, "test", options.Client().ApplyURI(baseURL))

	if err != nil {
		return err
	}

	return nil
}
