package gamelist_test

import (
	"get_game_names/pkg"
	"io/ioutil"
	"os"
	"testing"
)

// constants
const totalGames = 78
const gameName = "Aladdin"
const fileSize = 2867


// // test game list functions
// func TestGameList(t *testing.T) {
// 	// test total games
// 	gameList := gamelist.GetGameList("gamelists", "megadrive")
// 	if len(gameList) != totalGames {
// 		t.Errorf("Can't read the test file")
// 	}
// 	// test game name
// 	gameNames := gameList[0]
// 	if gameNames != gameName {
// 		t.Errorf("Can't extract game name")
// 	}
// 	// test game list error
// 	gameListerr := gamelist.GetGameList("notvalid", "notvalid")
// 	if gameListerr != nil {
// 		t.Errorf("It should return nil due to the invalid path")
// 	}
// }

// test extract list functions
func TestExtractList(t *testing.T) {
	// test gamelist.txt creation
	os.Chmod("gamelist.txt", 0600)
	gamelist.ExtractGameList("gamelists")
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