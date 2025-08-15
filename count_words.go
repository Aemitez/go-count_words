package main

import (
	"fmt"
	"strings"
)

// CountWords นับจำนวนคำใน string
func CountWords(text string) int {
	// strings.Fields จะตัดช่องว่างและแบ่งเป็น slice ของคำ
	words := strings.Fields(text)
	return len(words)
}

func main() {
	// ตัวอย่างข้อความ
	text := "Hello Golang world! This is an example."

	// นับจำนวนคำ
	count := CountWords(text)

	// แสดงผล
	fmt.Printf("ข้อความ: %q\n", text)
	fmt.Printf("จำนวนคำทั้งหมด: %d\n", count)
}
