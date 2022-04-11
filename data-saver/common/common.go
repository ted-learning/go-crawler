package common

import "net/http"

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func HandleServerError(writer http.ResponseWriter, err error) {
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}
