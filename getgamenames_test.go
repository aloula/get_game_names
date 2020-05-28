package main

import (
	"io/ioutil"
	"os"
	"testing"
)

// test main
// func TestMain(m *testing.M) {
//     exitcode := m.Run()
// 	os.Exit(exitcode)
// }

// test game list functions
func TestGameList(t *testing.T) {
	// test total games
	gameList := getGameList("megadrive")
	if len(gameList) != 94 {
		t.Errorf("Não leu arquivo corretamente")
	}
	// test game name
	gameNames := gameList[0]
	if gameNames != "Aero the Acro-Bat" {
		t.Errorf("Não extraiu nome corretamente")
	}
	// test game list error
	gameListerr := getGameList("notvalid")
	if gameListerr != nil {
		t.Errorf("Deveria retornar nil pois path é inválido")
	}
}

// test extract list functions
func TestExtractList(t *testing.T) {
	// test gamelist.txt creation
	os.Chmod("gamelist.txt", 0600)
	extractGameList("roms")
	gameListFile, err := ioutil.ReadFile("gamelist.txt")
	//fmt.Println(len(gameListFile))
	//fmt.Println(err)
	if len(gameListFile) != 5872 || err != nil{		
		t.Errorf("Não leu diretório corretamente")
	}
	//test if gamelist.txt can be created
	os.Chmod("gamelist.txt", 0400)
	expectedCreationErr := "open gamelist.txt: permission denied"
	creationErr := extractGameList("roms") 
	if creationErr.Error() != expectedCreationErr{
		t.Errorf("Não era para criar gamelist.txt")
	}
	//test if roms path exist
	os.Chmod("gamelist.txt", 0600)
	expectedRomsErr := "open rom/.: no such file or directory"
	romsErr := extractGameList("rom") 
	if romsErr.Error() != expectedRomsErr{
		t.Errorf("Não era para criar gamelist.txt")
	}
}