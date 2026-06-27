# 🚀 Simple REST API en Go + Fiber + GORM + SQLite

![Go](https://img.shields.io/badge/Go-1.XX-00ADD8?style=for-the-badge&logo=go)
![Fiber](https://img.shields.io/badge/Fiber-Web_Framework-black?style=for-the-badge)
![GORM](https://img.shields.io/badge/GORM-ORM-blue?style=for-the-badge)
![SQLite](https://img.shields.io/badge/SQLite-Database-003B57?style=for-the-badge&logo=sqlite)

## 📌 Sobre el proyecto

Este proyecto es una **API REST básica desarrollada en Go** con el objetivo de practicar y mejorar mis conocimientos en este lenguaje.

La aplicación utiliza:

- ⚡ **Fiber** como framework web para crear la API REST.
- 🗄️ **GORM** como ORM para manejar la comunicación con la base de datos.
- 💾 **SQLite** como base de datos ligera para almacenar información.

Este proyecto representa mis **primeros pasos desarrollando APIs REST con Go**, por lo que la implementación es sencilla y está enfocada principalmente en aprendizaje.

Actualmente mi mayor experiencia está en **TypeScript**, donde tengo un dominio más fuerte, y también cuento con un buen nivel trabajando con **C#**. Este proyecto nace como una oportunidad para conocer mejor el ecosistema de Go, su sintaxis, herramientas y forma de desarrollar aplicaciones backend.

> Este proyecto no busca ser una API avanzada, sino una base práctica para aprender Go y seguir mejorando mis habilidades backend.

---

# 🛠️ Tecnologías utilizadas

- Go
- Fiber
- GORM
- SQLite
- Air (Hot Reload)

---

# 📂 Estructura del proyecto

```bash
.
├── main.go
├── go.mod
├── go.sum
├── database/
├── models/
├── handlers/
├── services/
├── config/
├── routes/
└── tmp/
```

---

# ⚙️ Instalación

## Clonar el repositorio

```bash
git clone <url-del-repositorio>
```

Entrar al proyecto:

```bash
cd nombre-del-proyecto
```

---

# 📦 Instalar dependencias

Descargar las dependencias del proyecto:

```bash
go mod download
```

Organizar y limpiar dependencias:

```bash
go mod tidy
```

---

# ▶️ Ejecutar el proyecto

## Ejecutar normalmente

```bash
go run main.go
```

---

## Ejecutar con Air (Hot Reload)

Air permite reiniciar automáticamente el servidor cuando existen cambios en el código.

Instalar Air:

```bash
go install github.com/air-verse/air@latest
```

Ejecutar el proyecto usando Air:

```bash
air
```

---

# 📚 Manejo de dependencias

## Agregar una nueva dependencia

```bash
go get nombre-del-paquete
```

Ejemplo:

```bash
go get github.com/gofiber/fiber/v2
```

---

## Actualizar dependencias

```bash
go get -u ./...
```

---

## Limpiar dependencias no utilizadas

```bash
go mod tidy
```

---

## Desinstalar una dependencia

Remover una dependencia:

```bash
go get nombre-del-paquete@none
```

Ejemplo:

```bash
go get github.com/gofiber/fiber/v2@none
```

Después limpiar el módulo:

```bash
go mod tidy
```

---

# 🌐 Endpoints principales

| Método | Endpoint            | Descripción                    |
| ------ | ------------------- | ------------------------------ |
| GET    | `/api/v1/tasks`     | Obtener todos los registros    |
| GET    | `/api/v1/tasks/:id` | Obtener un registro específico |
| POST   | `/api/v1/tasks`     | Crear un registro              |
| PATCH  | `/api/v1/tasks/:id` | Actualizar un registro         |
| DELETE | `/api/v1/tasks/:id` | Eliminar un registro           |

---

# 🎯 Objetivos de aprendizaje

Con este proyecto busco practicar:

- Creación de APIs REST con Go.
- Manejo de rutas con Fiber.
- Uso de GORM como ORM.
- Conexión con bases de datos.
- Organización de proyectos backend.
- Manejo del sistema de módulos de Go.
- Uso de herramientas de desarrollo como Air.

---

# 👨‍💻 Autor

Proyecto creado como práctica personal para aprender y fortalecer mis habilidades en **Go Backend Development**.
