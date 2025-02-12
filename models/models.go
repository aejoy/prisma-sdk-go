package models

type Photo struct {
	ID string
}

type PhotoType string

const (
	PhotoTypePhoto  PhotoType = "PHOTO"
	PhotoTypeAvatar PhotoType = "AVATAR"
	PhotoTypeBanner PhotoType = "BANNER"
)
