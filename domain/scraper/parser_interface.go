package scraper

import (
	"context"

	"github.com/kundu-ramit/dozer_match/domain/entity"
)

type Parser interface {
	Parse(ctx context.Context, html string) (*entity.BullDozer, error)
}
