package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"wedding/models"
	"wedding/service"

	"github.com/gorilla/mux"
)

type AllController struct {
	service *service.AllService
}
func NewAllController(service *service.AllService) *AllController {
	return &AllController{service: service}
}

func (c *AllController) GetAll(w http.ResponseWriter, r *http.Request) {
    products, err := c.service.GetAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(products)
}

    func (c *AllController) AddData(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
            return
        }

        var user models.User
        if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
            http.Error(w, "Failed to decode JSON: "+err.Error(), http.StatusBadRequest)
            return
        }

        fmt.Println("id :", user.ID)
        fmt.Println("Nama :", user.Name)
        
        endCodename := url.QueryEscape(user.Name)
        invitationLink := fmt.Sprintf("https://married-templte1.vercel.app/undangan/%s", endCodename)


        
        user.Link = invitationLink
    
        // Menyimpan data user terlebih dahulu ke database
        if err := c.service.AddData(&user); err != nil {
            http.Error(w, "Failed to add data: "+err.Error(), http.StatusBadRequest)
            return
        }


        fmt.Println("id :", user.ID)
        // Setelah data berhasil disimpan, buat link undangan dengan ID yang sudah terisi
        

        // Kirim response sukses
        w.WriteHeader(http.StatusCreated)
        response := map[string]interface{}{
            "message":        "Produk berhasil ditambahkan!",
            "invitationLink": invitationLink,
            "namaUser": user.Name,
            "user":           user,
        }
        json.NewEncoder(w).Encode(response)
    }



// Mengambil ID dari URL menggunakan mux.Vars
func (c *AllController) GetInvitationLink(w http.ResponseWriter, r *http.Request) {
    name := mux.Vars(r)["name"] // Ambil nama dari URL

    // Validasi apakah nama tidak kosong
    if name == "" {
        http.Error(w, "Nama tidak ditemukan", http.StatusBadRequest)
        return
    }

    // Membuat objek user untuk menampung data yang diambil dari database
    var user models.User

    // Memanggil service untuk mengambil data berdasarkan nama
    if err := c.service.GetByName(&user, name); err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // Mengembalikan response JSON dengan link undangan dan nama user
    response := map[string]interface{}{
        "message":        "Undangan ditemukan",
        "namaUser":       user.Name,  // Menggunakan nama yang diambil dari database
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}





func (c *AllController) EditData(w http.ResponseWriter, r *http.Request) {
	var User models.User
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = c.service.EditData(&User)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(User)
}

func (c *AllController) DeleteData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.service.DeleteData(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Data deleted successfully"})
}	
