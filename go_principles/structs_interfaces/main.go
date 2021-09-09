package main

import (
	"bufio"
	"fmt"
	"os"
	. "structs_interfaces/multimedia"
	. "structs_interfaces/multimedia/audio"
	. "structs_interfaces/multimedia/image"
	. "structs_interfaces/multimedia/video"
	. "structs_interfaces/web_content"
)

func main() {
	var option string
	webContent := WebContent{Content: []Multimedia{}}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\033[H\033[2J")

	for option != "5" {
		fmt.Print("Main menu\n1. Add image\n2. Add audio\n3. Add video\n4. Show\n5. Exit\n\nOption: ")
		fmt.Scanln(&option)

		switch option {
			case "1":
				var title, format, channels string
				fmt.Print("Title: ")
				scanner.Scan()
				title = scanner.Text()
				fmt.Print("Format: ")
				fmt.Scanln(&format)
				fmt.Print("Channels: ")
				fmt.Scanln(&channels)
				webContent.Content = append(webContent.Content, &Image{Title: title, Format: format, Channels: channels})
			case "2":
				var title, format string
				var duration int
				fmt.Print("Title: ")
				scanner.Scan()
				title = scanner.Text()
				fmt.Print("Format: ")
				fmt.Scanln(&format)
				fmt.Print("Duration: ")
				fmt.Scanln(&duration)
				webContent.Content = append(webContent.Content, &Audio{Title: title, Format: format, Duration: duration})
			case "3":
				var title, format string
				var fps float32
				fmt.Print("Title: ")
				scanner.Scan()
				title = scanner.Text()
				fmt.Print("Format: ")
				fmt.Scanln(&format)
				fmt.Print("Fps: ")
				fmt.Scanln(&fps)
				webContent.Content = append(webContent.Content, &Video{Title: title, Format: format, Fps: fps})
			case "4":
				for _, multimedia := range webContent.Content {
					multimedia.Show()
				}
			case "5":
				return
			default:
				fmt.Println("Invalid option")
		}

		fmt.Println("\nEnter to continue")
		fmt.Scan()
		fmt.Print("\033[H\033[2J")
	}

}
