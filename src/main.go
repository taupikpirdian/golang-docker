package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/check", handler)
	fmt.Println("Server is running on port 8081...")
	http.ListenAndServe(":8081", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	n := []int{1, 8, 6, 2, 5, 4, 8, 3, 7} // Anda dapat mengambil nilai dari query jika diinginkan
	result := maxArea(n)

	// Struct untuk respons
	response := struct {
		Status string `json:"status"`
		Data   int    `json:"data"`
	}{
		Status: "success",
		Data:   result,
	}

	// Set header ke application/json
	w.Header().Set("Content-Type", "application/json")
	// Encode ke JSON dan kirimkan ke response writer
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode JSON: %v", err), http.StatusInternalServerError)
	}
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxAmountWater := 0
	for left < right {
		// rumus
		// v = h * w
		// karena ini air, maka gunakan tinggi minimal, karena air mengikuti wadahnya

		width := right - left
		volume := min(height[left], height[right]) * width

		if volume > maxAmountWater {
			maxAmountWater = volume
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxAmountWater
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
