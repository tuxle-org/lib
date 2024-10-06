package entities

type Server struct {
	Entity
	Name        string  `gorm:"not null;uniqueIndex"`
	Description string  `gorm:"type:text;not null"`
	Rules       string  `gorm:"type:text;not null"`
	IconURI     *string `gorm:"size:255"`
	BannerURI   *string `gorm:"size:255"`
	OwnerId     uint    `gorm:"not null"`
	Region      string  `gorm:"size:8;not null"`
}

func NewServer(
	name string,
	description string,
	rules string,
	iconURI string,
	bannerURI string,
	ownerId uint,
	region string,
) *Server {
	return &Server{
		Entity:      Entity{},
		Name:        name,
		Description: description,
		Rules:       rules,
		IconURI:     &iconURI,
		BannerURI:   &bannerURI,
		OwnerId:     ownerId,
		Region:      region,
	}
}
