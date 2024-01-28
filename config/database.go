package config

import (
	"gogo/prisma/db"

	"github.com/rs/zerolog/log"
)

func ConnectDB() (*db.PrismaClient, error) {
	client := db.NewClient()
	err := client.Connect()
	if err != nil {
		return nil, err
	}
	log.Info().Msg("Connected to DB")
	return client, nil
}
