package entity

type BullDozer struct {
	Make            string
	Model           string
	Picture         string
	Category        DozerCategory
	EngineHP        string
	OperatingWeight int64
}

type DozerCategory int

const (
	SmallDozer DozerCategory = iota
	MediumDozer
	LargeDozer
	WheelDozer
)
