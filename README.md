# Microservicio de Productos (Go + Gin + GORM + SQL Server)
git clone https://github.com/JoseFelixTheOne/PRODUCTS-MS.git

cd products-ms

Cambia tus credenciales en el .env


# Instalar dependencias
go mod tidy


### Ejecutar migraciones manuales (SQL Server Management Studio o sqlcmd)
### migrations/001_init.sql


## Correr servidor
go run ./cmd/server



## Endpoints


- GET /health → estado
- GET /api/v1/categories → lista de categorías
- GET /api/v1/products → lista paginada de productos


### Query params de `/api/v1/products`


- `page` (int, default 1)
- `page_size` (int, default 20, máx 200)
- `q` (string) → busca por `Name` o `SKU`
- `category_id` (uint)
- `min_price` (float)
- `max_price` (float)
- `in_stock` (bool) → true: stock > 0, false: stock <= 0
- `active` (bool)
- `sort` (string) → `name|price|created_at` (default: `created_at`)
- `order` (string) → `asc|desc` (default: `asc`)


### Ejemplos



GET /api/v1/products?page=1&page_size=10&sort=price&order=desc
GET /api/v1/products?q=bluetooth&in_stock=true
GET /api/v1/products?category_id=2&min_price=20&max_price=60

###⚠️ Nota importante sobre SQL Server

Para que el microservicio pueda conectarse a SQL Server asegúrate de:
Habilitar el protocolo TCP/IP en SQL Server Configuration Manager.
Configurar el puerto estático 1433 en la pestaña IPAll.
Reiniciar el servicio de SQL Server.
Sin esto, el DSN configurado en .env no funcionará y obtendrás errores de conexión.
