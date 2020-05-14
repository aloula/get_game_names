/* 
Read gamelist.xml files recursively and extract the Game Name saving it to a gamelist.txt file
Code by: Alexsander Loula
Email: alex.loula@gmail.com
Usage: getgamenames <path to roms file> 
*/ 

package main

import (
	"encoding/xml"
	"fmt"
	"github.com/cheggaaa/pb"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"
)

// Gamelist XML Struct
type gameList struct {
	XMLName xml.Name `xml:"gameList"`
	Text    string   `xml:",chardata"`
	Game    []struct {
		Text        string `xml:",chardata"`
		Path        string `xml:"path"`
		Name        string `xml:"name"`
		Sortname    string `xml:"sortname"`
		Desc        string `xml:"desc"`
		Rating      string `xml:"rating"`
		Releasedate string `xml:"releasedate"`
		Developer   string `xml:"developer"`
		Publisher   string `xml:"publisher"`
		Genre       string `xml:"genre"`
		Players     string `xml:"players"`
		Image       string `xml:"image"`
		Playcount   string `xml:"playcount"`
		Lastplayed  string `xml:"lastplayed"`
		Md5         string `xml:"md5"`
		Hash        string `xml:"hash"`
		Core        string `xml:"core"`
		Emulator    string `xml:"emulator"`
		Video       string `xml:"video"`
		Hidden      string `xml:"hidden"`
		Favorite    string `xml:"favorite"`
	} `xml:"game"`
	Folder struct {
		Text   string `xml:",chardata"`
		Path   string `xml:"path"`
		Image  string `xml:"image"`
		Hidden string `xml:"hidden"`
		Name   string `xml:"name"`
	} `xml:"folder"`
} 


// Function to extract the game names and sort it
func getGameList(folder string) []string {
	data, err := ioutil.ReadFile("roms/" + folder + "/gamelist.xml")
	if err != nil {
		//fmt.Println(err)
		return nil
	}
	gamelist := &gameList{}
	_ = xml.Unmarshal([]byte(data), &gamelist)
	gameNames := []string{}
	for i := 0; i < len(gamelist.Game); i++ {
		// Check if game is not hidden
		if gamelist.Game[i].Hidden != "true" { 
			gameNames = append(gameNames, gamelist.Game[i].Name)
		}
	}
	sort.Strings(gameNames)
	return gameNames
}

// Main
func main() {
	// Check arguments
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <path to roms file> \n ", os.Args[0])
		fmt.Printf("Example: %s /recalbox/share/roms \n ", os.Args[0])
		fmt.Printf("Example: %s roms \n", os.Args[0])
		os.Exit(1)
	}
	romsPath := os.Args[1]
	file, err := os.Create("gamelist.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// List folders (system names) inside roms dir 
	folders, err := ioutil.ReadDir(romsPath + "/.")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// var countSize int = int(fileSize/1000)
	// count := countSize
	// bar := pb.StartNew(count)
	// for i := 0; i < count; i++ {
    // 	bar.Increment()
    // 	time.Sleep(time.Millisecond)
	// }
	// bar.FinishPrint("The End!")

	// Extract game names recursively and save it to gamelist.txt file
	fmt.Println("Starting game names extraction...")
	for _, f := range folders {
		folders := f.Name()
		gameNames := getGameList(folders)
		bar := pb.StartNew(len(gameNames))
		if len(gameNames) != 0 {
			file.WriteString("================================================================================\n")
			file.WriteString(strings.Title(folders))
			file.WriteString("\n================================================================================\n")
			for i := 0; i < len(gameNames); i++ {
				//fmt.Println(gameNames[i])
				file.WriteString(gameNames[i] + "\n")
				bar.Increment()
				time.Sleep(time.Millisecond*5)
			}
		}
	bar.Finish()
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
