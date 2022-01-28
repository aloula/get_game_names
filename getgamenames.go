/*
Read gamelist.xml files recursively and extract the Game Name saving it to a gamelist.txt file
Code by: Alexsander Loula
Email: alex.loula@gmail.com
Usage: getgamenames <path to roms file>
*/

package main

import (
	"fmt"
	"get_game_names/pkg"
	"github.com/common-nighthawk/go-figure"
	"os"
)

// Main
func main() {
	// print ascii art
	myFigure := figure.NewColorFigure("Retro Gaming", "rectangles", "green", true)
	myFigure.Print()
	fmt.Println()
	// check command line arguments
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <path to roms file> \n ", os.Args[0])
		fmt.Printf("Example: %s /recalbox/share/roms \n ", os.Args[0])
		fmt.Printf("Example: %s roms \n", os.Args[0])
		os.Exit(1)
	}
	romsPath := os.Args[1]
	gamelist.ExtractGameList(romsPath)
	os.Exit(0)
}
