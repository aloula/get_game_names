package gamelist

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
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

// Global variables
var totalGamesAllSystems int


// Extract the game names and sort it
func GetGameList(folder string) []string {
	data, err := ioutil.ReadFile("roms/" + folder + "/gamelist.xml")
	if err != nil {
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


// Generate Game List function
func ExtractGameList(romsPath string) error {
	file, err := os.Create("gamelist.txt")
	if err != nil {
		return err
	}
	// List folders (system names) inside roms dir 
	folders, err := ioutil.ReadDir(romsPath + "/.")
	if err != nil {
		return err
	}

	// Extract game names recursively and save it to gamelist.txt file
	for _, f := range folders {
		folders := f.Name()
		gameNames := GetGameList(folders)
		platform := "Platform: " + strings.Title(folders) + "\n\n"
		totalGames := len(gameNames)
		totalGamesAllSystems = totalGamesAllSystems + totalGames
		strTotalGames := "\n- Total Games = " + strconv.Itoa(totalGames)
		if totalGames != 0 {
			file.WriteString(platform)
			for i := 0; i < totalGames; i++ {
				file.WriteString(gameNames[i] + "\n")
			}
			file.WriteString(strTotalGames)
			file.WriteString("\n\n================================================================================\n\n")
		}
	}
	strTotalGamesAllSystems := "- Total Games All Systems = " + strconv.Itoa(totalGamesAllSystems)
	file.WriteString(strTotalGamesAllSystems)
	file.Close()
	return err
}