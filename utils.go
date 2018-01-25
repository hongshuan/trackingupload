package main

import (
    "os"
    "log"
)

func CheckError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// Exists reports whether the named file or directory exists.
func FileExists(name string) bool {
    _, err := os.Stat(name)
    if os.IsNotExist(err) {
        return false
    }
    return err == nil
}

func ExistFile(name string) bool {
    _, err := os.Stat(name)
    return !os.IsNotExist(err)
}

func ExistsFile(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err != nil, err
}

func FileIsExists(name string) bool {
    if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
		// ??
    }
    return true
}
