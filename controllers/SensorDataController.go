package controllers

import (
	"io"
	"mime"
	"mime/multipart"
	"model"
	"net/http"
	"proto"
)

func SensorDataJsonHandler(writer http.ResponseWriter, request *http.Request) {
	if ok := readForm(writer, request); !ok {
		return
	}

	uid := request.Form.Get("uid")

	if len(uid) == 0 {
		writeJSON(writer, &proto.SensorDataResponse{Res: 1})
		return
	}

	var data proto.SensorDataList
	err := mime.JSON(request.Body, &data)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		io.WriteString(writer, err.Error())
		return
	}

	sensorDatas := make([]model.SensorData, len(data.Data))

	for i, d := range data.Data {
		sensorDatas[i] = model.SensorData{
			Uid:   uid,
			X:     d.X,
			Y:     d.Y,
			Z:     d.Z,
			Time:  d.Time,
			State: d.State,
		}
	}

	model.SensorDataModel.CreateManay(sensorDatas)

	writeJSON(writer, &proto.SensorDataResponse{Res: 0})
}

func SensorDataFormHandler(writer http.ResponseWriter, request *http.Request) {
	multiReader, err := request.MultipartReader()

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		io.WriteString(writer, err.Error())
		return
	}

	var part *multipart.Part

	for {
		part, err = multiReader.NextPart()

		if err == io.EOF {
			break
		}

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			io.WriteString(writer, err.Error())
			return
		}

		if "sensordata" == part.FormName() {

		}
	}
}
