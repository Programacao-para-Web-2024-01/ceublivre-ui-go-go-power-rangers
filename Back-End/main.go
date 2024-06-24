package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

type Promotion struct {
	ID             int       `json:"id"`
	Desconto       float64   `json:"desconto"`
	Produtos       []string  `json:"produtos"`
	ValidadeInicio time.Time `json:"validade_inicio"`
	ValidadeFim    time.Time `json:"validade_fim"`
	Codigo         string    `json:"codigo,omitempty"`
}

var (
	promotions = []Promotion{}
	nextID     = 1
	mu         sync.Mutex
)

func main() {
	http.HandleFunc("/promotions", promotionsHandler)
	http.HandleFunc("/promotions/apply", applyPromotionHandler)
	http.HandleFunc("/promotions/report", reportHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func promotionsHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	var promo Promotion
	if err := json.NewDecoder(r.Body).Decode(&promo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPost:
		promo.ID = nextID
		nextID++
		promotions = append(promotions, promo)
		w.WriteHeader(http.StatusCreated)
	case http.MethodPut:
		for i, p := range promotions {
			if p.ID == promo.ID {
				promotions[i] = promo
				w.WriteHeader(http.StatusOK)
				break
			}
		}
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	json.NewEncoder(w).Encode(promo)
}

func applyPromotionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	var request struct {
		Codigo      string   `json:"codigo"`
		Produtos    []string `json:"produtos"`
		TotalCompra float64  `json:"total_compra"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var appliedPromo *Promotion
	mu.Lock()
	defer mu.Unlock()
	now := time.Now()
	for _, promo := range promotions {
		if promo.Codigo == request.Codigo && now.After(promo.ValidadeInicio) && now.Before(promo.ValidadeFim) {
			for _, produto := range request.Produtos {
				for _, p := range promo.Produtos {
					if p == produto {
						appliedPromo = &promo
						break
					}
				}
				if appliedPromo != nil {
					break
				}
			}
		}
	}

	if appliedPromo != nil {
		desconto := request.TotalCompra * (appliedPromo.Desconto / 100)
		json.NewEncoder(w).Encode(map[string]float64{"desconto": desconto, "total": request.TotalCompra - desconto})
	} else {
		http.Error(w, "Promoção não aplicável", http.StatusBadRequest)
	}
}

func reportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	json.NewEncoder(w).Encode(promotions)
}
