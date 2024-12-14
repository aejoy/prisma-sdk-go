package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/aejoy/go-pkg/utils"
	"github.com/aejoy/prisma-sdk-go/api"
	"github.com/aejoy/prisma-sdk-go/types"
)

func main() {
	instance := api.New("http://localhost:4810/api/v1")

	file, err := os.ReadFile("./samples/Poster-Design-20240301-3-1.jpeg")
	if err != nil {
		panic(err)
	}

	uploaded, err := instance.Upload("photo", bytes.NewReader(file))
	if err != nil {
		panic(err)
	}

	fmt.Println(uploaded.Photo.URL)

	photos, err := instance.GetPhotos(types.GetPhotos{
		Count: utils.PointerValue(1),
	})
	if err != nil {
		panic(err)
	}

	for _, photo := range photos.Photos {
		fmt.Printf("Photo(\n\tID: %s,\n\tURL: %s,\n\theight: %d,\n\twidth: %d,\n\tsize: %f,\n\tblurHash: %s\n)\n", photo.ID, photo.URL, photo.Height, photo.Width, photo.SizeInKB, photo.BlurHash)
	}
}
