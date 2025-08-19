package structs

type Update struct {
	Label       string `json:"label" bson:"label"`
	Description string `json:"description" bson:"ref"`
	// TODO: add type enum
	Type string `json:"type" bson:"ref"`
	// Link to update
	Ref string `json:"ref" bson:"ref"`
	// Update timestamp (unix time)
	Timestamp int64 `json:"timestamp" bson:"timestamp"`
}
