package models

// import (
// 	"../utils"
// 	"bufio"
// 	"fmt"
// 	"io"
// 	. "link/database"
// 	"os"
// )

// type LinkMap struct {
// 	LongUrl string `gorm:"primary_key"`
// 	ShorUrl string
// }

// func init() {
// 	DB.AutoMigrate(&LinkMap{})
// }
// func readValues(infile string) (values []string, err error) {
// 	file, err := os.Open(infile)
// 	if err != nil {
// 		fmt.Println("Failed to open the input file", infile)
// 		return
// 	}
// 	defer file.Close()

// 	br := bufio.NewReader(file)

// 	values = make([]string, 0)
// 	for {
// 		line, isPrefix, err1 := br.ReadLine()
// 		if err1 != nil {
// 			if err1 != io.EOF {
// 				err = err1
// 			}
// 			break
// 		}
// 		if isPrefix {
// 			fmt.Println("A too long line, seems  unexpected")
// 			return
// 		}

// 		str := string(line)
// 		if err1 != nil {
// 			err = err1
// 			return
// 		}
// 		values = append(values, str)

// 	}

// 	return

// }

// func writeValues(values []string, outfile string) error {
// 	file, err := os.Create(outfile)
// 	if err != nil {
// 		fmt.Println("Failed to create write file", outfile)
// 	}
// 	defer file.Close()

// 	for _, value := range values {
// 		file.WriteString("http://127.0.0.1:8080/getLongUrl/" + value + "\n")
// 	}

// 	return nil
// }

// func checkErr(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
// func main() {
// 	values, err := readValues("../static/longurls.txt")
// 	var shortvalues []string
// 	checkErr(err)
// 	for _, value := range values {
// 		var lin LinkMap
// 		shorturls, err := utils.Transform(value)
// 		if err != nil {
// 			fmt.Println(err)
// 			panic(err)
// 		}
// 		lin.LongUrl = value
// 		lin.ShorUrl = shorturls[0]
// 		DB.Create(&lin)
// 		shortvalues = append(shortvalues, shorturls[0])
// 	}

// 	err = writeValues(shortvalues, "../static/shorturls.txt")
// 	checkErr(err)
// }
