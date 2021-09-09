package video

import "fmt"

type Video struct {
	Title string
	Format string
	Fps float32
}

func (video *Video) Show() {
	fmt.Println("Video { Title:", video.Title, "\b, Format:", video.Format, "\b, Fps:", video.Fps, "}")
}
