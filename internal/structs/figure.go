package structs

type FigureID string

type Figure struct {
	ID FigureID
	// Figure release title
	Title string
	// Item collection
	Collection string
	// Cloud figure preview hash
	Preview *string
	// Cloud figure animation hash
	Animation *string

	Styles map[int]string

	BaseCost        int
	UpgradeCode     int
	DefaultMetadata ItemMetadata

	UpgradedAmount int
	ReleaseAmount  int

	ReleasedAt int64
	LatestSell int64
}
