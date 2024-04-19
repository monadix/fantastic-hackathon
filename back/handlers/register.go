package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type RegisterRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	SecondName  string `json:"second_name"`
	LastName    string `json:"last_name"`
	PicturePath string `json:"picture_path"`
}

type RegisterResponse struct {
}

func (s *HandlersServer) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var resp RegisterRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln("Enable to read data: ", err)
		return
	} else {
		err = json.Unmarshal(body, &resp)
		if err != nil {
			log.Fatalln("Enable to convert data to json: ", err)
			return
		}
	}
	var cnt int64
	s.DB.Where("google_id = ?", resp.ID).Count(&cnt)
	if cnt == 0 {
		s.DB.Table("users").Create(&User{
			ID:           resp.ID,
			FirstName:    resp.FirstName,
			SecondName:   &resp.SecondName,
			LastName:     resp.LastName,
			PasswordHash: resp.PasswordHash,
			PhoneNumber:  resp.PhoneNumber,
			PicturePath:  &resp.PicturePath,
		})
	} else {
		log.Fatalln("User with current phone number already exist.")
		return
	}
}
