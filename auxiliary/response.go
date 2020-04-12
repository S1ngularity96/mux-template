package auxiliary

import "net/http"

func OkResponse(w http.ResponseWriter, r *http.Request, data string){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, data string){
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(data))
}

func UnauthorizedResponse(w http.ResponseWriter, r *http.Request, data string){
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(data))
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request, data string){
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(data))
}

func ForbiddenResponse(w http.ResponseWriter, r *http.Request, data string){
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte(data))

}

func InternalErrorResponse(w http.ResponseWriter, r *http.Request, data error){
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(data.Error()))
}


