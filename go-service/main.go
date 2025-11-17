package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"

    elastic "github.com/elastic/go-elasticsearch/v8"
    "github.com/elastic/go-elasticsearch/v8/esapi"
)

var es *elastic.Client


func init() {
    cfg := elastic.Config{
        Addresses: []string{"http://elasticsearch:9200"},
    }

    client, err := elastic.NewClient(cfg)
    if err != nil {
        log.Fatalf("Error creating ES client: %s", err)
    }

    es = client
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q")
    if q == "" {
        http.Error(w, "missing query param 'q'", http.StatusBadRequest)
        return
    }

    query := fmt.Sprintf(
        `{"query": {"multi_match": {"query": "%s", "fields": ["user_id"]}}}`,
        q,
    )

    req := esapi.SearchRequest{
        Index: []string{"user_activity"},
        Body:  strings.NewReader(query),
    }

    res, err := req.Do(context.Background(), es)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer res.Body.Close()

    var data map[string]interface{}
    if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}

func main() {
    http.HandleFunc("/search", searchHandler)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Listening on :%s...", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
