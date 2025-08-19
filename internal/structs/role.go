package structs

type RoleID string

type Role struct {
	ID    RoleID
	Title string
	Color string
	// TODO: Implement permissions
}
