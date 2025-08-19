package structs

type ItemID string

type Item struct {
	// Internal ID
	ID ItemID `json:"id" bson:"_id"`
	// Item actual holder
	Holder UserID `json:"holder" bson:"holder"`
	// Indicates whether the item is unique
	// Only unique items can be sold on the stock
	Unique bool `json:"unique" bson:"unique"`

	// Indicates whether the gold frame is displayed
	Gold bool `json:"gold" bson:"gold"`

	// Model Style ID
	Style int `json:"style" bson:"style"`
	// Model of the item
	FigureID FigureID `json:"figure" bson:"figure"`

	// If locked, all non-GET requests from non-admin users involving this item will be refused
	// Usually if item not in storage
	Locked bool `json:"locked" bson:"locked"`
	// Storage name where the item is located
	StorageAddress string `json:"storage" bson:"storage"`
	// Describes minecraft item metadata
	Metadata ItemMetadata `json:"metadata" bson:"metadata"`
	// Item description, cannot be changed after initial setting;
	// Mentioning using <@user_id>
	//
	// Limit is 500 symbols
	Description *string `json:"description" bson:"description"`

	// Item tags
	Tags []string `json:"tags" bson:"tags"`

	// Item visibility status in user profile
	Visible bool `json:"visible" bson:"visible"`
	// Previous holders of the item
	Holders []UserID `json:"holders" bson:"holders"`

	// Locks the ability to trade or sell this item on stock;
	// if -1, locked for indefinitely
	TradeLockedUntil int64 `json:"trade_locked_until" bson:"trade_locked_until"`
	// Locks the ability to send this item to another user;
	TransferLockedUnti int64 `json:"transfer_locked_until" bson:"transfer_locked_until"`
	// Sell timestamp
	SoldAt int64 `json:"sold_at" bson:"sold_at"`
}

// Describes minecraft item metadata
type ItemMetadata struct {
	Signature *string
	Lore      *string
	Timestamp *int64
	Note      *string
}
