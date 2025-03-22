package models

type Photo struct {
	ID        string
	HasAvatar bool
	HasBanner bool
}

type PhotoType string

const (
	PhotoTypePhoto  PhotoType = "PHOTO"
	PhotoTypeAvatar PhotoType = "AVATAR"
	PhotoTypeBanner PhotoType = "BANNER"
)
