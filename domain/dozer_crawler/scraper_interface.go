package scraper

import (
	"context"

	"github.com/kundu-ramit/dozer_match/domain/entity"
)

type Scraper interface {
	ScrapePage(ctx context.Context) ([]entity.BullDozer, error)
}
