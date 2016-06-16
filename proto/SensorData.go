package proto

const (
	STATE_RUN   uint8 = 1
	STATE_WALK  uint8 = 2
	STATE_STAND uint8 = 3
	STATE_JUMP  uint8 = 4
)

type SensorData struct {
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Z     float64 `json:"z"`
	Time  int64   `json:"time"`
	State uint8   `json:"state"`
}

type SensorDataList struct {
	Data []SensorData `json:"data"`
}

type SensorDataResponse struct {
	Res int `json:"res"`
}
