package gamelist_test

import (
	//"fmt"
	"get_game_names/pkg"
	"io/ioutil"
	"os"
	"testing"
)


// test game list functions
func TestGameList(t *testing.T) {
	// test total games
	gameList := gamelist.GetGameList("megadrive")
	//fmt.Println(gameList)
	if len(gameList) != 94 {
		t.Errorf("Can't read the test file")
	}
	// test game name
	gameNames := gameList[0]
	if gameNames != "Aero the Acro-Bat" {
		t.Errorf("Can't extract game name")
	}
	// test game list error
	gameListerr := gamelist.GetGameList("notvalid")
	if gameListerr != nil {
		t.Errorf("It should return nil due to the invalid path")
	}
}

// test extract list functions
func TestExtractList(t *testing.T) {
	// test gamelist.txt creation
	os.Chmod("gamelist.txt", 0600)
	gamelist.ExtractGameList("roms")
	gameListFile, err := ioutil.ReadFile("gamelist.txt")
	// fmt.Println(len(gameListFile))
	// fmt.Println(err)
	if len(gameListFile) != 2017 || err != nil{		
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