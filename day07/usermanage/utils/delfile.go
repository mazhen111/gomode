package utils

import "os"

func DelFile() {
	fileInos := GetSortedFileList()
	if len(fileInos) > 5 {
		for _, file := range fileInos[:len(fileInos)-5] {
			os.Remove("data/" + file.Name())

		}
	}

}
