CREATE SCHEMA IF NOT EXISTS "negocio";

-- ============================================================
-- FUNCION GLOBAL DE AUDITORIA
-- ============================================================

CREATE OR REPLACE FUNCTION fn_update_fecha_modificacion()
RETURNS TRIGGER AS $$
BEGIN
    NEW.fecha_modificacion = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


-- ============================================================
-- ESQUEMA: negocio
-- Tablas: banco, producto_crediticio, asesor_bancario,
--         contacto_asesor, lead, conversacion_usuario_asesor,
--         credito_desembolsado, transaccion_comision
-- ============================================================

-- negocio.banco
CREATE TABLE "negocio"."banco" (
    "id_banco"              SERIAL PRIMARY KEY,
    "nombre_banco"          VARCHAR(100) UNIQUE NOT NULL,
    "ciudad"                VARCHAR(100),
    "contacto"              VARCHAR(100),
    "telefono"              VARCHAR(20),
    "email"                 VARCHAR(150),
    "comision_porcentaje"   DECIMAL(5,2),
    "url_logo"              VARCHAR(500),
    "descripcion"           TEXT,
    "sitio_web"             VARCHAR(300),
    "estado"                VARCHAR(20)  NOT NULL DEFAULT 'activo'
                                CHECK (estado IN ('activo','inactivo')),
    "fecha_registro"        TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
    "activo"                BOOLEAN      NOT NULL DEFAULT true,
    "fecha_creacion"        TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_banco_mod
    BEFORE UPDATE ON "negocio"."banco"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- negocio.producto_crediticio
CREATE TABLE "negocio"."producto_crediticio" (
    "id_producto"         SERIAL PRIMARY KEY,
    "id_banco"            INT           NOT NULL
                              REFERENCES "negocio"."banco"("id_banco") ON DELETE RESTRICT,
    "nombre_producto"     VARCHAR(150)  NOT NULL,
    "descripcion"         TEXT,
    "monto_minimo"        DECIMAL(12,2),
    "monto_maximo"        DECIMAL(12,2),
    "tasa_minima"         DECIMAL(5,2),
    "tasa_maxima"         DECIMAL(5,2),
    "plazo_minimo"        INT,
    "plazo_maximo"        INT,
    "requisitos"          TEXT,
    "activo"              BOOLEAN       NOT NULL DEFAULT true,
    "fecha_creacion"      TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"  TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_producto_crediticio_mod
    BEFORE UPDATE ON "negocio"."producto_crediticio"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- negocio.asesor_bancario
CREATE TABLE "negocio"."asesor_bancario" (
    "id_asesor"           SERIAL PRIMARY KEY,
    "id_banco"            INT          NOT NULL
                              REFERENCES "negocio"."banco"("id_banco") ON DELETE RESTRICT,
    "nombre"              VARCHAR(100) NOT NULL,
    "apellido"            VARCHAR(100) NOT NULL,
    "email"               VARCHAR(150) NOT NULL,
    "telefono"            VARCHAR(20),
    "especialidad"        VARCHAR(100),
    "activo"              BOOLEAN      NOT NULL DEFAULT true,
    "fecha_creacion"      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_asesor_bancario_mod
    BEFORE UPDATE ON "negocio"."asesor_bancario"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- negocio.contacto_asesor
CREATE TABLE "negocio"."contacto_asesor" (
    "id_contacto"         SERIAL PRIMARY KEY,
    "id_asesor"           INT         NOT NULL
                              REFERENCES "negocio"."asesor_bancario"("id_asesor") ON DELETE RESTRICT,
    "whatsapp"            VARCHAR(20),
    "email"               VARCHAR(150),
    "telefono"            VARCHAR(20),
    "disponible_desde"    TIME,
    "disponible_hasta"    TIME,
    "dias_disponibles"    VARCHAR(100),
    "activo"              BOOLEAN     NOT NULL DEFAULT true,
    "fecha_creacion"      TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_contacto_asesor_mod
    BEFORE UPDATE ON "negocio"."contacto_asesor"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- negocio.lead
CREATE TABLE "negocio"."lead" (
    "id_lead"             SERIAL PRIMARY KEY,
    "id_usuario"          INT           NOT NULL,
    "id_producto"         INT           NOT NULL
                              REFERENCES "negocio"."producto_crediticio"("id_producto") ON DELETE RESTRICT,
    "id_asesor"           INT
                              REFERENCES "negocio"."asesor_bancario"("id_asesor") ON DELETE RESTRICT,
    "tipo_credito"        VARCHAR(100),
    "monto_interes"       DECIMAL(12,2),
    "plazo_interes"       INT,
    "estado_lead"         VARCHAR(30)   NOT NULL DEFAULT 'nuevo'
                              CHECK (estado_lead IN ('nuevo','contactado','en_proceso','aprobado','rechazado','cancelado')),
    "fecha_generacion"    TIMESTAMP     DEFAULT CURRENT_TIMESTAMP,
    "fecha_contacto"      TIMESTAMP,
    "observaciones"       TEXT,
    "activo"              BOOLEAN       NOT NULL DEFAULT true,
    "fecha_creacion"      TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"  TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_lead_mod
    BEFORE UPDATE ON "negocio"."lead"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- negocio.conversacion_usuario_asesor
CREATE TABLE "negocio"."conversacion_usuario_asesor" (
    "id_conversacion"     SERIAL PRIMARY KEY,
    "id_lead"             INT         NOT NULL
                              REFERENCES "negocio"."lead"("id_lead") ON DELETE RESTRICT,
    "id_usuario"          INT         NOT NULL,
    "id_asesor"           INT         NOT NULL
                              REFERENCES "negocio"."asesor_bancario"("id_asesor") ON DELETE RESTRICT,
    "tipo_contacto"       VARCHAR(30) NOT NULL DEFAULT 'email'
                              CHECK (tipo_contacto IN ('email','telefono','whatsapp','presencial')),
    "asunto"              VARCHAR(200),
    "contenido"           TEXT,
    "fecha_mensaje"       TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
    "activo"              BOOLEAN     NOT NULL DEFAULT true,
    "fecha_creacion"      TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_conversacion_mod
    BEFORE UPDATE ON "negocio"."conversacion_usuario_asesor"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- negocio.credito_desembolsado
CREATE TABLE "negocio"."credito_desembolsado" (
    "id_credito"            SERIAL PRIMARY KEY,
    "id_lead"               INT           NOT NULL
                                REFERENCES "negocio"."lead"("id_lead") ON DELETE RESTRICT,
    "id_usuario"            INT           NOT NULL,
    "id_producto"           INT           NOT NULL
                                REFERENCES "negocio"."producto_crediticio"("id_producto") ON DELETE RESTRICT,
    "id_banco"              INT           NOT NULL
                                REFERENCES "negocio"."banco"("id_banco") ON DELETE RESTRICT,
    "numero_credito"        VARCHAR(50)   UNIQUE,
    "monto_aprobado"        DECIMAL(12,2) NOT NULL,
    "tasa_interes_final"    DECIMAL(5,2),
    "plazo_meses"           INT,
    "fecha_aprobacion"      DATE,
    "fecha_desembolso"      DATE,
    "estado_credito"        VARCHAR(20)   NOT NULL DEFAULT 'activo'
                                CHECK (estado_credito IN ('activo','pagado','vencido','cancelado')),
    "saldo_actual"          DECIMAL(12,2),
    "activo"                BOOLEAN       NOT NULL DEFAULT true,
    "fecha_creacion"        TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"    TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_credito_mod
    BEFORE UPDATE ON "negocio"."credito_desembolsado"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- negocio.transaccion_comision
CREATE TABLE "negocio"."transaccion_comision" (
    "id_transaccion"        SERIAL PRIMARY KEY,
    "id_credito"            INT           NOT NULL
                                REFERENCES "negocio"."credito_desembolsado"("id_credito") ON DELETE RESTRICT,
    "id_banco"              INT           NOT NULL
                                REFERENCES "negocio"."banco"("id_banco") ON DELETE RESTRICT,
    "monto_comision"        DECIMAL(12,2) NOT NULL,
    "porcentaje_aplicado"   DECIMAL(5,2),
    "fecha_transaccion"     TIMESTAMP     DEFAULT CURRENT_TIMESTAMP,
    "estado"                VARCHAR(20)   NOT NULL DEFAULT 'pendiente'
                                CHECK (estado IN ('pendiente','pagada','cancelada')),
    "referencia_pago"       VARCHAR(100),
    "activo"                BOOLEAN       NOT NULL DEFAULT true,
    "fecha_creacion"        TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"    TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_transaccion_mod
    BEFORE UPDATE ON "negocio"."transaccion_comision"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();


-- ============================================================
-- INDICES
-- ============================================================

CREATE INDEX idx_lead_usuario    ON "negocio"."lead"("id_usuario");
CREATE INDEX idx_lead_estado     ON "negocio"."lead"("estado_lead");
CREATE INDEX idx_credito_usuario ON "negocio"."credito_desembolsado"("id_usuario");
CREATE INDEX idx_credito_estado  ON "negocio"."credito_desembolsado"("estado_credito");
