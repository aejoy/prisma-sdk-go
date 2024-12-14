package types

type GetPhotos struct {
	Count    *int
	Offset   *int
	PhotoIDs []string
}
