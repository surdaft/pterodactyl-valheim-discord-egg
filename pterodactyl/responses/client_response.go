package responses

import (
	"github.com/surdaft/pterodactyl-valheim-discord-egg/pterodactyl/entities"
)

type ClientResponse struct {
	Data []entities.Server `json:"data"`
}