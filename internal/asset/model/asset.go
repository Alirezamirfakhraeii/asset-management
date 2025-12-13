package model

import "time"

// ---> define Models + JSON Tags

type CategoryStatus string

const (
	Pending  CategoryStatus = "pending"
	Approved CategoryStatus = "approved"
	Rejected CategoryStatus = "rejected"
)

// User Model

type Users struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" gorm:"unique"`
	Password  string    `json:"password"`
	FullName  string    `json:"fullName"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Asset Model
type Asset struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	EmployeeCode string         `json:"employee_code"`
	AssetType    string         `json:"asset_type"`
	UserFullName string         `json:"user_full_name"`
	Location     string         `json:"location,omitempty"`
	Material     string         `json:"material,omitempty"`
	SerialNumber string         `json:"serial_number,omitempty"`
	Color        string         `json:"color,omitempty"`
	Description  string         `json:"description,omitempty"`
	CategoryId   string         `json:"category_id"`
	Status       CategoryStatus `json:"status"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

// Category Model
type Category struct {
	ID          uint      `json:"id" gorm:"column:id"`             // اتصال به ستون category_id در دیتابیس
	Name        string    `json:"name" gorm:"column:name"`         // اتصال به ستون category_name در دیتابیس
	Description string    `json:"description,omitempty,omitempty"` // اتصال به ستون description در دیتابیس
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// SubCategory Model

type SubCategory struct {
	SubCategoryId   string    `json:"id"`
	SubCategoryName string    `json:"name"`
	Description     string    `json:"description,omitempty"` // اضافه کردن توصیف برای زیر دسته‌بندی
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
