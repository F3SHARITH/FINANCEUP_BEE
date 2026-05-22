-- ============================================================
-- SCHEMA: auth (INDEPENDIENTE)
-- Base de datos temporal para desarrollo con Beego
-- ============================================================

CREATE SCHEMA IF NOT EXISTS "auth";

-- Función global de auditoría
CREATE OR REPLACE FUNCTION fn_update_fecha_modificacion()
RETURNS TRIGGER AS $$
BEGIN
    NEW.fecha_modificacion = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- ============================================================
-- TABLA: auth.rol
-- ============================================================
CREATE TABLE "auth"."rol" (
    "id_rol"              SERIAL PRIMARY KEY,
    "nombre_rol"          VARCHAR(100) UNIQUE NOT NULL,
    "descripcion"         TEXT,
    "activo"              BOOLEAN     NOT NULL DEFAULT true,
    "fecha_creacion"      TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_rol_mod
    BEFORE UPDATE ON "auth"."rol"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- ============================================================
-- TABLA: auth.tipo_documento
-- ============================================================
CREATE TABLE "auth"."tipo_documento" (
    "id_tipo_documento"   SERIAL PRIMARY KEY,
    "nombre"              VARCHAR(50)  NOT NULL,
    "codigo"              VARCHAR(10)  UNIQUE NOT NULL,
    "activo"              BOOLEAN      NOT NULL DEFAULT true,
    "fecha_creacion"      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_tipo_documento_mod
    BEFORE UPDATE ON "auth"."tipo_documento"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- ============================================================
-- TABLA: auth.usuario
-- ============================================================
CREATE TABLE "auth"."usuario" (
    "id_usuario"          SERIAL PRIMARY KEY,
    "tipo_documento"      INT          REFERENCES "auth"."tipo_documento"("id_tipo_documento") ON DELETE RESTRICT,
    "nombre"              VARCHAR(100) NOT NULL,
    "apellido"            VARCHAR(100) NOT NULL,
    "email"               VARCHAR(150) UNIQUE NOT NULL,
    "telefono"            VARCHAR(20),
    "cedula"              VARCHAR(50)  UNIQUE,
    "ciudad"              VARCHAR(100),
    "estado"              VARCHAR(20)  NOT NULL DEFAULT 'activo'
                              CHECK (estado IN ('activo','inactivo','suspendido')),
    "fecha_registro"      TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
    "fecha_ultima_sesion" TIMESTAMP,
    "activo"              BOOLEAN      NOT NULL DEFAULT true,
    "fecha_creacion"      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_usuario_mod
    BEFORE UPDATE ON "auth"."usuario"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- ============================================================
-- TABLA: auth.credencial
-- ============================================================
CREATE TABLE "auth"."credencial" (
    "id_credencial"         SERIAL PRIMARY KEY,
    "id_usuario"            INT         UNIQUE NOT NULL
                                REFERENCES "auth"."usuario"("id_usuario") ON DELETE RESTRICT,
    "contrasena_hash"       VARCHAR(255) NOT NULL,
    "salt"                  VARCHAR(100) NOT NULL,
    "algoritmo"             VARCHAR(20)  NOT NULL DEFAULT 'bcrypt'
                                CHECK (algoritmo IN ('bcrypt','argon2','sha256')),
    "fecha_actualizacion"   TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
    "fecha_ultimo_cambio"   TIMESTAMP,
    "intentos_fallidos"     INT          DEFAULT 0,
    "bloqueado_hasta"       TIMESTAMP,
    "requiere_cambio"       BOOLEAN      DEFAULT false,
    "activo"                BOOLEAN      NOT NULL DEFAULT true,
    "fecha_creacion"        TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_credencial_mod
    BEFORE UPDATE ON "auth"."credencial"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- ============================================================
-- TABLA: auth.usuario_rol
-- ============================================================
CREATE TABLE "auth"."usuario_rol" (
    "id_usuario_rol"      SERIAL PRIMARY KEY,
    "id_usuario"          INT       NOT NULL
                              REFERENCES "auth"."usuario"("id_usuario") ON DELETE RESTRICT,
    "id_rol"              INT       NOT NULL
                              REFERENCES "auth"."rol"("id_rol") ON DELETE RESTRICT,
    "fecha_asignacion"    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "activo"              BOOLEAN   NOT NULL DEFAULT true,
    "fecha_creacion"      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE("id_usuario", "id_rol")
);

CREATE TRIGGER trg_usuario_rol_mod
    BEFORE UPDATE ON "auth"."usuario_rol"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- ============================================================
-- TABLA: auth.auditoria_login
-- ============================================================
CREATE TABLE "auth"."auditoria_login" (
    "id_auditoria"        SERIAL PRIMARY KEY,
    "id_usuario"          INT         NOT NULL
                              REFERENCES "auth"."usuario"("id_usuario") ON DELETE RESTRICT,
    "tipo_evento"         VARCHAR(30) NOT NULL DEFAULT 'login'
                              CHECK (tipo_evento IN ('login','logout','cambio_contrasena','acceso_denegado')),
    "ip_address"          VARCHAR(45),
    "navegador"           VARCHAR(200),
    "fecha_evento"        TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
    "estado_evento"       VARCHAR(20) NOT NULL DEFAULT 'exitoso'
                              CHECK (estado_evento IN ('exitoso','fallido')),
    "fecha_creacion"      TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP
);