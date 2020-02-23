package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestSortFileDefault(t *testing.T) {
	file, err := ioutil.ReadFile("test_data1")
	if err != nil{
		fmt.Println(err)
	}
	expectedStr := "Apple\nBOOK\nBook\nGo\nHauptbahnhof\nJanuary\nJanuary\nNapkin"

	receivedStr := strings.Join(sortFile(file), "\n")
	if receivedStr != expectedStr {
		t.Error("Expected\n", expectedStr, "\ngot\n", receivedStr)
	}

	receivedStr = strings.Join(sortByDefault(file), "\n")
	if receivedStr != expectedStr {
		t.Error("Expected\n", expectedStr, "\ngot\n", receivedStr)
	}
}

func TestSortFileInDescendingOrder(t *testing.T) {
	file, err := ioutil.ReadFile("test_data1")
	if err != nil{
		fmt.Println(err)
	}
	*inDescendingOrder = true
	expectedStr := "Napkin\nJanuary\nJanuary\nHauptbahnhof\nGo\nBook\nBOOK\nApple"
	receivedStr := strings.Join(sortFile(file), "\n")
	if receivedStr != expectedStr {
		t.Error("Expected\n", expectedStr, "\ngot\n", receivedStr)
	}
	*inDescendingOrder = false
}

func TestSortFileOnlyNumbers(t *testing.T) {
	file, err := ioutil.ReadFile("test_data2")
	if err != nil{
		fmt.Println(err)
	}
	*onlyNumbers = true
	expectedStr := "1\n2\n6\n10"
	receivedStr := strings.Join(sortFile(file), "\n")
	if receivedStr != expectedStr {
		t.Error("Expected\n", expectedStr, "\ngot\n", receivedStr)
	}
	*onlyNumbers = false
}

func TestSortFileIgnoreLetterCaseAndOuputOnlyFirst(t *testing.T) {
	file, err := ioutil.ReadFile("test_data1")
	if err != nil{
		fmt.Println(err)
	}
	*outputOnlyFirst = true
	*ignoreLetterCase = true
	expectedStr := "Apple\nBOOK\nGo\nHauptbahnhof\nJanuary\nNapkin"
	receivedStr := strings.Join(sortFile(file), "\n")
	if receivedStr != expectedStr {
		t.Error("Expected\n", expectedStr, "\ngot\n", receivedStr)
	}
	*outputOnlyFirst = false
	*ignoreLetterCase = false
}

func TestSortFileByColumn(t *testing.T) {
	file, err := ioutil.ReadFile("test_data3")
	if err != nil{
		fmt.Println(err)
	}
	*byColumn = 2
	expectedStr := "January aaa\nGo book\nHauptbahnhof fef\nNapkin kek\nBook l\nApple lol\nBOOK mykek\nJanuary table"
	receivedStr := strings.Join(sortFile(file), "\n")
	if receivedStr != expectedStr {
		t.Error("Expected\n", expectedStr, "\ngot\n", receivedStr)
	}
	*byColumn = -1
}

func TestOutFile(t *testing.T) {
	file, err := ioutil.ReadFile("test_data3")
	if err != nil{
		fmt.Println(err)
	}
	*outputToFile = "tmp_file"
	expectedStr := strings.Join(sortFile(file), "\n")
	outToFile(expectedStr)

	fileR, err := ioutil.ReadFile(*outputToFile)
	if expectedStr != string(fileR) {
		t.Error("Expected\n", expectedStr, "\ngot\n", string(fileR))
	}
	*outputToFile = ""
}

func TestMakeUniqueByDefault(t *testing.T){
	originArr := []string {"Hauptbahnhof\n", "January\n","January\n", "Napkin\n", "Napkin"}
	receivedArr := makeUniqueByDefault(originArr)
	expectedArr := []string {"Hauptbahnhof\n", "January\n", "Napkin\n"}

	for i, item := range expectedArr{
		if receivedArr[i] != item {
			t.Error("Expected\n", item, "\ngot\n", receivedArr[i])
		}
	}
}

func TestMakeUniqueByColumn(t *testing.T){
	*ignoreLetterCase = true
	*byColumn = 1
	file, err := ioutil.ReadFile("test_data3")
	if err != nil{
		fmt.Println(err)
	}
	tmpArr := strings.Split(string(file), "\n")
	arrFile := make([][]string, len(tmpArr))
	for i := range arrFile{
		arrFile[i] = make([]string, len(strings.Split(tmpArr[0], " ")))
		arrFile[i] = append(strings.Split(tmpArr[i], " "))
	}
	makeUniqueByColumn(arrFile)

	receivedArr := make([]string, len(arrFile))
	for i, item := range arrFile{
		receivedArr[i] = strings.Join(item, " ")
	}
	expectedArr := []string {"Napkin kek","Apple lol","January table","BOOK mykek","Hauptbahnhof fef","Go book"}
	for i, item := range expectedArr{
		if receivedArr[i] != item {
			t.Error("Expected\n", item, "\ngot\n", receivedArr[i])
		}
	}
	*ignoreLetterCase = false
	*byColumn = -1
}


