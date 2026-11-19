# Sistema de Gestión de Biblioteca

CRUD de Libros y Autores con Go, React, PostgreSQL, DDD, Facade y Visitor.

## Requisitos
- Go 1.22 o superior
- Node.js 18 o superior
- PostgreSQL 14 o superior

## 1. Base de datos
Crea la base de datos:
```sql
CREATE DATABASE biblioteca;
```
Ejecuta el contenido de `backend/migrations/001_init.sql` sobre esa base de datos.

## 2. Backend
Desde la carpeta `backend`:
```bash
cp .env.example .env
```
Edita `.env` y coloca tu usuario y contraseña de PostgreSQL. En Windows puedes copiar `.env.example` como `.env` manualmente.

Luego:
```bash
go mod tidy
go run ./cmd/api
```
API: http://localhost:8080

## 3. Frontend
En otra terminal, desde `frontend`:
```bash
npm install
npm run dev
```
Abre http://localhost:5173

## Endpoints
| Método | Ruta | Acción |
|---|---|---|
| GET | /api/libros | Listar libros |
| GET | /api/libros/{id} | Obtener libro |
| POST | /api/libros | Crear libro |
| PUT | /api/libros/{id} | Actualizar libro |
| DELETE | /api/libros/{id} | Eliminar libro |
| GET | /api/autores | Listar autores |
| GET | /api/autores/{id} | Obtener autor |
| POST | /api/autores | Crear autor |
| PUT | /api/autores/{id} | Actualizar autor |
| DELETE | /api/autores/{id} | Eliminar autor |
| GET | /api/reportes/inventario | Reporte Visitor |

## Ejemplos JSON
Autor:
```json
{"nombre":"Gabriel García Márquez","nacionalidad":"Colombiana"}
```
Libro:
```json
{"titulo":"Cien años de soledad","isbn":"9780307474728","anio":1967,"autor_id":1}
```

## Patrones
- **Facade:** `internal/interfaces/http/facade/biblioteca_facade.go` centraliza el acceso de los handlers a los servicios de aplicación.
- **Visitor:** `internal/domain/libro/visitor.go` define el visitante y `InventarioVisitor` recorre los libros para generar un resumen sin añadir esa lógica a la entidad Libro.

## Reglas
- No se puede registrar un libro sin un autor existente.
- No se puede eliminar un autor que tenga libros asociados (integridad referencial).
- El ISBN debe ser único.
