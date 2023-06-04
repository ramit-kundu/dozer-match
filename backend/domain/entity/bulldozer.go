package entity

type BullDozer struct {
	Make            string `validate:"required"`
	Model           string `validate:"required"`
	Picture         string `validate:"required"`
	Category        string `validate:"required"`
	EngineHP        string `validate:"required"`
	OperatingWeight string `validate:"required"`
	ScrapeIndex     string
}
