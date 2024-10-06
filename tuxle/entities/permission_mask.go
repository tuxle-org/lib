package entities

type PermissionMask struct {
	Entity
	CanSendMessage *bool `gorm:"default:null"`
}

func NewPermissionMask(
	canSendMessage bool,
) *PermissionMask {
	return &PermissionMask{
		CanSendMessage: &canSendMessage,
	}
}
