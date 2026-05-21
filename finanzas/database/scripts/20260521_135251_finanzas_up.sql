CREATE SCHEMA IF NOT EXISTS "finanzas";

-- ============================================================
-- FUNCIÓN GLOBAL DE AUDITORÍA
-- ============================================================

CREATE OR REPLACE FUNCTION fn_update_fecha_modificacion()
RETURNS TRIGGER AS $$
BEGIN
    NEW.fecha_modificacion = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


-- ============================================================
-- ESQUEMA: finanzas
-- Tablas: categoria, movimiento_ingreso_egreso, tipo_ingreso,
--         finanzas, tipo_inversion, nivel_riesgo,
--         movimiento_inversion, tipo_ingreso_inversion,
--         inversion, editar_meta, movimiento_meta,
--         tipo_ingreso_meta, meta
-- ============================================================

-- finanzas.categoria
CREATE TABLE IF NOT EXISTS "finanzas"."categoria" (
    "id_categoria"        SERIAL PRIMARY KEY,
    "nombre"              VARCHAR(100) NOT NULL,
    "descripcion"         TEXT,
    "activo"              BOOLEAN      NOT NULL DEFAULT true,
    "fecha_creacion"      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_categoria_mod
    BEFORE UPDATE ON "finanzas"."categoria"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- finanzas.movimiento_ingreso_egreso
CREATE TABLE IF NOT EXISTS "finanzas"."movimiento_ingreso_egreso" (
    "id_movimiento_dinero"  SERIAL PRIMARY KEY,
    "nombre"                VARCHAR(120)  NOT NULL,
    "monto"                 DECIMAL(18,2) NOT NULL,
    "es_ingreso"            BOOLEAN       NOT NULL,
    "activo"                BOOLEAN       NOT NULL DEFAULT true,
    "fecha_creacion"        TIMESTAMP     DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"    TIMESTAMP     DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_fin_movimiento_mod
    BEFORE UPDATE ON "finanzas"."movimiento_ingreso_egreso"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- finanzas.tipo_ingreso
CREATE TABLE IF NOT EXISTS "finanzas"."tipo_ingreso" (
    "id_tipo_ingreso"         SERIAL PRIMARY KEY,
    "id_movimiento_dinero"    INT REFERENCES "finanzas"."movimiento_ingreso_egreso"("id_movimiento_dinero") ON DELETE RESTRICT,
    "nombre_movimiento_pago"  VARCHAR(120) NOT NULL,
    "descripcion"             VARCHAR(150),
    "activo"                  BOOLEAN      NOT NULL DEFAULT true,
    "fecha_creacion"          TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"      TIMESTAMP    DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_fin_tipo_ingreso_mod
    BEFORE UPDATE ON "finanzas"."tipo_ingreso"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- finanzas.finanzas
CREATE TABLE IF NOT EXISTS "finanzas"."finanzas" (
    "id_finanzas"           SERIAL PRIMARY KEY,
    "id_usuario"            INT           NOT NULL
                                REFERENCES "auth"."usuario"("id_usuario") ON DELETE RESTRICT,
    "id_movimiento_dinero"  INT           REFERENCES "finanzas"."movimiento_ingreso_egreso"("id_movimiento_dinero") ON DELETE RESTRICT,
    "id_categoria"          INT           REFERENCES "finanzas"."categoria"("id_categoria") ON DELETE RESTRICT,
    "monto_presupuesto"     DECIMAL(12,2),
    "gasto"                 DECIMAL(12,2),
    "disponible"            DECIMAL(12,2),
    "fecha"                 DATE          DEFAULT CURRENT_DATE,
    "activo"                BOOLEAN       NOT NULL DEFAULT true,
    "fecha_creacion"        TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"    TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_finanzas_mod
    BEFORE UPDATE ON "finanzas"."finanzas"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- finanzas.tipo_inversion
CREATE TABLE IF NOT EXISTS "finanzas"."tipo_inversion" (
    "id_tipo_inversion"   SERIAL PRIMARY KEY,
    "nombre"              VARCHAR(60)  NOT NULL,
    "descripcion"         VARCHAR(150),
    "activo"              BOOLEAN      NOT NULL DEFAULT true,
    "fecha_creacion"      TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"  TIMESTAMP    DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_tipo_inversion_mod
    BEFORE UPDATE ON "finanzas"."tipo_inversion"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- finanzas.nivel_riesgo
CREATE TABLE IF NOT EXISTS "finanzas"."nivel_riesgo" (
    "id_nivel_riesgo"     SERIAL PRIMARY KEY,
    "nombre"              VARCHAR(40) NOT NULL,
    "activo"              BOOLEAN     NOT NULL DEFAULT true,
    "fecha_creacion"      TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"  TIMESTAMP   DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_nivel_riesgo_mod
    BEFORE UPDATE ON "finanzas"."nivel_riesgo"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- finanzas.movimiento_inversion
CREATE TABLE IF NOT EXISTS "finanzas"."movimiento_inversion" (
    "id_movimiento_dinero"  SERIAL PRIMARY KEY,
    "nombre"                VARCHAR(120)  NOT NULL,
    "monto"                 DECIMAL(18,2) NOT NULL,
    "es_ingreso"            BOOLEAN       NOT NULL,
    "activo"                BOOLEAN       NOT NULL DEFAULT true,
    "fecha_creacion"        TIMESTAMP     DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"    TIMESTAMP     DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_inv_movimiento_mod
    BEFORE UPDATE ON "finanzas"."movimiento_inversion"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- finanzas.tipo_ingreso_inversion
CREATE TABLE IF NOT EXISTS "finanzas"."tipo_ingreso_inversion" (
    "id_tipo_ingreso"         SERIAL PRIMARY KEY,
    "id_movimiento_dinero"    INT REFERENCES "finanzas"."movimiento_inversion"("id_movimiento_dinero") ON DELETE RESTRICT,
    "nombre_movimiento_pago"  VARCHAR(120) NOT NULL,
    "descripcion"             VARCHAR(150),
    "activo"                  BOOLEAN      NOT NULL DEFAULT true,
    "fecha_creacion"          TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"      TIMESTAMP    DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_inv_tipo_ingreso_mod
    BEFORE UPDATE ON "finanzas"."tipo_ingreso_inversion"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- finanzas.inversion
CREATE TABLE IF NOT EXISTS "finanzas"."inversion" (
    "id_inversion"          SERIAL PRIMARY KEY,
    "id_usuario"            INT           NOT NULL
                                REFERENCES "auth"."usuario"("id_usuario") ON DELETE RESTRICT,
    "id_tipo_inversion"     INT           NOT NULL
                                REFERENCES "finanzas"."tipo_inversion"("id_tipo_inversion") ON DELETE RESTRICT,
    "id_nivel_riesgo"       INT           NOT NULL
                                REFERENCES "finanzas"."nivel_riesgo"("id_nivel_riesgo") ON DELETE RESTRICT,
    "id_movimiento_dinero"  INT           REFERENCES "finanzas"."movimiento_inversion"("id_movimiento_dinero") ON DELETE RESTRICT,
    "nombre"                VARCHAR(120),
    "monto"                 DECIMAL(18,2),
    "rentabilidad"          DECIMAL(8,4),
    "fecha_inicio"          DATE,
    "fecha_fin"             DATE,
    "activo"                BOOLEAN       NOT NULL DEFAULT true,
    "fecha_creacion"        TIMESTAMP     DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"    TIMESTAMP     DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_inversion_mod
    BEFORE UPDATE ON "finanzas"."inversion"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- finanzas.editar_meta
CREATE TABLE IF NOT EXISTS "finanzas"."editar_meta" (
    "id_editar_meta"      SERIAL PRIMARY KEY,
    "nombre"              VARCHAR(40),
    "monto_actual"        DECIMAL(18,2),
    "monto_objetivo"      DECIMAL(18,2) NOT NULL,
    "ahorro_mensual"      DECIMAL(18,2),
    "fecha_objetivo"      DATE          DEFAULT CURRENT_DATE,
    "descripcion"         VARCHAR(100),
    "activo"              BOOLEAN       NOT NULL DEFAULT true,
    "fecha_creacion"      TIMESTAMP     DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"  TIMESTAMP     DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_editar_meta_mod
    BEFORE UPDATE ON "finanzas"."editar_meta"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- finanzas.movimiento_meta
CREATE TABLE IF NOT EXISTS "finanzas"."movimiento_meta" (
    "id_movimiento_dinero"  SERIAL PRIMARY KEY,
    "nombre"                VARCHAR(120)  NOT NULL,
    "monto"                 DECIMAL(18,2) NOT NULL,
    "es_ingreso"            BOOLEAN       NOT NULL,
    "activo"                BOOLEAN       NOT NULL DEFAULT true,
    "fecha_creacion"        TIMESTAMP     DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"    TIMESTAMP     DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_meta_movimiento_mod
    BEFORE UPDATE ON "finanzas"."movimiento_meta"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- finanzas.tipo_ingreso_meta
CREATE TABLE IF NOT EXISTS "finanzas"."tipo_ingreso_meta" (
    "id_tipo_ingreso"         SERIAL PRIMARY KEY,
    "id_movimiento_dinero"    INT REFERENCES "finanzas"."movimiento_meta"("id_movimiento_dinero") ON DELETE RESTRICT,
    "nombre_movimiento_pago"  VARCHAR(120) NOT NULL,
    "descripcion"             VARCHAR(150),
    "activo"                  BOOLEAN      NOT NULL DEFAULT true,
    "fecha_creacion"          TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"      TIMESTAMP    DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_meta_tipo_ingreso_mod
    BEFORE UPDATE ON "finanzas"."tipo_ingreso_meta"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- finanzas.meta
CREATE TABLE IF NOT EXISTS "finanzas"."meta" (
    "id_meta"               SERIAL PRIMARY KEY,
    "id_usuario"            INT           NOT NULL
                                REFERENCES "auth"."usuario"("id_usuario") ON DELETE RESTRICT,
    "id_editar_meta"        INT           NOT NULL
                                REFERENCES "finanzas"."editar_meta"("id_editar_meta") ON DELETE RESTRICT,
    "id_movimiento_dinero"  INT           REFERENCES "finanzas"."movimiento_meta"("id_movimiento_dinero") ON DELETE RESTRICT,
    "nombre"                VARCHAR(120),
    "descripcion"           TEXT,
    "monto_objetivo"        DECIMAL(18,2),
    "monto_actual"          DECIMAL(18,2),
    "fecha_limite"          DATE,
    "color"                 VARCHAR(20),
    "activo"                BOOLEAN       NOT NULL DEFAULT true,
    "fecha_creacion"        TIMESTAMP     DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"    TIMESTAMP     DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_meta_mod
    BEFORE UPDATE ON "finanzas"."meta"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();


-- ============================================================
-- ÍNDICES
-- ============================================================

CREATE INDEX IF NOT EXISTS idx_meta_usuario       ON "finanzas"."meta"("id_usuario");
CREATE INDEX IF NOT EXISTS idx_inversion_usuario  ON "finanzas"."inversion"("id_usuario");
