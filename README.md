# FINANCEUP_BEE

API REST en Go para la gestión integral de la plataforma **FinanceUp**. Administra autenticación, soporte, educación financiera, finanzas personales y el ecosistema bancario-comercial sobre una base de datos PostgreSQL organizada en 5 esquemas independientes.

---

## Especificaciones Técnicas

| Tecnología | Versión |
|---|---|
| [Go](https://go.dev/) | 1.26.2 |
| [Beego v2](https://github.com/beego/beego) | v2.3.10 |
| [lib/pq](https://github.com/lib/pq) — driver PostgreSQL | v1.12.3 |
| [godotenv](https://github.com/joho/godotenv) | v1.5.1 |
| [PostgreSQL](https://www.postgresql.org/) | — |

- **Base de datos:** `financeup_bee`
- **Esquemas:** `auth`, `soporte`, `educacion`, `finanzas`, `negocio`
- **Puerto por defecto:** `8080`

---

## Estructura del Proyecto

```
FINANCEUP_BEE/
├── conf/
│   └── app.conf          # Configuración de Beego
├── routers/              # Definición de rutas
├── swagger/              # Documentación Swagger (modo dev)
├── .env                  # Variables de entorno (no se sube a GitHub)
├── .gitignore
├── go.mod
├── go.sum
└── main.go
```

---

## Configuración

### Variables de entorno

Crea un archivo `.env` en la raíz del proyecto:

```env
FINANCE_BEE_HTTP_PORT=8080
FINANCE_BEE_RUN_MODE=dev
FINANCE_BEE_PG_HOST=localhost
FINANCE_BEE_PG_PORT=5432
FINANCE_BEE_PG_USER=postgres
FINANCE_BEE_PG_PASS=postgres
FINANCE_BEE_PG_DBNAME=financeup_bee
FINANCE_BEE_PG_SCHEMA=negocio
```

> ⚠️ El archivo `.env` está incluido en `.gitignore` y **nunca debe subirse al repositorio**.

La cadena de conexión se construye automáticamente:
```
postgres://<PGuser>:<PGpass>@<PGhost>:<PGport>/<PGdb>?sslmode=disable&search_path=<PGschema>
```

---

## Ejecución del Proyecto

```bash
# 1. Clonar el repositorio
git clone https://github.com/<tu-usuario>/FINANCEUP_BEE.git

# 2. Entrar a la carpeta
cd FINANCEUP_BEE

# 3. Descargar dependencias
go mod tidy

# 4. Crear el archivo .env con tus credenciales

# 5. Ejecutar el servidor
go run main.go
```

El servidor quedará escuchando en **`http://localhost:8080`**.

En modo `dev` la documentación Swagger estará disponible en:
```
http://localhost:8080/swagger/
```

---

## Modelo de Datos

La base de datos está organizada en **5 esquemas**. Todos comparten la misma convención de auditoría: campos `activo`, `fecha_creacion` y `fecha_modificacion` en cada tabla, actualizados automáticamente por el trigger global `fn_update_fecha_modificacion()`.

> El esquema `auth` es el **núcleo transversal** — `auth.usuario` es referenciado por los 4 esquemas restantes.

---

### Esquema `auth` — Autenticación y usuarios

Gestiona el acceso, identidad y trazabilidad de todos los usuarios de la plataforma.

| Tabla | Descripción |
|---|---|
| `tipo_documento` | Catálogo de tipos de documento de identidad |
| `rol` | Roles disponibles en el sistema |
| `usuario` | Usuarios registrados — nodo central de toda la base de datos |
| `credencial` | Hash de contraseña, algoritmo y control de bloqueo por usuario |
| `usuario_rol` | Relación muchos a muchos entre usuarios y roles |
| `auditoria_login` | Log inmutable de eventos de sesión (login, logout, acceso denegado) |

---

### Esquema `soporte` — PQR y actividad

Gestiona las peticiones, quejas y reclamos de los usuarios, así como el registro de actividad general del sistema.

| Tabla | Descripción |
|---|---|
| `estado_pqr` | Catálogo de estados de una PQR |
| `pqr` | Solicitudes PQR creadas por usuarios |
| `adjunto` | Archivos adjuntos vinculados a una PQR |
| `registro_actividad` | Log inmutable de acciones realizadas por usuarios |

---

### Esquema `educacion` — Contenido educativo

Gestiona el material educativo financiero y el progreso de aprendizaje de cada usuario.

| Tabla | Descripción |
|---|---|
| `modulo_educativo` | Módulos de aprendizaje con nivel (básico, intermedio, avanzado) |
| `contenido` | Contenidos multimedia independientes (videos) |
| `leccion` | Lecciones vinculadas a un módulo y opcionalmente a un contenido |
| `progreso_educativo` | Avance de un usuario por módulo (porcentaje y calificación) |
| `progreso_leccion` | Avance de un usuario por lección individual |

---

### Esquema `finanzas` — Finanzas personales

Gestiona ingresos, egresos, inversiones y metas de ahorro de cada usuario.

| Tabla | Descripción |
|---|---|
| `categoria` | Categorías de movimientos financieros |
| `movimiento_ingreso_egreso` | Movimientos de dinero personales |
| `tipo_ingreso` | Tipos de pago vinculados a movimientos personales |
| `finanzas` | Registro de presupuesto, gasto y saldo disponible por usuario |
| `tipo_inversion` | Catálogo de tipos de inversión |
| `nivel_riesgo` | Catálogo de niveles de riesgo de inversión |
| `movimiento_inversion` | Movimientos de dinero asociados a inversiones |
| `tipo_ingreso_inversion` | Tipos de pago vinculados a inversiones |
| `inversion` | Inversiones registradas por usuario |
| `editar_meta` | Configuración de una meta de ahorro |
| `movimiento_meta` | Movimientos de dinero asociados a metas |
| `tipo_ingreso_meta` | Tipos de pago vinculados a metas |
| `meta` | Metas de ahorro por usuario |

---

### Esquema `negocio` — Ecosistema bancario

Gestiona bancos, asesores, productos crediticios, leads comerciales, créditos desembolsados y comisiones.

| Tabla | Descripción |
|---|---|
| `banco` | Bancos aliados con contacto, comisión y estado |
| `producto_crediticio` | Productos de crédito con rangos de monto, tasa y plazo |
| `asesor_bancario` | Asesores vinculados a un banco con especialidad |
| `contacto_asesor` | Canales de contacto y horarios de disponibilidad del asesor |
| `lead` | Solicitudes de crédito generadas por usuarios |
| `conversacion_usuario_asesor` | Historial de mensajes entre usuario y asesor por lead |
| `credito_desembolsado` | Créditos aprobados con saldo, estado y fechas |
| `transaccion_comision` | Comisiones generadas por crédito desembolsado |

---

### Relaciones entre esquemas

Todos los esquemas dependen de `auth.usuario` como fuente de identidad del usuario:

| Esquema | Tablas que referencian `auth.usuario` |
|---|---|
| `soporte` | `pqr`, `registro_actividad` |
| `educacion` | `progreso_educativo`, `progreso_leccion` |
| `finanzas` | `finanzas`, `inversion`, `meta` |
| `negocio` | `lead`, `conversacion_usuario_asesor`, `credito_desembolsado` |

---

### Estados válidos por campo

| Esquema | Campo | Valores permitidos |
|---|---|---|
| `auth` | `usuario.estado` | `activo`, `inactivo`, `suspendido` |
| `auth` | `credencial.algoritmo` | `bcrypt`, `argon2`, `sha256` |
| `auth` | `auditoria_login.tipo_evento` | `login`, `logout`, `cambio_contrasena`, `acceso_denegado` |
| `auth` | `auditoria_login.estado_evento` | `exitoso`, `fallido` |
| `educacion` | `modulo_educativo.nivel` | `basico`, `intermedio`, `avanzado` |
| `negocio` | `banco.estado` | `activo`, `inactivo` |
| `negocio` | `lead.estado_lead` | `nuevo`, `contactado`, `en_proceso`, `aprobado`, `rechazado`, `cancelado` |
| `negocio` | `conversacion_usuario_asesor.tipo_contacto` | `email`, `telefono`, `whatsapp`, `presencial` |
| `negocio` | `credito_desembolsado.estado_credito` | `activo`, `pagado`, `vencido`, `cancelado` |
| `negocio` | `transaccion_comision.estado` | `pendiente`, `pagada`, `cancelada` |

---

> ⚠️ **Importante — `ON DELETE RESTRICT`:** Todas las claves foráneas usan `RESTRICT`. Para eliminar cualquier registro debes primero eliminar todos los que dependen de él, de adentro hacia afuera.

---

## Endpoints Disponibles

Todos los endpoints responden en formato **JSON** y soportan los métodos `GET`, `POST`, `PUT`, `DELETE`.

### `auth`
| Recurso | Ruta base |
|---|---|
| Roles | `/v1/rol` |
| Tipos de documento | `/v1/tipo_documento` |
| Usuarios | `/v1/usuario` |
| Credenciales | `/v1/credencial` |
| Usuario-Rol | `/v1/usuario_rol` |
| Auditoría login | `/v1/auditoria_login` |

### `soporte`
| Recurso | Ruta base |
|---|---|
| Estado PQR | `/v1/estado_pqr` |
| PQR | `/v1/pqr` |
| Adjuntos | `/v1/adjunto` |
| Registro de actividad | `/v1/registro_actividad` |

### `educacion`
| Recurso | Ruta base |
|---|---|
| Módulos educativos | `/v1/modulo_educativo` |
| Contenidos | `/v1/contenido` |
| Lecciones | `/v1/leccion` |
| Progreso educativo | `/v1/progreso_educativo` |
| Progreso lección | `/v1/progreso_leccion` |

### `finanzas`
| Recurso | Ruta base |
|---|---|
| Categorías | `/v1/categoria` |
| Movimientos ingreso/egreso | `/v1/movimiento_ingreso_egreso` |
| Tipos de ingreso | `/v1/tipo_ingreso` |
| Finanzas | `/v1/finanzas` |
| Tipos de inversión | `/v1/tipo_inversion` |
| Niveles de riesgo | `/v1/nivel_riesgo` |
| Movimientos inversión | `/v1/movimiento_inversion` |
| Inversiones | `/v1/inversion` |
| Metas | `/v1/meta` |

### `negocio`
| Recurso | Ruta base |
|---|---|
| Bancos | `/v1/banco` |
| Asesores bancarios | `/v1/asesor_bancario` |
| Contacto asesor | `/v1/contacto_asesor` |
| Productos crediticios | `/v1/producto_crediticio` |
| Leads | `/v1/lead` |
| Conversaciones | `/v1/conversacion_usuario_asesor` |
| Créditos desembolsados | `/v1/credito_desembolsado` |
| Transacciones comisión | `/v1/transaccion_comision` |

### Ejemplo de uso

```bash
# Obtener todos los registros
GET http://localhost:8080/v1/<recurso>

# Obtener un registro por id
GET http://localhost:8080/v1/<recurso>/1

# Crear un registro
POST http://localhost:8080/v1/<recurso>

# Actualizar un registro
PUT http://localhost:8080/v1/<recurso>/1

# Eliminar un registro
DELETE http://localhost:8080/v1/<recurso>/1
```

---

## CORS

El servidor tiene CORS habilitado para todos los orígenes con los siguientes métodos permitidos:

```
GET, POST, PUT, PATCH, OPTIONS, DELETE
```

---

## Pruebas Unitarias

```bash
# En proceso
```

---

## Estado CI

| Develop | Release 0.0.1 | Master |
|---|---|---|
| En proceso | En proceso | En proceso |