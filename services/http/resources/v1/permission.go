package v1

// PermissionItem ...
type PermissionItem struct {
	ID   int    `structs:"id" json:"id" copier:"field:ID"`
	Name string `structs:"name" json:"name" copier:"field:Name"`
}

// PermissionDetail ...
type PermissionDetail struct {
	ID   int    `structs:"id" json:"id" copier:"field:ID"`
	Name string `structs:"name" json:"name" copier:"field:Name"`
}

// PermissionListItems ..
type PermissionListItems []PermissionItem
