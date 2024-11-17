package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/carlosclavijo/pizza/internal/models"
	"github.com/carlosclavijo/pizza/internal/repository"
	"github.com/tidwall/gjson"
)

func (repo Repository) PostPedido(w http.ResponseWriter, r *http.Request) {
	var pedido models.Pedido
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error al leer el json request\n", err)
		return
	}

	fecha := gjson.Get(string(body), "fecha").String()
	pedido.Fecha, err = time.Parse("2006-01-02T15:04:05.999Z", fecha)
	if err != nil {
		log.Println("Error al concatenar la fecha\n", err)
	}

	pedido.PrecioDelivery = float32(gjson.Get(string(body), "precio_delivery").Float())

	//Insert de Pedido
	if int(pedido.Fecha.Weekday()) == 2 || int(pedido.Fecha.Weekday()) == 3 {
		pedido.Promocion = addPromocion(1)
	} else if int(pedido.Fecha.Weekday()) == 4 {
		pedido.Promocion = addPromocion(2)
	}
	result := repository.InsertPedido(&pedido)
	if result != nil {
		log.Println("Error al insertar pedido\n", result)
	}

	//Aglomerar de pizzas con sabor
	pizzasSabores := gjson.Get(string(body), "pizzas.sabores").Array()
	for _, s := range pizzasSabores {
		sabor := models.NewPizzaSabor(int(s.Int()), pedido.PedidoId)
		pedido.Pizzas = append(pedido.Pizzas, sabor)
	}

	//Aglomerar de pizzas personalizadas
	pizzaPersonalizadas := gjson.Get(string(body), "pizzas.personalizadas").Array()
	for _, s := range pizzaPersonalizadas {
		toppings := s.Array()
		var ingredientes []models.Ingrediente
		for _, t := range toppings {
			var ingrediente models.Ingrediente
			result := repository.GetIngredientById(&ingrediente, int(t.Int()))
			if result != nil {
				log.Println("Error al traer los ingredientes\n", result)
			}
			ingredientes = append(ingredientes, ingrediente)
		}
		personalizada := models.NewPizzaPersonalizada(ingredientes, pedido.PedidoId)
		pedido.Pizzas = append(pedido.Pizzas, personalizada)
	}

	//Insertar todas las pizzas
	for i := 0; i < len(pedido.Pizzas); i++ {
		result := repository.InsertarPizza(pedido.Pizzas[i])
		if result != nil {
			log.Println("Error al insertar las pizzas\n", result)
		}
		pedido.PrecioPizzas += pedido.Pizzas[i].CalcularPrecio()
	}

	//Aplicar Descuentos
	pedido.PrecioTotal = pedido.PrecioPizzas + pedido.PrecioDelivery
	if pedido.Promocion != nil {
		pedido.Descuento = aplicarPromocion(&pedido, pedido.Promocion)
	}
	pedido.PrecioTotal = pedido.CalcularTotal()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pedido)
}

func addPromocion(dia int) models.IPromocion {
	if dia == 1 {
		return models.NewPromo2x1()
	} else if dia == 2 {
		return models.NewPromoDelivery()
	}
	return nil
}

func aplicarPromocion(pedido *models.Pedido, promo models.IPromocion) float32 {
	return promo.Descuento(pedido)
}
