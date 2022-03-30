package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"strings"
)

func main() {
	s := "sha1 this string"

	h := sha1.New()
	m := md5.New()

	h.Write([]byte(s))
	m.Write([]byte("123"))

	bs := h.Sum(nil)
	m5 := m.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)

	newArr := fmt.Sprintf("%x\n", m5)
	sig := strings.ToTitle(newArr)
	fmt.Println(sig)
}
