## Overview

โปรเจกต์ **be-go-music-store-1** เป็นแอปพลิเคชันสำหรับการจัดการสินค้าที่ขายในร้านขายเครื่องดนตรี (Music Store) โดยสามารถทำการ **Create**, **Update**, **Remove**, และ **Read** รายการสินค้าได้

## Features

โปรเจกต์นี้มีฟังก์ชันหลักดังนี้:

- **Create**: เพิ่มสินค้าใหม่เข้าสู่ระบบ
- **Update**: อัปเดตข้อมูลของสินค้าที่มีอยู่
- **Remove**: ลบสินค้าที่ไม่ต้องการออกจากระบบ
- **Read**: ดึงข้อมูลสินค้าทั้งหมดหรือสินค้าตาม ID

## Tech Stack

โปรเจกต์นี้ใช้เทคโนโลยีดังนี้:

- **Go**: ภาษาโปรแกรมหลักที่ใช้ในการพัฒนา
- **Gin**: Web framework สำหรับ Go
- **GORM**: ORM (Object-Relational Mapping) สำหรับการจัดการฐานข้อมูล
- **PostgreSQL**: ฐานข้อมูลที่ใช้เก็บข้อมูลสินค้า

## How to Run the Project Locally

### 1. Clone the repository

````bash
git clone https://github.com/your-username/be-go-music-store-1.git
cd be-go-music-store-1
````

### 2. Install Dependencies
```bash
go mod tidy
````

### 3. Set up Environment Variables
```bash
PORT=8080
DB_HOST=
DB_USER=postgres
DB_PASSWORD=
DB_NAME=music_store
DB_PORT=5432
````

### 4. Run the Application
```bash
go run main.go
````

### 5. Testing the Endpoints
คุณสามารถทดสอบ API ของโปรเจกต์ได้โดยการใช้ Postman หรือ cURL เพื่อเรียกดู endpoints ต่าง ๆ:

```bash
GET /products - ดึงรายการสินค้าทั้งหมด

POST /admin/products - เพิ่มสินค้าใหม่

PUT /admin/products/:id - อัปเดตข้อมูลสินค้าตาม ID

DELETE /admin/products/:id - ลบสินค้าตาม ID
````


