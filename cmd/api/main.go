package main

import (
	"net/http"

	"github.com/VictorManoelCostaDeBarros/go-intensivo/internal/entity"
	// "github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	// chi
	// r := chi.NewRouter()
	// r.Use(middleware.Logger)
	// r.Get("/order", Order)
	// http.ListenAndServe(":8888", r)

	e := echo.New()
	e.GET("/order", Order)
	e.Logger.Fatal(e.Start(":8888"))
}

// func Order(w http.ResponseWriter, r *http.Request) {
// 	// if r.Method != http.MethodGet {
// 	// 	w.WriteHeader(http.StatusMethodNotAllowed)
// 	// 	json.NewEncoder(w).Encode("Method nto allowed")
// 	// 	return
// 	// }

// 	order, err := entity.NewOrder("123", 1000, 10)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(err.Error())
// 	}
// 	order.CalculateFinalPrice()
// 	w.Header().Set("Content-type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(order)
// }

func Order(c echo.Context) error {
	order := entity.Order{
		ID:    "1",
		Price: 10,
		Tax:   1,
	}

	err := order.CalculateFinalPrice()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, order)
}
