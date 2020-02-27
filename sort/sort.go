package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)


var (
	ignoreLetterCase = flag.Bool("f", false, "ignore word's letter case")
	outputOnlyFirst = flag.Bool("u", false, "output only first value among equal")
	inDescendingOrder = flag.Bool("r", false, "sort by descending order")
	outputToFile = flag.String("o", "", "output to file")
	onlyNumbers = flag.Bool("n", false, "sort only numbers")
	byColumn = flag.Int("k", -1, "sort by column")
)

func makeUniqueByColumn(arrFile [][]string) [][]string{
	for i, valI := range arrFile{
		for j := i + 1; j < len(arrFile); j++ {
			if *ignoreLetterCase {
				if strings.ToLower(valI[*byColumn - 1]) == strings.ToLower(arrFile[j][*byColumn -1]) {
					arrFile = append(arrFile[:j], arrFile[(j+1):]...)
					j--
				}
			} else {
				if valI[*byColumn - 1] == arrFile[j][*byColumn -1] {
					arrFile = append(arrFile[:j], arrFile[(j+1):]...)
					j--
				}
			}
		}
	}
	return arrFile
}

func makeUniqueByDefault(arrFile []string) []string {
	for i, valI := range arrFile{
		for j := i + 1; j < len(arrFile); j++ {
			if *ignoreLetterCase {
				if strings.ToLower(valI) == strings.ToLower(arrFile[j]) {
					arrFile = append(arrFile[:j], arrFile[(j+1):]...)
					j--
				}
			} else {
				if valI == arrFile[j] {
					arrFile = append(arrFile[:j], arrFile[(j+1):]...)
					j--
				}
			}
		}
	}
	return arrFile
}

func sortByColumn(file []uint8) []string{
	tmpArr := strings.Split(string(file), "\n")
	/* Создание двумерного слайса */
	arrFile := make([][]string, len(tmpArr))
	for i := range arrFile{
		arrFile[i] = make([]string, len(strings.Split(tmpArr[0], " ")))
		arrFile[i] = append(strings.Split(tmpArr[i], " "))
	}

	if *outputOnlyFirst {
		makeUniqueByColumn(arrFile)
	}

	sort.Slice(arrFile, func(i, j int) bool{
		if *onlyNumbers {
			numValI, _ := strconv.Atoi(arrFile[i][*byColumn -1])
			numValJ, _ := strconv.Atoi(arrFile[j][*byColumn -1])
			if *inDescendingOrder{
				return numValI > numValJ
			} else {
				return numValI < numValJ
			}
		} else {
			if *inDescendingOrder{
				return arrFile[i][*byColumn -1] > arrFile[j][*byColumn -1]
			} else {
				return arrFile[i][*byColumn -1] < arrFile[j][*byColumn -1]
			}
		}
	})
	newArrFile := make([]string, len(arrFile))
	for i, item := range arrFile{
		newArrFile[i] = strings.Join(item, " ")
	}
	return newArrFile
}

func sortByDefault(file []uint8) []string {
	arrFile := strings.Split(string(file), "\n")

	if *outputOnlyFirst {
		arrFile = makeUniqueByDefault(arrFile)
	}

	sort.Slice(arrFile, func(i, j int) bool{
		if *onlyNumbers {
			numValI, _ := strconv.Atoi(arrFile[i])
			numValJ, _ := strconv.Atoi(arrFile[j])
			if *inDescendingOrder{
				return numValI > numValJ
			} else {
				return numValI < numValJ
			}
		}
		if *inDescendingOrder{
			return arrFile[i] > arrFile[j]
		} else {
			return arrFile[i] < arrFile[j]
		}
	})
	return arrFile
}

func sortFile(file []uint8) []string {
	if *byColumn == -1 {
		return sortByDefault(file)
	} else {
		return sortByColumn(file)
	}
}

func outToFile(sortedFile string){
	file, err := os.Create(*outputToFile)
	if err != nil{
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(sortedFile)
}

func getLastArg() string {
	var tmpArg string
	for _, tmpArg = range os.Args[1:] {
	}
	return tmpArg
}

func main(){
	flag.Parse()
	file, err := ioutil.ReadFile(getLastArg())
	if err != nil{
		fmt.Println(err)
	}

	sortedFile := strings.Join(sortFile(file), "\n")
	if *outputToFile != "" {
		outToFile(sortedFile)
	}
	fmt.Print(sortedFile)

}
