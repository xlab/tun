package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func errorf(format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, format, v...)
}

func errorln(v ...interface{}) {
	fmt.Fprintln(os.Stderr, v...)
}

func fatalf(format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, format, v...)
	os.Exit(1)
}

func fatalln(v ...interface{}) {
	fmt.Fprintln(os.Stderr, v...)
	os.Exit(1)
}

func getAlphanumericPin(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	a := getAlphabet()
	buf := make([]byte, length)
	for i := 0; i < length; i++ {
		buf[i] = a[r.Intn(len(a))]
	}
	return string(buf)
}

func getAlphabet() []byte {
	len1, len2 := int('9'-'0')+1, int('z'-'a')+1
	chars := make([]byte, len1+len2)
	for i := 0; i < len1; i++ {
		chars[i] = byte('0' + i)
	}
	for i := 0; i < len2; i++ {
		chars[len1+i] = byte('a' + i)
	}
	return chars
}
