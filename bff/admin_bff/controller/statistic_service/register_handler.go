package statistic_service

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"store-bpel/bff/admin_bff/config"
	"store-bpel/bff/admin_bff/schema/statistic_service"
)

var statisticController IStatisticBffController

func RegisterEndpointHandler(mux *mux.Router, cfg *config.Config) {
	// init controller
	statisticController = NewController(cfg)
	// register handler
	mux.HandleFunc("/api/bff/statistic-service/overall-stat", handleGetOverallStat)
	mux.HandleFunc("/api/bff/statistic-service/revenue", handleGetRevenue)
	mux.HandleFunc("/api/bff/statistic-service/revenue-goods", handleGetRevenueOneGoods)
	mux.HandleFunc("/api/bff/statistic-service/profit", handleGetProfit)
	mux.HandleFunc("/api/bff/statistic-service/profit-goods", handleGetProfitOneGoods)
}

func handleGetOverallStat(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&statistic_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Statistic-handleGetOverallStat-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(statistic_service.GetOverallStatRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&statistic_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Statistic-handleGetOverallStat-xml.Unmarshal err %v", err),
			})
		}
		stat, err := statisticController.GetOverallStat(ctx, request)
		if err != nil {
			err = enc.Encode(&statistic_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Statistic-handleGetOverallStat-GetOverallStat err %v", err),
			})
		} else {
			err = enc.Encode(&statistic_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       stat,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleGetRevenue(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&statistic_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Statistic-handleGetRevenue-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(statistic_service.FilterGetStatisticRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&statistic_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Statistic-handleGetRevenue-xml.Unmarshal err %v", err),
			})
		}
		revenue, err := statisticController.GetRevenue(ctx, request)
		if err != nil {
			err = enc.Encode(&statistic_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Statistic-handleGetRevenue-GetOverallStat err %v", err),
			})
		} else {
			err = enc.Encode(&statistic_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       revenue,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleGetRevenueOneGoods(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&statistic_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Statistic-handleGetRevenueOneGoods-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(statistic_service.GetStatOneGoodsRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&statistic_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Statistic-handleGetRevenueOneGoods-xml.Unmarshal err %v", err),
			})
		}
		revenue, err := statisticController.GetRevenueOneGoods(ctx, request)
		if err != nil {
			err = enc.Encode(&statistic_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Statistic-handleGetRevenueOneGoods-GetOverallStat err %v", err),
			})
		} else {
			err = enc.Encode(&statistic_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       revenue,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleGetProfit(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&statistic_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Statistic-handleGetProfit-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(statistic_service.FilterGetStatisticRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&statistic_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Statistic-handleGetProfit-xml.Unmarshal err %v", err),
			})
		}
		profit, err := statisticController.GetProfit(ctx, request)
		if err != nil {
			err = enc.Encode(&statistic_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Statistic-handleGetProfit-GetOverallStat err %v", err),
			})
		} else {
			err = enc.Encode(&statistic_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       profit,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}

func handleGetProfitOneGoods(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = enc.Encode(&statistic_service.GetResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("BFF-Statistic-handleGetProfitOneGoods-ioutil.ReadAll err %v", err),
		})
		return
	}
	if r.Method == http.MethodPost {
		var request = new(statistic_service.GetStatOneGoodsRequest)
		err = xml.Unmarshal(payload, request)
		if err != nil {
			err = enc.Encode(&statistic_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Statistic-handleGetProfitOneGoods-xml.Unmarshal err %v", err),
			})
		}
		profit, err := statisticController.GetProfitOneGoods(ctx, request)
		if err != nil {
			err = enc.Encode(&statistic_service.GetResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("BFF-Statistic-handleGetProfitOneGoods-GetOverallStat err %v", err),
			})
		} else {
			err = enc.Encode(&statistic_service.GetResponse{
				StatusCode: 200,
				Message:    "OK",
				Data:       profit,
			})
		}
	} else {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
}
