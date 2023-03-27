package adapter

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"store-bpel/bff/admin_bff/config"
	"store-bpel/staff_service/schema"
)

type IStaffServiceAdapter interface {
	GetStaff(ctx context.Context, staffName, staffId string) ([]*schema.GetStaffResponseData, error)
	GetStaffDetail(ctx context.Context, staffId string) ([]*schema.GetStaffResponseData, error)
	AddStaff(ctx context.Context, request *schema.AddStaffRequest) error
	UpdateStaff(ctx context.Context, staffId string, request *schema.UpdateStaffRequest) error
	CreateAddRequest(ctx context.Context, request *schema.CreateAddRequest) error
	CreateDeleteRequest(ctx context.Context, staffId string) error
	GetRequestList(ctx context.Context) ([]*schema.GetRequestResponseData, error)
	GetStaffAttendance(ctx context.Context, staffId string) ([]*schema.GetStaffAttendanceResponseData, error)
	UpdateRequestStatus(ctx context.Context, requestId string, request *schema.UpdateRequestStatusRequest) error
}

type staffServiceAdapter struct {
	httpClient *http.Client
	port       int
}

func NewStaffAdapter(cfg *config.Config) IStaffServiceAdapter {
	return &staffServiceAdapter{
		httpClient: &http.Client{},
		port:       cfg.StaffServicePort,
	}
}

