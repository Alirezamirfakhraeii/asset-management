package model

type Asset struct {
	AssetID      string `json:"asset_id"`
	EmployeeID   string `json:"employee_id"`
	AssetType    string `json:"asset_type"`
	UserFullName string `json:"user_full_name"`
	Location     string `json:"location"`
	Material     string `json:"material"`
	SerialNumber string `json:"serial_number"`
	Color        string `json:"color"`
	Description  string `json:"description"`
}
