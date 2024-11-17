CREATE TABLE sabors(
	sabor_id SERIAL PRIMARY KEY,
	nombre VARCHAR(50) NOT NULL UNIQUE,
	precio NUMERIC(5, 2) NOT NULL
);

CREATE TABLE ingredientes(
	ingrediente_id SERIAL PRIMARY KEY,
	nombre VARCHAR(50) NOT NULL UNIQUE,
	precio NUMERIC(5, 2) NOT NULL
);

CREATE TABLE pedidos(
	pedido_id SERIAL PRIMARY KEY,
	fecha TIMESTAMP DEFAULT NOW(),
	precio_pizzas NUMERIC(5,2) NOT NULL DEFAULT 0,
	precio_delivery NUMERIC(5,2) NOT NULL DEFAULT 0,
	descuento NUMERIC(5, 2) NOT NULL DEFAULT 0,
	precio_total NUMERIC(5,2) NOT NULL DEFAULT 0,
	promocion INT NOT NULL DEFAULT 0 -- 0 = Sin promoción, 1 = 2x1 y 2 = Delivery gratis
);

CREATE TABLE pizza_sabors(
	pizza_sabor_id SERIAL PRIMARY KEY,
	pedido_id INT NOT NULL,
	sabor_id INT NOT NULL,
	FOREIGN KEY (pedido_id) REFERENCES pedidos(pedido_id),
	FOREIGN KEY (sabor_id) REFERENCES sabors(sabor_id)
);

CREATE TABLE pizza_personalizadas(
	pizza_personalizada_id SERIAL PRIMARY KEY,
	precio NUMERIC(5, 2) NOT NULL DEFAULT 0,
	pedido_id INT NOT NULL,
	FOREIGN KEY (pedido_id) REFERENCES pedidos(pedido_id)
);

CREATE TABLE ingrediente_pizzas(
	ingrediente_pizza_id SERIAL PRIMARY KEY,
	ingrediente_id INT NOT NULL,
	pizza_personalizada_id INT NOT NULL,
	FOREIGN KEY (ingrediente_id) REFERENCES ingredientes(ingrediente_id),
	FOREIGN KEY (pizza_personalizada_id) REFERENCES pizza_personalizadas(pizza_personalizada_id)
);

-- Seeders
INSERT INTO ingredientes(nombre, precio) VALUES('peperoni', 5), ('muzzarella', 8), ('pina', 4);
INSERT INTO sabors(nombre, precio) VALUES('4 quesos', 50), ('hawaiana', 40), ('anchoas', 60);
