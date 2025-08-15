# Write an API with Gin that has HTTP GET /ping and returns "pong". (Go)

โปรแกรมนี้เขียนด้วยภาษา **Golang** สำหรับนับจำนวนคำในข้อความ (`string`)  
ใช้ฟังก์ชัน `strings.Fields` จาก Go Standard Library เพื่อแบ่งข้อความออกเป็น slice ของคำ และนับจำนวนด้วย `len()`  

## 📂 ไฟล์ในโปรเจกต์
- `go.mod` — ใช้สำหรับจัดการ dependencies
- `go.sum` — ใช้สำหรับตรวจสอบความถูกต้องของโมดูล (modules) ที่ดาวน์โหลดมาใช้ในโปรเจกต์
- `main.go` — ไฟล์ Go หลัก ที่มีฟังก์ชัน `gin` และ `main` สำหรับรันทดสอบ

## 🛠 วิธีใช้งาน

### 1. ติดตั้ง Go
ดาวน์โหลดและติดตั้งจาก [https://go.dev/dl/](https://go.dev/dl/)  
ตรวจสอบว่า Go ติดตั้งแล้ว:
```bash
go version

go mod init go-ping-pong

go get -u github.com/gin-gonic/gin

go run main.go
```
### 1. TEST 
#### [GET] http://localhost:8080/ping
[Link For Test Call](http://localhost:8080/ping "/ping")

 - Response 
```
   "message": "pong"
```
