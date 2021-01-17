package main

import "log"

func main() {
	batchList := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	batchCnt := 3
	var insertData [][]string
	var data []string
	idx := 0
	for range batchList {
		for j := 0; j < batchCnt; j += 1 {

			if idx < len(batchList) {
				data = append(data, batchList[idx])
			}
			idx++
		}

		if len(data) > 0 {
			insertData = append(insertData, data)
		}
		data = []string{}
	}

	log.Println(insertData)
}
