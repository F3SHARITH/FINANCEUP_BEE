# APIS_FINANCEUP

✅ **API REST** en Go para la gestión integral de la plataforma **FinanceUp**: finanzas personales, educación financiera, intermediación bancaria y soporte al usuario.

El backend está compuesto por **5 microservicios independientes**, cada uno con su propio esquema en la base de datos PostgreSQL `FinanceUp`, desarrollados con Go nativo + Gorilla Mux y CORS habilitado.

| Microservicio | Descripción | Puerto |
|---|---|---|
| [`AUTH`](./AUTH/README.md) | Autenticación, usuarios, roles y auditoría | `8084` |
| [`FINANZAS`](./FINANZAS/README.md) | Ingresos, egresos, metas e inversiones | `8081` |
| [`EDUCACION`](./EDUCACION/README.md) | Módulos educativos y progreso del usuario | `8080` |
| [`NEGOCIO`](./NEGOCIO/README.md) | Bancos, asesores, leads y créditos | `8085` |
| [`SOPORTE`](./SOPORTE/README.md) | PQR, adjuntos y registro de actividad | `8083` |

---

## Arquitectura General

```text
APIS_FINANCEUP/
├── AUTH/               → Autenticación y autorización (puerto 8084)
│   ├── config/         → Conexión a base de datos
│   ├── controller/     → Controladores HTTP
│   ├── models/         → Estructuras de datos
│   └── router/         → Definición de rutas
├── EDUCACION/          → Módulo educativo financiero (puerto 8080)
│   ├── config/
│   ├── controllers/
│   ├── models/
│   └── routes/
├── FINANZAS/           → Gestión financiera personal (puerto 8081)
│   ├── config/
│   ├── controllers/
│   ├── models/
│   └── routes/
├── NEGOCIO/            → Lógica bancaria y comercial (puerto 8085)
│   ├── config/
│   ├── controllers/
│   ├── models/
│   └── routes/
├── SOPORTE/            → PQR y soporte al usuario (puerto 8083)
│   ├── config/
│   ├── controller/
│   ├── models/
│   └── routes/
└── PostgreSQL/         → Scripts SQL (crear, insertar, actualizar)
```

---

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones

- [Go 1.26.2](https://go.dev/doc/install)
- [Gorilla Mux v1.8.1](https://github.com/gorilla/mux) — enrutador HTTP
- [lib/pq v1.12.3](https://github.com/lib/pq) — driver PostgreSQL para Go
- [PostgreSQL](https://www.postgresql.org/) — base de datos relacional (5 esquemas independientes)

---

### Variables de Entorno

La conexión a la base de datos está definida directamente en cada `config/db.go`. Para modificar los parámetros de conexión sin tocar el código, se recomienda reemplazar los valores fijos por variables de entorno.

```
# Base de datos compartida por todos los servicios
PGHOST=localhost
PGPORT=5432
PGUSER=postgres
PGPASSWORD=<contraseña>
PGDATABASE=FinanceUp        # o FINANCE_UP / financeup según el servicio

# Esquemas por servicio (no configurables por variable, definidos en config/db.go)
# AUTH      → search_path=auth
# FINANZAS  → search_path=finanzas
# EDUCACION → search_path=educacion
# NEGOCIO   → search_path=negocio
# SOPORTE   → search_path=soporte
```

> **Nota:** Los valores actuales en `config/db.go` de cada microservicio son los que se usan por defecto. Se recomienda externalizar estas variables antes de desplegar en producción.

---

## Ejecución del Proyecto

### Requisitos previos

- Go 1.26.2 instalado
- PostgreSQL activo y accesible
- Base de datos `FinanceUp` creada con los scripts de `PostgreSQL/`

### Inicializar la base de datos

```bash
# Ejecutar los scripts SQL en el siguiente orden:
# 1. Crear estructura (esquemas, tablas, triggers)
psql -U postgres -d FinanceUp -f "PostgreSQL/crear sql (1).txt"

# 2. Insertar datos semilla
psql -U postgres -d FinanceUp -f "PostgreSQL/insertar sql (1).txt"

# 3. Aplicar actualizaciones (si aplica)
psql -U postgres -d FinanceUp -f "PostgreSQL/actualizar sql (1).txt"
```

### Clonar el repositorio

```bash
# 1. Clonar el repositorio
git clone https://github.com/F5Alejo/APIS_FINANCEUP.git

# 2. Moverse a la carpeta del repositorio
cd APIS_FINANCEUP

# 3. Moverse a la rama develop
git pull origin develop && git checkout develop
```
---

## Endpoints Disponibles

Todos los endpoints responden en **JSON** y soportan `GET`, `POST`, `PUT`, `DELETE`. CORS está habilitado en todos los servicios (`*`).


---

## Modelo de Datos

Todos los servicios comparten la misma base de datos PostgreSQL (`FinanceUp`) operando sobre **esquemas separados**.

```
FinanceUp (base de datos)
├── auth
│   ├── rol
│   ├── tipo_documento
│   ├── usuario
│   ├── credencial
│   ├── usuario_rol
│   └── auditoria_login
├── finanzas
│   ├── categoria
│   ├── movimiento_ingreso_egreso
│   ├── tipo_ingreso
│   ├── finanzas
│   ├── tipo_inversion
│   ├── nivel_riesgo
│   ├── movimiento_inversion
│   ├── tipo_ingreso_inversion
│   ├── inversion
│   ├── editar_meta
│   ├── movimiento_meta
│   ├── tipo_ingreso_meta
│   └── meta
├── educacion
│   ├── modulo_educativo
│   ├── contenido
│   ├── leccion
│   ├── progreso_educativo
│   └── progreso_leccion
├── negocio
│   ├── banco
│   ├── asesor_bancario
│   ├── contacto_asesor
│   ├── producto_crediticio
│   ├── lead
│   ├── conversacion_usuario_asesor
│   ├── credito_desembolsado
│   └── transaccion_comision
└── soporte
    ├── estado_pqr
    ├── pqr
    ├── adjunto
    └── registro_actividad
```

Todas las tablas incluyen los campos de auditoría `fecha_creacion` y `fecha_modificacion`, gestionados automáticamente mediante el trigger `fn_update_fecha_modificacion`.
---

📊 [Ver diagrama del modelo de datos](./modelo_datos.png)

---

## Ramas del Repositorio

| Rama | Propósito |
|---|---|
| `main` | Versión estable |
| `develop` | Integración de nuevas funcionalidades |
| `release` | Preparación para producción |
| `javier` | Rama de trabajo — módulo AUTH |
| `sharith` | Rama de trabajo — módulos EDUCACION / FINANZAS |

---