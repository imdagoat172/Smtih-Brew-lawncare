package main

import (
    "encoding/json"
    "net/http"
    "sync"
)

type Info struct {
    Name     string   `json:"name"`
    Type     string   `json:"type"`
    Services []string `json:"services"`
    Email    string   `json:"email"`
    Phone    string   `json:"phone"`
}

type QuoteRequest struct {
    Name    string `json:"name"`
    Email   string `json:"email"`
    Phone   string `json:"phone"`
    Details string `json:"details"`
}

type QuoteResponse struct {
    ID      int    `json:"id"`
    Message string `json:"message"`
}

var (
    quoteStore = []QuoteRequest{}
    quoteID    = 0
    quoteMu    = &sync.Mutex{}
)

func main() {
    cors := func(w http.ResponseWriter) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    }

    http.HandleFunc("/api/info", func(w http.ResponseWriter, r *http.Request) {
        cors(w)
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        info := Info{
            Name: "Smith/Brew Lawncare LLC",
            Type: "Black-owned lawn care",
            Services: []string{"lawn mowing", "grass planting", "edge trimming", "leaf removal", "bushes", "flowers"},
            Email: "Marcussmith481@gmail.com",
            Phone: "678-544-7911",
        }
        json.NewEncoder(w).Encode(info)
    })

    http.HandleFunc("/api/quotes", func(w http.ResponseWriter, r *http.Request) {
        cors(w)
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        if r.Method == http.MethodPost {
            var quote QuoteRequest
            if err := json.NewDecoder(r.Body).Decode(&quote); err != nil {
                w.WriteHeader(http.StatusBadRequest)
                json.NewEncoder(w).Encode(map[string]string{"error": "invalid payload"})
                return
            }
            if quote.Name == "" || quote.Email == "" || quote.Phone == "" || quote.Details == "" {
                w.WriteHeader(http.StatusBadRequest)
                json.NewEncoder(w).Encode(map[string]string{"error": "all fields are required"})
                return
            }

            quoteMu.Lock()
            quoteID++
            quoteStore = append(quoteStore, quote)
            quoteMu.Unlock()

            w.WriteHeader(http.StatusCreated)
            json.NewEncoder(w).Encode(QuoteResponse{ID: quoteID, Message: "Your quote request was received."})
            return
        }

        if r.Method == http.MethodGet {
            quoteMu.Lock()
            defer quoteMu.Unlock()
            json.NewEncoder(w).Encode(map[string]interface{}{"count": len(quoteStore), "quotes": quoteStore})
            return
        }

        w.WriteHeader(http.StatusMethodNotAllowed)
    })

    http.ListenAndServe(":8080", nil)
}
