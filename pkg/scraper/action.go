package scraper

import "github.com/stashapp/stash/pkg/models"

type scraperAction string

const (
	scraperActionScript scraperAction = "script"
	scraperActionStash  scraperAction = "stash"
	scraperActionXPath  scraperAction = "scrapeXPath"
	scraperActionJson   scraperAction = "scrapeJson"
)

var allScraperAction = []scraperAction{
	scraperActionScript,
	scraperActionStash,
	scraperActionXPath,
	scraperActionJson,
}

func (e scraperAction) IsValid() bool {
	switch e {
	case scraperActionScript, scraperActionStash, scraperActionXPath, scraperActionJson:
		return true
	}
	return false
}

type scrapeOptions struct {
	scraper      scraperTypeConfig
	config       config
	globalConfig GlobalConfig
}

type scraper interface {
	scrapePerformersByName(name string) ([]*models.ScrapedPerformer, error)
	scrapePerformerByFragment(scrapedPerformer models.ScrapedPerformerInput) (*models.ScrapedPerformer, error)
	scrapePerformerByURL(url string) (*models.ScrapedPerformer, error)

	scrapeSceneByFragment(scene models.SceneUpdateInput) (*models.ScrapedScene, error)
	scrapeSceneByURL(url string) (*models.ScrapedScene, error)

	scrapeMovieByURL(url string) (*models.ScrapedMovie, error)
}

func getScraper(scraper scraperTypeConfig, config config, globalConfig GlobalConfig) scraper {
	switch scraper.Action {
	case scraperActionScript:
		return newScriptScraper(scraper, config, globalConfig)
	case scraperActionStash:
		return newStashScraper(scraper, config, globalConfig)
	case scraperActionXPath:
		return newXpathScraper(scraper, config, globalConfig)
	case scraperActionJson:
		return newJsonScraper(scraper, config, globalConfig)
	}

	panic("unknown scraper action: " + scraper.Action)
}
