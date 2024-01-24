package main

import "github.com/rs/zerolog/log"

func main() {
	log.Info().Str("planet", "earth").Msg("Hello World!")
}
