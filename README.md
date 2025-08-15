# Word Counter (Go)

โปรแกรมนี้เขียนด้วยภาษา **Golang** สำหรับนับจำนวนคำในข้อความ (`string`)  
ใช้ฟังก์ชัน `strings.Fields` จาก Go Standard Library เพื่อแบ่งข้อความออกเป็น slice ของคำ และนับจำนวนด้วย `len()`  

## 📂 ไฟล์ในโปรเจกต์
- `count_words.go` — ไฟล์ Go หลัก ที่มีฟังก์ชัน `CountWords` และ `main` สำหรับรันทดสอบ

## 🛠 วิธีใช้งาน

### 1. ติดตั้ง Go
ดาวน์โหลดและติดตั้งจาก [https://go.dev/dl/](https://go.dev/dl/)  
ตรวจสอบว่า Go ติดตั้งแล้ว:
```bash
go version

go run count_words.go

ข้อความ: "Hello Golang world! This is an example."
จำนวนคำทั้งหมด: 7
