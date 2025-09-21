# Microservicio de Productos (Go + Gin + GORM + SQL Server)
git clone https://github.com/yourname/products-ms.git
cd products-ms
cp .env.example .env
# Edita .env con tus credenciales


# Instalar dependencias
go mod tidy


# Ejecutar migraciones manuales (SQL Server Management Studio o sqlcmd)
# migrations/001_init.sql


# Correr servidor
go run ./cmd/server
```


## Endpoints


- `GET /health` → estado
- `GET /api/v1/categories` → lista de categorías
- `GET /api/v1/products` → lista paginada de productos


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


```
GET /api/v1/products?page=1&page_size=10&sort=price&order=desc
GET /api/v1/products?q=bluetooth&in_stock=true
GET /api/v1/products?category_id=2&min_price=20&max_price=60
```


## Respuesta JSON (ejemplo)
```json
{
"items": [
{
"id": 1,
"sku": "ELEC-001",
"name": "Audífonos Bluetooth",
"price": 99.9,
"stock": 50,
"active": true,
"category_id": 1,
"category": {"id":1, "name":"Electrónica", "slug":"electronica"}
}
],
"page": 1,
"page_size": 10,
"total_items": 100,
"total_pages": 10,
"has_next": true,
"has_prev": false
}
```


## Notas
- El repo usa **GORM** con el driver `sqlserver`. Alternativamente puedes escribir repos `database/sql` con `OFFSET ... FETCH NEXT` si prefieres SQL puro.
- Si quieres exponer **HATEOAS** o links de páginas siguientes/previas, agrega helpers en `pkg/pagination`.
- Para entornos productivos, añade logs estructurados, métricas y validaciones adicionales.