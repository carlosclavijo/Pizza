CREATE TABLE sabores(
	sabor_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	nombre VARCHAR(50),
	precio NUMERIC(5, 2)
);

CREATE TABLE ingredientes(
	ingrediente_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	nombre VARCHAR(50) NOT NULL,
	precio NUMERIC(5, 2) NOT NULL
);

CREATE TABLE pizzas(
	pizza_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	tipo INT NOT NULL -- 1 = Pizza preestablecida 2 = pizza personalizada
);

CREATE TABLE pizzarecetas(
	pizzareceta_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	sabor_id UUID NOT NULL,
	pizza_id UUID NOT NULL,
	FOREIGN KEY (sabor_id) REFERENCES sabores(sabor_id),
	FOREIGN KEY (pizza_id) REFERENCES pizzas(pizza_id)
);

CREATE TABLE pizzapersonalizadas(
	pizzapersonalizada_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	precio NUMERIC(5, 2) NOT NULL DEFAULT 0,
	pizza_id UUID NOT NULL,
	FOREIGN KEY (pizza_id) REFERENCES pizzas(pizza_id)
);

CREATE TABLE ingredientespizzapersonalizadas(
	ingredientepizzapersonalizada_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	ingrediente_id UUID NOT NULL,
	pizzapersonalizada_id UUID NOT NULL,
	FOREIGN KEY (ingrediente_id) REFERENCES ingredientes(ingrediente_id),
	FOREIGN KEY (pizzapersonalizada_id) REFERENCES pizzapersonalizadas(pizzapersonalizada_id)
);

CREATE TABLE pedidos(
	pedido_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	fecha TIMESTAMP NOT NULL,
	total NUMERIC(5,2) NOT NULL
);

CREATE TABLE pedidopizza(
	pedidopizza_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	pedido_id UUID NOT NULL,
	pizza_id UUID NOT NULL,
	FOREIGN KEY (pedido_id) REFERENCES pedidos(pedido_id),
	FOREIGN KEY (pizza_id) REFERENCES pizzas(pizza_id)
);