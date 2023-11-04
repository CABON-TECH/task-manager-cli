package handlers

import (
    "testing"
)

func TestStatusHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/status", nil)
    w := httptest.NewRecorder()

    
    StatusHandler(w, req)

    
    if w.Code != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
    }
}
