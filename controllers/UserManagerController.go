package controllers

import (
	"io"
	"net/http"

	"github.com/zjwdmlmx/ssensor/mime"
	"github.com/zjwdmlmx/ssensor/model"
	"github.com/zjwdmlmx/ssensor/proto"
)

func RegistUserHandler(writer http.ResponseWriter, request *http.Request) {
	var data proto.UserRegistry

	if err := mime.JSON(request.Body, &data); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		io.WriteString(writer, err.Error())
		return
	}

	if err := model.UserModel.CreateOne(&model.User{Uid: data.Uid}); err == model.ErrUserExists {
		writeJSON(writer, &proto.UserRegistryResponse{Res: 1})
	} else {
		writeJSON(writer, &proto.UserRegistryResponse{Res: 0})
	}
}
