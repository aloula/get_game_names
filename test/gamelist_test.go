package gamelist_test

import (
	"get_game_names/pkg"
	"io/ioutil"
	"os"
	"testing"
)

// test parameters
var systemFolder = "mega"
var firstGame = "Aero the Acro-Bat"
var lastGame = "World Heroes"
var lastIndex = 92
var fileSize = 1956


// test game list functions
func TestGameList(t *testing.T) {
	// test total games
	gameList := gamelist.GetGameList(systemFolder)
	if len(gameList) != lastIndex + 1 {
		t.Errorf("Total games is wrong!")
	}
	// test game name
	firstGameName := gameList[0]
	if firstGameName != firstGame {
		t.Errorf("First game should be:" + firstGame)
	}
	lastGameName := gameList[lastIndex]
	if lastGameName != lastGame {
		t.Errorf("Last game should be:" + lastGame )
	}
	// test game list error
	gameListerr := gamelist.GetGameList("notvalid")
	if gameListerr != nil {
		t.Errorf("It should return nil due to the invalid path")
	}
}

//test extract list functions
func TestExtractList(t *testing.T) {
	// test gamelist.txt was created
	os.Chmod("gamelist.txt", 0600)
	gamelist.ExtractGameList("roms")
	gameListFile, err := ioutil.ReadFile("gamelist.txt")
	if len(gameListFile) != fileSize || err != nil{		
		t.Errorf("Can't read the folder correctly")
	}
	//test if gamelist.txt can be created
	os.Chmod("gamelist.txt", 0400)
	expectedCreationErr := "open gamelist.txt: permission denied"
	creationErr := gamelist.ExtractGameList("roms") 
	if creationErr.Error() != expectedCreationErr{
		t.Errorf("It shouldn't create gamelist.txt file")
	}
	//test if roms path exist
	os.Chmod("gamelist.txt", 0600)
	expectedRomsErr := "open rom/.: no such file or directory"
	romsErr := gamelist.ExtractGameList("rom") 
	if romsErr.Error() != expectedRomsErr{
		t.Errorf("It shouldn't create gamelist.txt file")
	}
}