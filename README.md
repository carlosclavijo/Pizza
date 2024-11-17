# Pizza

Proyecto 2 del módulo de Arquitectura de Software.

- Versión de Go 1.23
- Base de datos PostgreSQL

Librerías
- [uuid](https://www.github.com/gofrs/uuid) para el uso de UUID
- [godotenv](https://www.github.com/joho/godotenv) para cargar el archivo ".env"
- [gorm](https://www.gorm.io/gorm) ORM de Go
- [postgresDriver](https://www.gorm.io/driver/postgres) Driver de Postgres de GORM

Para instalar cada librería se utiliza el comando: "go get <url>"
Ejemplo: "go get github.com/gofrs/uuid"

EndPoints:

1. PostIngrediente(localhost:8080/ingredientes)
    Parámetros: nombre y precio

2. GetIngredientes(localhost:8080/ingredientes)

3. GetIngredienteById(localhost:8080/ingredientes/{id})

4. PostSabor(localhost:8080/sabores)
    Parámetros: nombre y precio

5. GetSabores(localhost:8080/sabores)

6. GetSaborById(localhost:8080/sabores/{id})

7. PostPedido(localhost:8080/pedidos)
    Parámetros: Pizza{sabores, personalizadas}, fecha, precio_delivery
    Pizza.sabores es un array de los id de los sabores de pizza
    Pizza.personalizadas es un array que contiene arrays de los id de ingredientes.

    Ejemplo: {
    "pizzas": {
        "sabores": [1, 2], //2 pizza de sabor 1 y 2
        "personalizadas": [ //2 pizzas, una con 3 topings y otra con 2
            [1, 2, 3],
            [1, 3]
        ]
    },
    "fecha": "2024-11-15T00:00:00.000000000Z", //Si se reconoce que es día de descuento se aplica este al final
    "precio_delivery": 10
    }

    Al usar este endpoint se realiza toda la transaccion del pedido, desde el armado de pizzas, tanto predefinidas como personalizadas, la aplicación de descuentos y el total a pagar
