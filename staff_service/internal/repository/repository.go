package repository

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type IStaffServiceRepository interface {
	GetStaff(ctx context.Context, staffName, staffId string) ([]*StaffModel, error)
	AddStaff(ctx context.Context, staff *StaffModel) error
	GetStaffDetail(ctx context.Context, staffId string) (*StaffModel, error)
	UpdateStaff(ctx context.Context, data *StaffModel) error
	DeleteStaffUpdateStatus(ctx context.Context, staffId string) error
	DeleteStaffRemove(ctx context.Context, staffId string) error
	CreateAccount(ctx context.Context, data *AccountModel) error
	GetStaffAttendance(ctx context.Context, staffId string) ([]*AttendanceModel, error)
	CreateStaffRequest(ctx context.Context, request *RequestsModel) error
	UpdateRequestStatus(ctx context.Context, status, requestId string) error
	GetStaffRequest(ctx context.Context, requestId string) (*RequestsModel, error)
	GetListRequest(ctx context.Context) ([]*GetRequestResponseData, error)
}

type GetRequestResponseData struct {
	Id            string
	RequestDate   time.Time
	RequestType   string // ADD or DELETE
	Status        string
	StaffId       string
	StaffName     string
	Province      string
	District      string
	Ward          string
	Street        string
	Hometown      string
	CitizenId     string
	StaffPosition string
	Birthdate     string
	StartDate     time.Time
	Salary        int
	Gender        string
	Phone         string
	Email         string
	BranchId      string
}

func NewRepository(db *gorm.DB) IStaffServiceRepository {
	return &staffServiceRepository{
		db:                  db,
		accountTableName:    "account",
		staffTableName:      "staff",
		attendanceTableName: "attendance",
		requestsTaleName:    "requests",
	}
}

func (r *staffServiceRepository) GetStaff(ctx context.Context, staffName, staffId string) ([]*StaffModel, error) {
	var result []*StaffModel
	query := r.db.WithContext(ctx).Where("status = 'APPROVED'").Table(r.staffTableName)
	if staffName != "" && staffId != "" {
		query = query.Where("staff_name LIKE ? OR staff_id LIKE ?", "%"+staffName+"%", "%"+staffId+"%")
	} else if staffName != "" {
		query = query.Where("staff_name LIKE ?", "%"+staffName+"%")
	} else if staffId != "" {
		query = query.Where("staff_id LIKE ?", "%"+staffId+"%")
	}
	return result, query.Find(&result).Error
}

func (r *staffServiceRepository) GetStaffDetail(ctx context.Context, staffId string) (*StaffModel, error) {
	var result *StaffModel
	query := r.db.WithContext(ctx).Table(r.staffTableName).Where("staff_id = ?", staffId).First(&result)
	return result, query.Error
}

func (r *staffServiceRepository) AddStaff(ctx context.Context, staff *StaffModel) error {
	// add to staff table
	return r.db.WithContext(ctx).Table(r.staffTableName).Select(`staff_id`, `staff_name`, `province`, `district`, `ward`, `street`, `hometown`, `citizen_id`, `staff_position`, `birthdate`, `salary`,
		`gender`, `phone`, `email`, `status`, `branch_id`).Create(staff).Error
}

func (r *staffServiceRepository) UpdateStaff(ctx context.Context, data *StaffModel) error {
	return r.db.WithContext(ctx).Table(r.staffTableName).Where("staff_id = ?", data.StaffId).Updates(data).Error
}

func (r *staffServiceRepository) DeleteStaffUpdateStatus(ctx context.Context, staffId string) error {
	// update status to 'DELETED' in staff table
	return r.db.WithContext(ctx).Table(r.staffTableName).Where("staff_id = ?", staffId).Update("status", "DELETED").Error
}

func (r *staffServiceRepository) DeleteStaffRemove(ctx context.Context, staffId string) error {
	// update status to 'DELETED' in staff table
	return r.db.WithContext(ctx).Table(r.staffTableName).Where("staff_id = ?", staffId).Delete(staffId).Error
}

func (r *staffServiceRepository) CreateAccount(ctx context.Context, data *AccountModel) error {
	return r.db.WithContext(ctx).Table(r.accountTableName).Create(data).Error
}
func (r *staffServiceRepository) GetStaffAttendance(ctx context.Context, staffId string) ([]*AttendanceModel, error) {
	result := make([]*AttendanceModel, 0)
	query := r.db.WithContext(ctx).Table(r.attendanceTableName).Where("staff_id = ?", staffId).Find(&result)
	return result, query.Error
}

func (r *staffServiceRepository) CreateStaffRequest(ctx context.Context, request *RequestsModel) error {
	return r.db.WithContext(ctx).Table(r.requestsTaleName).Create(request).Error
}

func (r *staffServiceRepository) UpdateRequestStatus(ctx context.Context, status, requestId string) error {
	return r.db.WithContext(ctx).Table(r.requestsTaleName).Where("id = ?", requestId).Update("status", status).Error
}

func (r *staffServiceRepository) GetStaffRequest(ctx context.Context, requestId string) (*RequestsModel, error) {
	var result *RequestsModel
	query := r.db.WithContext(ctx).Table(r.requestsTaleName).Where("id = ?", requestId).First(&result)
	return result, query.Error
}

func (r *staffServiceRepository) GetListRequest(ctx context.Context) ([]*GetRequestResponseData, error) {
	var result []*GetRequestResponseData
	query := r.db.WithContext(ctx).Table(r.requestsTaleName).Select("`requests`.`id`,`requests`.`request_date`," +
		"`requests`.`request_type`,`requests`.`status`,`requests`.`staff_id`,`staff`.`staff_name`,`staff`.`province`," +
		"`staff`.`district`,`staff`.`ward`,`staff`.`street`,`staff`.`hometown`,`staff`.`citizen_id`," +
		"`staff`.`staff_position`,`staff`.`birthdate`,`staff`.`start_date`,`staff`.`salary`,`staff`.`gender`," +
		"`staff`.`phone`,`staff`.`email`,`staff`.`branch_id`").Joins("JOIN `staff` on requests.staff_id = staff.staff_id").
		Where("requests.status = 'PENDING'").Find(&result)
	return result, query.Error
}
