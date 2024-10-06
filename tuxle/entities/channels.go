package entities

const TEXT_CHANNEL = 0

type Directory struct {
	Entity
	Name             string
	ParentId         *uint          `gorm:"not null"`
	Permissions      PermissionMask `gorm:"foreignKey:FkPermissionMask"`
	FkPermissionMask uint
}

func NewDirectory(name string, parentId uint, fkPerm uint) *Directory {
	return &Directory{
		Entity:           Entity{},
		Name:             name,
		ParentId:         &parentId,
		Permissions:      PermissionMask{},
		FkPermissionMask: fkPerm,
	}
}

type Channel struct {
	Entity
	Name             string
	Type             uint8          `gorm:"not null"`
	Permissions      PermissionMask `gorm:"foreignKey:FkPermissionMask"`
	Directory        Directory      `gorm:"foreignKey:FkDirectory"`
	FkPermissionMask uint
	FkDirectory      uint
}

func NewChannel(name string, t uint8, fkPermission uint, fkDirectory uint) *Channel {
	return &Channel{
		Entity:           Entity{},
		Name:             name,
		Type:             t,
		Permissions:      PermissionMask{},
		Directory:        Directory{},
		FkPermissionMask: fkPermission,
		FkDirectory:      fkDirectory,
	}
}

type TextMessage struct {
	Entity
	Content   string  `gorm:"type:text"`
	User      User    `gorm:"foreignKey:FkUser"`
	Channel   Channel `gorm:"foreignKey:FkChannel"`
	FkUser    uint
	FkChannel uint
}

func NewTextMessage(content string, fkUser uint, fkChannel uint) *TextMessage {
	return &TextMessage{
		Entity:    Entity{},
		Content:   content,
		User:      User{},
		Channel:   Channel{},
		FkUser:    fkUser,
		FkChannel: fkChannel,
	}
}

type MessageVote struct {
	Entity
	IsPositive bool    `gorm:"not null;default:0"`
	MessageId  uint    `gorm:"not null"`
	Channel    Channel `gorm:"foreignKey:FkChannel"`
	User       User    `gorm:"foreignKey:FkUser"`
	FkChannel  uint
	FkUser     uint
}

func NewMessageVote(isPositive bool, messageId uint, fkChannel uint, fkUser uint) *MessageVote {
	return &MessageVote{
		Entity:     Entity{},
		IsPositive: isPositive,
		MessageId:  messageId,
		Channel:    Channel{},
		User:       User{},
		FkChannel:  fkChannel,
		FkUser:     fkUser,
	}
}
