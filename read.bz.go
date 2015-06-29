package main

import (
	"bufio"
	"compress/bzip2"
	"fmt"
	"os"
	//	"strconv"
	//	"strings"
)

func main() {
	//	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	//		advance, token, err = bufio.ScanWords(data, atEOF)
	//		if err == nil && token != nil {
	//			_, err = strconv.ParseInt(string(token), 10, 32)
	//		}
	//		return
	//	}
	f, _ := os.Open("test.data.bz2")
	//fbz2:= strings.NewReader(f)
	bzfile := bzip2.NewReader(f)
	scanner := bufio.NewScanner(bzfile)
	// Set the split function for the scanning operation.
	//scanner.Split(split)
	// Validate the input
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text()) //seems to read 1 line at a time as a big string
		//fmt.Println("adfadfadfasd")
	}
}
