package libtcolorist

import (
	"fmt"
	"io/fs"
	"os"
)

const PSEP = string(os.PathSeparator)

func multiplatformHomelikePath() string {
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		homeDir = os.Getenv("LOCALAPPDATA")
	}
	return homeDir
}

func initVault() (*os.File, error) {
	homeDir := multiplatformHomelikePath()
	tcoloristDir := homeDir + PSEP + ".tclrst"
	vaultFile := tcoloristDir + PSEP + "vault.tclr"

	_, dirErr := os.Stat(tcoloristDir)
	if os.IsNotExist(dirErr) {
		os.Mkdir(tcoloristDir, 0755)
	}

	_, fileErr := os.Stat(vaultFile)
	if os.IsNotExist(fileErr) {
		err := os.WriteFile(tcoloristDir+PSEP+"vault.tclr", []byte(""), fs.FileMode(os.O_RDWR))
		if err != nil {
			panic(err)
		}
	}
	return os.OpenFile(vaultFile, os.O_RDWR, fs.FileMode(os.O_RDWR))
}

type ColoristVault struct {
	fp *os.File
}

func (cv *ColoristVault) AddColorRGB8(color RGB8) error {
	_, seekErr := cv.fp.Seek(0, 2)
	if seekErr != nil {
		return seekErr
	}
	dataString := fmt.Sprintf("RGB8/%v,%v,%v\n", color.Red, color.Green, color.Blue)
	_, writeErr := cv.fp.WriteString(dataString)
	return writeErr
}

func (cv *ColoristVault) GetColors() error {
	_, err := cv.fp.Seek(0, 0)
	if err != nil {
		return err
	}

	buff := make([]byte, 4096)
	for {
		n, err := cv.fp.Read(buff)
		if err != nil {
			return err
		}
		if n == 0 {
			break
		}
		fmt.Printf("%s", buff[:n])
	}
	return nil
}

func GetVault() *ColoristVault {
	fp, err := initVault()
	if err != nil {
		panic(err)
	}
	vault := ColoristVault{fp: fp}

	return &vault
}