func (a *staffServiceAdapter) GetStaff(ctx context.Context, staffName, staffId string) ([]*schema.GetStaffResponseData, error) {
	log.Printf("Start to call staff service for GetStaff, filtering staffName %s, staffId %s", staffName, staffId)
	defer log.Println("End call staff service for GetStaff")

	var result *schema.GetStaffResponse

	url := fmt.Sprintf("http://localhost:%d/api/staff-service/staff?name=%s&id=%s", a.port, staffName, staffId)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-GetStaff-NewRequestWithContext error %v", err)
		return nil, err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-GetStaff-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-GetStaff-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-GetStaff-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func (a *staffServiceAdapter) GetStaffDetail(ctx context.Context, staffId string) ([]*schema.GetStaffResponseData, error) {
	log.Printf("Start to call staff service for GetStaffDetail, staffId %s", staffId)
	defer log.Println("End call staff service for GetStaffDetail")

	if staffId == "" {
		return nil, errors.New("[BFF-Adapter-StaffServiceAdapter-GetStaffDetail] staffId must not be empty")
	}

	var result *schema.GetStaffResponse

	url := fmt.Sprintf("http://localhost:%d/api/staff-service/staff/%s", a.port, staffId)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-GetStaffDetail-NewRequestWithContext error %v", err)
		return nil, err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-GetStaffDetail-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-GetStaffDetail-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-GetStaffDetail-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func (a *staffServiceAdapter) AddStaff(ctx context.Context, request *schema.AddStaffRequest) error {
	log.Printf("Start to call staff service for AddStaff")
	defer log.Println("End call staff service for AddStaff")

	var result *schema.UpdateResponse

	url := fmt.Sprintf("http://localhost:%d/api/staff-service/staff", a.port)

	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-AddStaff-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-AddStaff-NewRequestWithContext error %v", err)
		return err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-AddStaff-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-AddStaff-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-AddStaff-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *staffServiceAdapter) UpdateStaff(ctx context.Context, staffId string, request *schema.UpdateStaffRequest) error {
	log.Printf("Start to call staff service for UpdateStaff")
	defer log.Println("End call staff service for UpdateStaff")

	if staffId == "" {
		return errors.New("[BFF-Adapter-StaffServiceAdapter-UpdateStaff] staffId must not be empty")
	}

	var result *schema.UpdateResponse

	url := fmt.Sprintf("http://localhost:%d/api/staff-service/staff/%s", a.port, staffId)

	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-UpdateStaff-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-UpdateStaff-NewRequestWithContext error %v", err)
		return err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-UpdateStaff-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-UpdateStaff-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-UpdateStaff-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *staffServiceAdapter) CreateAddRequest(ctx context.Context, request *schema.CreateAddRequest) error {
	log.Printf("Start to call staff service for CreateAddRequest")
	defer log.Println("End call staff service for CreateAddRequest")

	var result *schema.UpdateResponse

	url := fmt.Sprintf("http://localhost:%d/api/staff-service/request/add", a.port)

	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-CreateAddRequest-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-CreateAddRequest-NewRequestWithContext error %v", err)
		return err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-CreateAddRequest-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-CreateAddRequest-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-CreateAddRequest-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *staffServiceAdapter) CreateDeleteRequest(ctx context.Context, staffId string) error {
	log.Printf("Start to call staff service for CreateDeleteRequest")
	defer log.Println("End call staff service for CreateDeleteRequest")

	if staffId == "" {
		return errors.New("[BFF-Adapter-StaffServiceAdapter-CreateDeleteRequest] staffId must not be empty")
	}

	var result *schema.UpdateResponse

	url := fmt.Sprintf("http://localhost:%d/api/staff-service/request/delete/%s", a.port, staffId)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-CreateDeleteRequest-NewRequestWithContext error %v", err)
		return err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-CreateDeleteRequest-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-CreateDeleteRequest-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-CreateDeleteRequest-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}

func (a *staffServiceAdapter) GetRequestList(ctx context.Context) ([]*schema.GetRequestResponseData, error) {
	log.Printf("Start to call staff service for GetRequestList")
	defer log.Println("End call staff service for GetRequestList")

	var result *schema.GetRequestResponse

	url := fmt.Sprintf("http://localhost:%d/api/staff-service/request", a.port)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-GetRequestList-NewRequestWithContext error %v", err)
		return nil, err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-GetRequestList-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-GetRequestList-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-GetRequestList-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func (a *staffServiceAdapter) GetStaffAttendance(ctx context.Context, staffId string) ([]*schema.GetStaffAttendanceResponseData, error) {
	log.Printf("Start to call staff service for GetStaffAttendance, staffId %s", staffId)
	defer log.Println("End call staff service for GetStaffAttendance")

	if staffId == "" {
		return nil, errors.New("[BFF-Adapter-StaffServiceAdapter-GetStaffAttendance] staffId must not be empty")
	}

	var result *schema.GetStaffAttendanceResponse

	url := fmt.Sprintf("http://localhost:%d/api/staff-service/staff/attendance/%s", a.port, staffId)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-GetStaffAttendance-NewRequestWithContext error %v", err)
		return nil, err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-GetStaffAttendance-httpClient.Do error %v", err)
		return nil, err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-GetStaffAttendance-ioutil.ReadAll error %v", err)
		return nil, err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-GetStaffAttendance-json.Unmarshal error %v", err)
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func (a *staffServiceAdapter) UpdateRequestStatus(ctx context.Context, requestId string, request *schema.UpdateRequestStatusRequest) error {
	log.Printf("Start to call staff service for UpdateRequestStatus")
	defer log.Println("End call staff service for UpdateRequestStatus")

	if requestId == "" {
		return errors.New("[BFF-Adapter-StaffServiceAdapter-UpdateRequestStatus] requestId must not be empty")
	}

	var result *schema.UpdateResponse

	url := fmt.Sprintf("http://localhost:%d/api/staff-service/request/%s", a.port, requestId)

	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-UpdateRequestStatus-Marshal error %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewReader(data))
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-UpdateRequestStatus-NewRequestWithContext error %v", err)
		return err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-UpdateRequestStatus-httpClient.Do error %v", err)
		return err
	}

	respByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-UpdateRequestStatus-ioutil.ReadAll error %v", err)
		return err
	}

	err = json.Unmarshal(respByteArr, &result)
	if err != nil {
		log.Printf("BFF-Adapter-StaffServiceAdapter-UpdateRequestStatus-json.Unmarshal error %v", err)
		return err
	}

	if result.StatusCode != http.StatusOK {
		return errors.New(result.Message)
	}

	return nil
}
