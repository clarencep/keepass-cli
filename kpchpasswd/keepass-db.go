package main

import (
	"errors"
	"os"

	gkp "github.com/tobischo/gokeepasslib"
)

type KeepassDbFile struct {
	file *os.File
	db   *gkp.Database
}

func OpenKeepassDb(filePath string) (f *KeepassDbFile, err error) {
	f = new(KeepassDbFile)
	f.file, err = os.Open(filePath)
	if err != nil {
		return
	}

	return
}

func (f *KeepassDbFile) UnlockWithPassword(passwd string) error {
	f.db = gkp.NewDatabase()
	f.db.Credentials = gkp.NewPasswordCredentials(passwd)
	err := gkp.NewDecoder(f.file).Decode(f.db)
	if err != nil {
		return err
	}

	f.db.UnlockProtectedEntries()
	return nil
}

func (f *KeepassDbFile) ChPassword(newPasswd string) error {
	if f.db == nil {
		return errors.New("Database not open/unlock yet!")
	}

	f.db.Credentials = gkp.NewPasswordCredentials(newPasswd)
	return nil
}

func (f *KeepassDbFile) SaveTo(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	f.db.LockProtectedEntries()
	defer f.db.UnlockProtectedEntries()

	err = gkp.NewEncoder(file).Encode(f.db)
	if err != nil {
		return err
	}

	return nil
}

func (f *KeepassDbFile) CloseFile() {
	if f.file != nil {
		f.file.Close()
		f.file = nil
	}
}

func (f *KeepassDbFile) CloseDb() {
	if f.db != nil {
		f.db = nil
	}
}

func (f *KeepassDbFile) Close() {
	f.CloseFile()
	f.CloseDb()
}
