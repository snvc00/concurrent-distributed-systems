package image

import "fmt"

type Image struct {
	Title string
	Format string
	Channels string
}

func (image *Image) Show() {
	fmt.Println("Image { Title:", image.Title, "\b, Format:", image.Format, "\b, Channels:", image.Channels, "}")
}