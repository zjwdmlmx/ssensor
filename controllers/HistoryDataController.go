package controllers

import (
	"io"
	"net/http"
	"time"

	"github.com/zjwdmlmx/ssensor/mime"
	"github.com/zjwdmlmx/ssensor/model"
	"github.com/zjwdmlmx/ssensor/proto"
)

func HistoryJsonHandler(writer http.ResponseWriter, request *http.Request) {
	if ok := readForm(writer, request); !ok {
		return
	}

	uid := request.Form.Get("uid")

	if len(uid) == 0 {
		writeJSON(writer, &proto.SensorDataResponse{Res: 1})
		return
	}

	var data proto.HistoryDataList

	err := mime.JSON(request.Body, &data)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		io.WriteString(writer, err.Error())
		return
	}

	historyDatas := make([]model.HistoryData, len(data.Data))

	for i, d := range data.Data {
		historyDatas[i] = model.HistoryData{
			Uid:       uid,
			Time:      time.Unix(d.Time, 0),
			State:     d.State,
			Latitude:  d.Latitude,
			Longitude: d.Longitude,
		}
	}

	model.HistoryDataModel.CreateManay(historyDatas)

	writeJSON(writer, &proto.HistoryDataResponse{Res: 0})
}
