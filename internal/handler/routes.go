package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"test/internal/service"
	"test/pkg"
)

type Handlers struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handlers {
	return &Handlers{service: s}
}

func (h *Handlers) Init() *mux.Router {
	router := mux.NewRouter()
	router.Use(CORS, RecoverAllPanic)
	router.HandleFunc("/users", h.GetUsers).Methods(http.MethodGet, http.MethodOptions)
	//router.HandleFunc("/user-by-id", h.GetUserById).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/drop", h.UserDelete).Methods("DELETE", "OPTION")
	//router.HandleFunc("/create", CreateUser).Methods(http.MethodPost, http.MethodOptions)
	return router
}

func (h *Handlers) UserDelete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("user_id")
	_, err := strconv.Atoi(id)
	if err != nil {
		pkg.ErrorResponse(w, 400, "только цифры")
		return
	}
	err = h.service.DeleteUser(id)
	if err != nil {
		log.Println(err)
		pkg.ErrorResponse(w, 500, "sorry")
		return
	}
	pkg.Response(w, "success")
}

func (h *Handlers) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetUsers()
	if err != nil {
		log.Println(err)
		pkg.ErrorResponse(w, 500, "sorry")
		return
	}
	data, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
		w.Write([]byte("на сервере возникла ошибка"))
		return
	}
	w.Write(data)
}

//func GetUserById(w http.ResponseWriter, r *http.Request) {
//	idStr := r.URL.Query().Get("user_id")
//
//	connect := NewConnectPostgres()
//	var user UsersT
//	connect.Table("users_t").Where("id", idStr).First(&user)
//	if connect.Error != nil {
//		log.Println(connect.Error)
//		w.Write([]byte("на сервере возникла ошибка"))
//		return
//	}
//	data, err := json.Marshal(user)
//	if err != nil {
//		log.Println(err)
//		w.Write([]byte("на сервере возникла ошибка"))
//		return
//	}
//
//	w.Write(data)
//
//}
//
//func ChangeUserState(w http.ResponseWriter, r *http.Request) {
//	idStr := r.URL.Query().Get("user_id")
//
//	connect := NewConnectPostgres()
//	connect.Table("users_t").Where("id", idStr).UpdateColumn("active", false)
//	if connect.Error != nil {
//		log.Println(connect.Error)
//		w.Write([]byte("на сервере возникла ошибка"))
//		return
//	}
//	w.Write([]byte("Пользователь успешно деактивирован"))
//
//}
//
//func CreateUser(w http.ResponseWriter, r *http.Request) {
//	var userParams UsersT
//	err := json.NewDecoder(r.Body).Decode(&userParams)
//	if err != nil {
//		log.Println(err)
//		pkg.ErrorResponse(w, 400, "некорректные данные клиента")
//		return
//	}
//	if userParams.CiteId == 0 {
//		pkg.ErrorResponse(w, 400, "выберите город")
//		return
//	}
//	if userParams.RoleId == 0 {
//		pkg.ErrorResponse(w, 400, "укажите роль клиента")
//		return
//	}
//
//	conn := NewConnectPostgres()
//
//	err = conn.Table("users_t").Create(&userParams).Error
//	if err != nil {
//		log.Println(err)
//		pkg.ErrorResponse(w, 500, err.Error())
//		return
//	}
//	pkg.Response(w, userParams)
//}
