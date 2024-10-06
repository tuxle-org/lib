package entities

type User struct {
	Entity
	Name   string `gorm:"uniqueIndex"`
	RoleId uint   `gorm:"not null"`
	Role   Role
	Tags   []*Tag `gorm:"many2many:user_tags"`
}

func NewUser(name string, roleId uint) *User {
	return &User{
		Entity: Entity{},
		Name:   name,
		RoleId: roleId,
		Role:   Role{},
		Tags:   []*Tag{},
	}
}

type Role struct {
	Entity
	Name             string         `gorm:"uniqueIndex"`
	Color            uint32         `gorm:"default:0"`
	IconURI          *string        `gorm:"size:255"`
	Permissions      PermissionMask `gorm:"foreignKey:FkPermissionMask"`
	FkPermissionMask uint
}

func NewRole(name string, color uint32, iconURI string, fkPermission uint) *Role {
	return &Role{
		Entity:           Entity{},
		Name:             name,
		Color:            color,
		IconURI:          &iconURI,
		Permissions:      PermissionMask{},
		FkPermissionMask: fkPermission,
	}
}

type Tag struct {
	Entity
	Name             string         `gorm:"uniqueIndex"`
	Priority         uint32         `gorm:"default:0"`
	Permissions      PermissionMask `gorm:"foreignKey:FkPermissionMask"`
	FkPermissionMask uint
}

func NewTag(name string, priority uint32, fkPermission uint) *Tag {
	return &Tag{
		Entity:           Entity{},
		Name:             name,
		Priority:         priority,
		Permissions:      PermissionMask{},
		FkPermissionMask: fkPermission,
	}
}
