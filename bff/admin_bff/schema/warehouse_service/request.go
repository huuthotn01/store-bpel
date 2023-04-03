package warehouse_service

type GetWarehouseId struct {
	WarehouseId string
}

type AddWarehouseStaffRequest struct {
	StaffId     string
	WarehouseId string
	Role        string
}

type UpdateWarehouseStaffRequest struct {
	StaffId     string
	WarehouseId string
	Role        string
}

type DeleteWarehouseStaffRequest struct {
	StaffId string
}

type UpdateManagerRequest struct {
	StaffId     string
	WarehouseId string
}

type AddWarehouseRequest struct {
	WarehouseName string
	Capacity      int
	Street        string
	Ward          string
	District      string
	Province      string
}

type UpdateWarehouseRequest struct {
	WarehouseCode string
	WarehouseName string
	Capacity      int
	Street        string
	Ward          string
	District      string
	Province      string
}

type DeleteWarehouseRequest struct {
	WarehouseCode string
}
