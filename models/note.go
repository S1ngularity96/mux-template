package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Note struct {
	gorm.Model
	Message string `json:"message"`
}

func  (note Note) Create() error{

	if note.Message == ""{
		return fmt.Errorf("Die Notitz enthält keinen Text")
	}
	if DbCon != nil{
		DbCon.Create(&note)
	}
	return nil
}

func GetNotes() ([]Note, error) {
	var notes []Note
	if err := DbCon.Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}

func GetNoteByID(ID uint) (Note, error){
	var note Note
	if err := DbCon.Where("ID = ?", ID).First(&note).Error; err != nil {
		return note,err
	}
	return note, nil
}
