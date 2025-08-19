package structs

type UserID string

type User struct {
	// Internal User ID
	ID UserID `json:"id" bson:"_id"`
	// User Discord ID
	DiscordID *string `json:"discord_id"  bson:"discord_id"`
	// User SPWorlds API account ID
	SPWorldsID *string `json:"spw_id" bson:"spw_id"`
	// Minecraft Username
	// Can be refreshed, but not changed
	Username string `json:"username" bson:"username"`
	// Minecraft UUID
	// Cannot be changed
	UUID string `json:"uuid" bson:"uuid"`
	// User profile description;
	// Mentioning using <@user_id>
	//
	// Limit is 500 symbols
	Description *string `json:"description" bson:"description"`
	// Cloud banner hash;
	// Premium feature
	Banner *string `json:"banner" bson:"banner"`
	// Displayed near the username;
	// Premium feature
	ProfileIcon *ItemID `json:"icon" bson:"icon"`
	// Items displayed over the user banner;
	// Premium feature
	BannerItems *[]ItemID `json:"banner_items" bson:"banner_items"`
	// User Balance in Dimond Ores;
	// Can be negative
	Balance int `json:"balance" bson:"balance"`
	// A list of user roles
	Roles []RoleID `json:"roles" bson:"roles"`
	// A map of users who cannot suggest trade offers
	//
	// map[UserID]deferred_until (unix time);
	// If -1, user is deferred indefinitely
	//
	// In the future, may change to a pointer if deferring becomes a premium feature
	DeferredUsers map[UserID]int64 `json:"deferred" bson:"deferred"`
	// List of updates
	Updates []Update `json:"updates" bson:"updates"`
	// If a user profile is disabled, all non-GET requests to the API will be refused.
	Disabled bool `json:"disabled" bson:"disabled"`
	// If a user profile is blocked, all interactions with the API will be refused.
	Banned bool `json:"banned" bson:"banned"`
	// Unix timestamp of user creation
	CreatedAt int64 `json:"created_at" bson:"created_at"`
}
