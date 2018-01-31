package main

import (
    "fmt"
    "os"
    "log"
    "time"
    "strings"
    "path/filepath"
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

func GetFileTime(name string) (mtime time.Time, err error) {
    fi, err := os.Stat(name)
    if err != nil {
        return
    }
    mtime = fi.ModTime()
    return
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

// Y-m-d H:i:s => 2006-01-02 15:04:05
var dtfmt = strings.NewReplacer(
    "Y", "2006",
    "m", "01",
    "d", "02",
    "H", "15",
    "i", "04",
    "s", "05",
)

func FormatDateTime(format string, t time.Time) string {
    format = dtfmt.Replace(format)
    return t.Format(format)
}

// archiveFiles("e:/amazon/*.xml")
func ArchiveFiles(pattern string) {
    files, err := filepath.Glob(pattern)
    CheckError(err)
    //fmt.Println(files)

    for _, fname := range files {
        dir, file := filepath.Split(fname)
        fileTime, _ := GetFileTime(fname)

        newDir := dir + "archive\\" + fileTime.Format("2006-01-02");
        os.MkdirAll(newDir, 0777)

        newFile := newDir + "\\" + file

        if time.Now().Sub(fileTime) > 15*24*time.Hour {
            //os.Rename(fname, newFile)
            fmt.Println(fname, newFile)
        }
    }
}

func InsertSql(table string, columns []string, data []map[string]string) string {
    columnStr := strings.Join(columns, "`, `")

    updateList := make([]string, len(columns))
    for i, col := range columns {
        updateList[i] = fmt.Sprintf("`%s`=VALUE(`%s`)", col, col)
    }
    updateStr := strings.Join(updateList, ",\n")

    valueList := make([]string, 0)
    for _, row := range data {
        valueRow := make([]string, 0)
        //for _, val := range row { // WRONG
        for _, col := range columns {
            valueRow = append(valueRow, "'" + row[col] + "'")
        }
        valueList = append(valueList, "(" + strings.Join(valueRow, ", ") + ")")
    }
    valueStr := strings.Join(valueList, ",\n")

    //return "INSERT INTO `" + table + "` (" + columnStr + ") VALUES\n" + valueStr + updateStr;
    return fmt.Sprintf("INSERT INTO `%s` (`%s`) VALUES\n%s\nON DUPLICATE KEY UPDATE\n%s",
        table, columnStr, valueStr, updateStr);
}
