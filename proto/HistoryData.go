package proto

type HistoryData struct {
	Time      int64   `json:"time"`
	State     uint8   `json:"state"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type HistoryDataList struct {
	Data []HistoryData `json:"data"`
}

type HistoryDataResponse struct {
	Res int `json:"res"`
}
