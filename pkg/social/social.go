package social

import "github.com/tks98/DetectBot/pkg/botdetector"

// Media describes what methods social media types require to run the bot detection
type Media interface {
	GetFeatures() (botdetector.Features, error)
}
