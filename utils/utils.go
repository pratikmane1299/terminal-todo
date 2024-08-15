package utils

import (
	"log"
	"os"

	gap "github.com/muesli/go-app-paths"
)

func SetupPath() string {
	scope := gap.NewScope(gap.User, "terminal.todo")
	dirs, err := scope.DataDirs()

	if err != nil {
		log.Fatal(err)
	}

	var termialDbString string

	if len(dirs) > 0 {
		termialDbString = dirs[0]
	} else {
		termialDbString, _ = os.UserHomeDir()
	}

	if err := initDataDir(termialDbString); err != nil {
		log.Fatal("could not setup data directory path", err)
	}

	return termialDbString
}

func initDataDir(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return os.Mkdir(path, 0o770)
		}

		return err
	}
	return nil
}
