package main

import (
	"net/http"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type RequestModel struct {
	OperatorID string `binding:"required"`
	Data string `binding:"required"`
	TimeStamp string `binding:"required"`
	Seq string `binding:"required"`
	Sig string `binding:"required"`
}

type QueryStationStatsBo struct {
	StationID string
	StartTime string
	EndTime string
}

type ResponseModel struct {
	Ret int
	Msg string
	Data string
	Sig string
}

// type SuccessData struct {
// 	StationStats StationStats `json:"StationStats"`
// }
// type StationStats struct {
// 	StationID string `json:"StationID"`
// 	StartTime string `json:"StartTime"`
// 	EndTime string `json:"EndTime"`
// 	StationElectricity float64 `json:"StationElectricity"`
// 	EquipmentStatsInfos []EquipmentStatsInfo `json:"EquipmentStatsInfos"`
// }
// type EquipmentStatsInfo struct {
// 	EquipmentID string `json:"EquipmentID"`
// 	EquipmentElectricity string `json:"EquipmentElectricity"`
// 	ConnectorStatsInfos []ConnectorStatsInfo `json:"ConnectorStatsInfos"`
// }
// type ConnectorStatsInfo struct {
// 	ConnectorID string `json:"ConnectorID"`
// 	ConnectorElectricity float64 `json:"ConnectorElectricity"`
// }

type SuccessData struct {
	StationStats StationStats
}
type StationStats struct {
	StationID string
	StartTime string
	EndTime string
	StationElectricity float64
	EquipmentStatsInfos []EquipmentStatsInfo
}
type EquipmentStatsInfo struct {
	EquipmentID string
	EquipmentElectricity string
	ConnectorStatsInfos []ConnectorStatsInfo
}
type ConnectorStatsInfo struct {
	ConnectorID string
	ConnectorElectricity float64
}

func main() {
	router := gin.Default()

	router.POST("/query_station_stats", func(ctx *gin.Context) {
		var req = RequestModel{}
		if err := ctx.ShouldBind(&req); err != nil {
			errorReq := ResponseModel{
				http.StatusInternalServerError,
				"系统错误",
				"",
				"123456789",
			}
			ctx.JSON(http.StatusInternalServerError, errorReq)
			return
		}
		successData := SuccessData{
			StationStats: StationStats{
				StationID: "StationID_57dce96ddcc7",
				StartTime: "StartTime_2cb1d439bf85",
				EndTime: "EndTime_0a2ea9ef4fce",
				StationElectricity: 0.00,
				EquipmentStatsInfos: []EquipmentStatsInfo{
					{
						EquipmentID: "EquipmentID_17e274204245",
						EquipmentElectricity: "EquipmentElectricity_f21cff1927af",
						ConnectorStatsInfos: []ConnectorStatsInfo{
							{
								ConnectorID: "ConnectorID_f217f74367dd",
								ConnectorElectricity: 0.00,
							},
						},
					},
				},
			},
		}
		successDataJSON, _ := json.Marshal(successData)
		successResp := ResponseModel{
			0,
			"success",
			string(successDataJSON),
			"123456",
		}
		ctx.JSON(http.StatusOK, successResp)
	})
	router.Run(":1234")
}						
				
