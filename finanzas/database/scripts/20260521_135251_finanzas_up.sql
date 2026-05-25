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
    "id_movimiento_dinero"    INT REFERENCES "finanzas"."movimiento_ingreso_egreso"("id_movimiento_dinero") ON DELETE CASCADE,
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
    "id_usuario"            INT           NOT NULL,
    "id_movimiento_dinero"  INT           REFERENCES "finanzas"."movimiento_ingreso_egreso"("id_movimiento_dinero") ON DELETE CASCADE,
    "id_categoria"          INT           REFERENCES "finanzas"."categoria"("id_categoria") ON DELETE CASCADE,
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
    "id_movimiento_dinero"    INT REFERENCES "finanzas"."movimiento_inversion"("id_movimiento_dinero") ON DELETE CASCADE,
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
    "id_usuario"            INT           NOT NULL,
    "id_tipo_inversion"     INT           NOT NULL
                                REFERENCES "finanzas"."tipo_inversion"("id_tipo_inversion") ON DELETE CASCADE,
    "id_nivel_riesgo"       INT           NOT NULL
                                REFERENCES "finanzas"."nivel_riesgo"("id_nivel_riesgo") ON DELETE CASCADE,
    "id_movimiento_dinero"  INT           REFERENCES "finanzas"."movimiento_inversion"("id_movimiento_dinero") ON DELETE CASCADE,
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
    "id_movimiento_dinero"    INT REFERENCES "finanzas"."movimiento_meta"("id_movimiento_dinero") ON DELETE CASCADE,
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
    "id_usuario"            INT           NOT NULL,
    "id_editar_meta"        INT           NOT NULL
                                REFERENCES "finanzas"."editar_meta"("id_editar_meta") ON DELETE CASCADE,
    "id_movimiento_dinero"  INT           REFERENCES "finanzas"."movimiento_meta"("id_movimiento_dinero") ON DELETE CASCADE,
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

-- Idempotent seed for finanzas schema: inserts sample rows only if table is empty
BEGIN;

-- categoria
INSERT INTO finanzas.categoria (nombre, descripcion, activo)
SELECT 'Arriendo', 'Gastos de vivienda', true
WHERE NOT EXISTS (SELECT 1 FROM finanzas.categoria);

-- movimiento_ingreso_egreso
INSERT INTO finanzas.movimiento_ingreso_egreso (nombre, monto, es_ingreso, activo)
SELECT 'Salario Abril', 4000000, true, true
WHERE NOT EXISTS (SELECT 1 FROM finanzas.movimiento_ingreso_egreso);

-- tipo_ingreso
INSERT INTO finanzas.tipo_ingreso (id_movimiento_dinero, nombre_movimiento_pago, descripcion, activo)
SELECT m.id_movimiento_dinero, 'Salario', 'Ingreso por salario', true
FROM finanzas.movimiento_ingreso_egreso m
WHERE NOT EXISTS (SELECT 1 FROM finanzas.tipo_ingreso) LIMIT 1;

-- tipo_inversion
INSERT INTO finanzas.tipo_inversion (nombre, descripcion, activo)
SELECT 'Acciones', 'Inversion en acciones', true
WHERE NOT EXISTS (SELECT 1 FROM finanzas.tipo_inversion);

-- nivel_riesgo
INSERT INTO finanzas.nivel_riesgo (nombre, activo)
SELECT 'Medio', true
WHERE NOT EXISTS (SELECT 1 FROM finanzas.nivel_riesgo);

-- movimiento_inversion
INSERT INTO finanzas.movimiento_inversion (nombre, monto, es_ingreso, activo)
SELECT 'Inversion Inicial', 100000, false, true
WHERE NOT EXISTS (SELECT 1 FROM finanzas.movimiento_inversion);

-- tipo_ingreso_inversion
INSERT INTO finanzas.tipo_ingreso_inversion (id_movimiento_dinero, nombre_movimiento_pago, descripcion, activo)
SELECT m.id_movimiento_dinero, 'Rendimiento', 'Rendimiento inversion', true
FROM finanzas.movimiento_inversion m
WHERE NOT EXISTS (SELECT 1 FROM finanzas.tipo_ingreso_inversion) LIMIT 1;

-- editar_meta
INSERT INTO finanzas.editar_meta (nombre, monto_actual, monto_objetivo, ahorro_mensual, fecha_objetivo, descripcion, activo)
SELECT 'Fondo Emergencia', 1000000, 5000000, 50000, CURRENT_DATE + INTERVAL '180 days', 'Ahorro inicial', true
WHERE NOT EXISTS (SELECT 1 FROM finanzas.editar_meta);

-- movimiento_meta
INSERT INTO finanzas.movimiento_meta (nombre, monto, es_ingreso, activo)
SELECT 'Aporte Meta', 50000, true, true
WHERE NOT EXISTS (SELECT 1 FROM finanzas.movimiento_meta);

-- tipo_ingreso_meta
INSERT INTO finanzas.tipo_ingreso_meta (id_movimiento_dinero, nombre_movimiento_pago, descripcion, activo)
SELECT m.id_movimiento_dinero, 'Aporte Regular', 'Aporte a meta', true
FROM finanzas.movimiento_meta m
WHERE NOT EXISTS (SELECT 1 FROM finanzas.tipo_ingreso_meta) LIMIT 1;

-- finanzas (budget) - requires usuario and categoria; use defaults if available
INSERT INTO finanzas.finanzas (id_usuario, id_movimiento_dinero, id_categoria, monto_presupuesto, gasto, disponible, fecha, activo)
SELECT 1, (SELECT id_movimiento_dinero FROM finanzas.movimiento_ingreso_egreso LIMIT 1), (SELECT id_categoria FROM finanzas.categoria LIMIT 1), 1000000, 200000, 800000, CURRENT_DATE, true
WHERE NOT EXISTS (SELECT 1 FROM finanzas.finanzas);

-- inversion sample data
INSERT INTO finanzas.inversion (id_usuario, id_tipo_inversion, id_nivel_riesgo, id_movimiento_dinero, nombre, monto, rentabilidad, fecha_inicio, fecha_fin, activo)
SELECT 1,
       (SELECT id_tipo_inversion FROM finanzas.tipo_inversion LIMIT 1),
       (SELECT id_nivel_riesgo FROM finanzas.nivel_riesgo LIMIT 1),
       (SELECT id_movimiento_dinero FROM finanzas.movimiento_inversion LIMIT 1),
       'Inversion Semilla', 500000, 7.25,
       CURRENT_DATE - INTERVAL '30 days', CURRENT_DATE + INTERVAL '335 days', true
WHERE NOT EXISTS (SELECT 1 FROM finanzas.inversion);

-- meta sample data
INSERT INTO finanzas.meta (id_usuario, id_editar_meta, id_movimiento_dinero, nombre, descripcion, monto_objetivo, monto_actual, fecha_limite, color, activo)
SELECT 1,
       (SELECT id_editar_meta FROM finanzas.editar_meta LIMIT 1),
       (SELECT id_movimiento_dinero FROM finanzas.movimiento_meta LIMIT 1),
       'Meta Semilla', 'Meta de ejemplo', 5000000, 100000, CURRENT_DATE + INTERVAL '365 days', '#FFEEAA', true
WHERE NOT EXISTS (SELECT 1 FROM finanzas.meta);

COMMIT;



