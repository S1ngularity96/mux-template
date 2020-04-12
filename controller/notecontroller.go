package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"mux-template/auxiliary"
	"mux-template/models"
	"net/http"
	"strconv"
)

func GetNotes(w http.ResponseWriter, r *http.Request){
	notes, err := models.GetNotes()

	if err != nil {
		auxiliary.InternalErrorResponse(w,r,err)
	}

	data, err := json.Marshal(notes)

	if err != nil {
		auxiliary.InternalErrorResponse(w,r,err)
	}
	auxiliary.OkResponse(w,r,string(data))
}

func PostNote(w http.ResponseWriter, r *http.Request){
	var note models.Note
	note.Message = r.FormValue("message")
	if err :=note.Create(); err != nil {
		auxiliary.InternalErrorResponse(w,r,err)
		return
	}
	auxiliary.OkResponse(w,r,"OK")
}

func GetNoteById(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id := params["id"]
	INTid, err := strconv.Atoi(id)

	if err != nil {
		auxiliary.InternalErrorResponse(w,r,err)
		return
	}
	var note models.Note
	note, err = models.GetNoteByID(uint(INTid))
	if err != nil {
		auxiliary.InternalErrorResponse(w,r,err)
		return
	}
	jsonData, err := json.Marshal(&note)
	if err != nil {
		auxiliary.InternalErrorResponse(w,r,err)
		return
	}
	auxiliary.OkResponse(w,r,string(jsonData))
}

func DeleteNoteById(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id := params["id"]
	INTid, err := strconv.Atoi(id)

	if err != nil {
		auxiliary.InternalErrorResponse(w,r,err)
		return
	}
	var note models.Note
	models.DbCon.First(&note,INTid)

	if note.ID == 0{
		auxiliary.NotFoundResponse(w,r,"Note with ID does not exist")
	}
	models.DbCon.Delete(&note)
	auxiliary.OkResponse(w,r,"OK")
}