package main

import (
    "os"
    "log"
    "time"
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

// today := time.Now().Format("2006-01-02")
// modTime := FileModTime(filename).Format("2006-01-02")

func FileModTime(filename string) time.Time {
	info, err := os.Stat(filename)
	if err != nil {
		 return time.Time{}
	}

	return info.ModTime()
}

func GetFileSize(filename string) int64 {
	info, err := os.Stat(filename)
	if err != nil {
		 return 0
	}

	return info.Size()
}

func SameDay(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year() &&
           t1.Month() == t2.Month() &&
		   t1.Day() == t2.Day()
}
