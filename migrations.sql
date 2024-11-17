CREATE TABLE sabores(
	sabor_id SERIAL PRIMARY KEY,
	nombre VARCHAR(50),
	precio NUMERIC(5, 2)
);

CREATE TABLE ingredientes(
	ingrediente_id SERIAL PRIMARY KEY,
	nombre VARCHAR(50) NOT NULL,
	precio NUMERIC(5, 2) NOT NULL
);

CREATE TABLE pizzas(
	pizza_id SERIAL PRIMARY KEY,
	tipo INT NOT NULL -- 1 = Pizza preestablecida 2 = pizza personalizada
);

CREATE TABLE pizzarecetas(
	pizzareceta_id SERIAL PRIMARY KEY,
	sabor_id INT NOT NULL,
	pizza_id INT NOT NULL,
	FOREIGN KEY (sabor_id) REFERENCES sabores(sabor_id),
	FOREIGN KEY (pizza_id) REFERENCES pizzas(pizza_id)
);

CREATE TABLE pizzapersonalizadas(
	pizzapersonalizada_id SERIAL PRIMARY KEY,
	precio NUMERIC(5, 2) NOT NULL DEFAULT 0,
	pizza_id INT NOT NULL,
	FOREIGN KEY (pizza_id) REFERENCES pizzas(pizza_id)
);

CREATE TABLE ingredientespizzapersonalizadas(
	ingredientepizzapersonalizada_id SERIAL PRIMARY KEY,
	ingrediente_id INT NOT NULL,
	pizzapersonalizada_id INT NOT NULL,
	FOREIGN KEY (ingrediente_id) REFERENCES ingredientes(ingrediente_id),
	FOREIGN KEY (pizzapersonalizada_id) REFERENCES pizzapersonalizadas(pizzapersonalizada_id)
);

CREATE TABLE pedidos(
	pedido_id SERIAL PRIMARY KEY,
	fecha TIMESTAMP NOT NULL,
	total NUMERIC(5,2) NOT NULL,
	promocion INT NOT NULL DEFAULT 0 -- 0 = Sin promoción, 1 = 2x1 y 2 = Delivery gratis
);

CREATE TABLE pedidopizza(
	pedidopizza_id SERIAL PRIMARY KEY,
	pedido_id INT NOT NULL,
	pizza_id INT NOT NULL,
	FOREIGN KEY (pedido_id) REFERENCES pedidos(pedido_id),
	FOREIGN KEY (pizza_id) REFERENCES pizzas(pizza_id)
);
