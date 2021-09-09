package audio

import "fmt"

type Audio struct {
	Title string
	Format string
	Duration int
}

func (audio *Audio) Show() {
	fmt.Println("Audio { Title:", audio.Title, "\b, Format:", audio.Format, "\b, Duration:", audio.Duration, "}")
}