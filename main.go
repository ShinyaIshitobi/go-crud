package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// 構造体の埋め込み
type Water struct {
	Name    string   `json:"name"`
	Amount  int      `json:"amount"`
	Color   string   `json:"color"`
	Company *Company `json:"company"`
}

type Company struct {
	Name          string `json:"name"`
	Home          string `json:"home"`
	Establishment int    `json:"establishment"`
}

var waters []Water

func main() {
	// コンストラクタ的な
	r := mux.NewRouter()

	waters = append(waters, Water{Name: "irohasu", Amount: 500, Color: "green", Company: &Company{Name: "Coca-Cola", Home: "US", Establishment: 1892}})
	waters = append(waters, Water{"evian", 330, "red", &Company{"Danone", "Spain", 1919}})

	// ルーティング. FastAPIのadd_api_routeとかデコレータとか
	r.PathPrefix("/waters")
	r.HandleFunc("/", getWaters).Methods("GET")
	r.HandleFunc("/{name}", getWater).Methods("GET")
	r.HandleFunc("/", createWaters).Methods("POST")
	r.HandleFunc("/{name}", updateWater).Methods("PUT")
	r.HandleFunc("/{name}", deleteWater).Methods("DELETE")

	fmt.Println("start server at port :8080")
	http.ListenAndServe("localhost:8080", r)
}

func getWater(w http.ResponseWriter, r *http.Request) {
	// レスポンスヘッダをset
	w.Header().Set("Content-Type", "application/json")
	// map
	vars := mux.Vars(r)
	for _, water := range waters {
		if water.Name == vars["name"] {
			// JSON形式で書き込み送信
			err := json.NewEncoder(w).Encode(water)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
	}
}

func getWaters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(waters)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteWater(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	for i, water := range waters {
		if water.Name == vars["name"] {
			// それだけ退ける
			waters = append(waters[:i], waters[i+1:]...)
			break
		}
	}
	err := json.NewEncoder(w).Encode(waters)
	if err != nil {
		log.Fatal(err)
	}
}

func createWaters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// JSON形式のrequestを読み込む
	reqBody, _ := ioutil.ReadAll(r.Body)

	var water Water
	err := json.Unmarshal(reqBody, &water)
	if err != nil {
		log.Fatal(err)
	}

	waters = append(waters, water)
	err = json.NewEncoder(w).Encode(water)
	if err != nil {
		log.Fatal(err)
	}
}

func updateWater(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	reqBody, _ := ioutil.ReadAll(r.Body)
	var updateWater Water
	err := json.Unmarshal(reqBody, &updateWater)
	if err != nil {
		log.Fatal(err)
	}

	for i, water := range waters {
		if water.Name == vars["name"] {
			waters[i] = Water{
				Name:    water.Name,
				Amount:  updateWater.Amount,
				Color:   updateWater.Color,
				Company: updateWater.Company,
			}
		}
	}
}
