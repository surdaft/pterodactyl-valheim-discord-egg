package pterodactyl_valheim_discord_egg

import (
	"github.com/surdaft/pterodactyl-valheim-discord-egg/pterodactyl"
)

type Config struct {
	Pterodactyl pterodactyl.Config `yaml:"pterodactyl"`
}
