package startup

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rommel96/mock-auth-server/config"
	"github.com/rommel96/mock-auth-server/database"
	"github.com/rommel96/mock-auth-server/repository"
)

const FILE_PATH = "mock_data.txt"
const PERM_RONLY = 0444

func init() {
	config.LoadVars()
	database.Connect()
	repository.Init()
}

func LoadData() {
	f, err := os.OpenFile(FILE_PATH, os.O_RDONLY, PERM_RONLY)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)
	for {
		line, _, err := br.ReadLine()
		if err != nil {
			break
		}
		var user repository.User
		fmt.Sscanf(string(line), "%s %s %s", &user.Name, &user.Email, &user.Password)
		user.Insert()
	}

}
