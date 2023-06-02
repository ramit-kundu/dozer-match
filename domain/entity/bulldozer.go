package entity

type BullDozer struct {
	Make            string
	Model           string
	Picture         string
	Category        string //should be made enum future scope
	EngineHP        string
	OperatingWeight int64
	ScrapeIndex     int64
}
