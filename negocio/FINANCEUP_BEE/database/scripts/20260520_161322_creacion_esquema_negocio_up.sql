-- ============================================================
-- SCHEMA: negocio (INDEPENDIENTE)
-- Base de datos temporal para desarrollo con Beego
-- NOTA: Tipo TIME convertido a VARCHAR para evitar errores
-- ============================================================

CREATE SCHEMA IF NOT EXISTS "negocio";

-- Función global de auditoría
CREATE OR REPLACE FUNCTION fn_update_fecha_modificacion()
RETURNS TRIGGER AS $$
BEGIN
    NEW.fecha_modificacion = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- ============================================================
-- TABLA: negocio.banco
-- ============================================================
CREATE TABLE "negocio"."banco" (
    "id_banco"            SERIAL PRIMARY KEY,
    "nombre_banco"        VARCHAR(150) NOT NULL,
    "ciudad"              VARCHAR(100),
    "contacto"            VARCHAR(100),
    "telefono"            VARCHAR(20),
    "email"               VARCHAR(150),
    "comision_porcentaje" DECIMAL(5,2),
    "url_logo"            VARCHAR(500),
    "descripcion"         TEXT,
    "sitio_web"           VARCHAR(200),
    "estado"              VARCHAR(20)  NOT NULL DEFAULT 'activo'
                              CHECK (estado IN ('activo','inactivo','suspendido')),
    "activo"              BOOLEAN      NOT NULL DEFAULT true,
    "fecha_creacion"      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_banco_mod
    BEFORE UPDATE ON "negocio"."banco"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- ============================================================
-- TABLA: negocio.producto_crediticio
-- ============================================================
CREATE TABLE "negocio"."producto_crediticio" (
    "id_producto"         SERIAL PRIMARY KEY,
    "id_banco"            INT          NOT NULL
                              REFERENCES "negocio"."banco"("id_banco") ON DELETE RESTRICT,
    "nombre_producto"     VARCHAR(150) NOT NULL,
    "descripcion"         TEXT,
    "monto_minimo"        DECIMAL(12,2),
    "monto_maximo"        DECIMAL(12,2),
    "tasa_minima"         DECIMAL(5,2),
    "tasa_maxima"         DECIMAL(5,2),
    "plazo_minimo"        INT,
    "plazo_maximo"        INT,
    "requisitos"          TEXT,
    "activo"              BOOLEAN      NOT NULL DEFAULT true,
    "fecha_creacion"      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_producto_crediticio_mod
    BEFORE UPDATE ON "negocio"."producto_crediticio"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- ============================================================
-- TABLA: negocio.asesor_bancario
-- ============================================================
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

-- ============================================================
-- TABLA: negocio.contacto_asesor
-- CORRECCION: TIME convertido a VARCHAR (formato HH:MM)
-- ============================================================
CREATE TABLE "negocio"."contacto_asesor" (
    "id_contacto"         SERIAL PRIMARY KEY,
    "id_asesor"           INT         NOT NULL
                              REFERENCES "negocio"."asesor_bancario"("id_asesor") ON DELETE RESTRICT,
    "whatsapp"            VARCHAR(20),
    "email"               VARCHAR(150),
    "telefono"            VARCHAR(20),
    "disponible_desde"    VARCHAR(5),
    "disponible_hasta"    VARCHAR(5),
    "dias_disponibles"    VARCHAR(100),
    "activo"              BOOLEAN     NOT NULL DEFAULT true,
    "fecha_creacion"      TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_contacto_asesor_mod
    BEFORE UPDATE ON "negocio"."contacto_asesor"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- ============================================================
-- TABLA: negocio.lead
-- ============================================================
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

-- ============================================================
-- TABLA: negocio.conversacion_usuario_asesor
-- ============================================================
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

-- ============================================================
-- TABLA: negocio.credito_desembolsado
-- ============================================================
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

-- ============================================================
-- TABLA: negocio.transaccion_comision
-- ============================================================
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
-- ÍNDICES
-- ============================================================
CREATE INDEX idx_lead_usuario       ON "negocio"."lead"("id_usuario");
CREATE INDEX idx_lead_estado        ON "negocio"."lead"("estado_lead");
CREATE INDEX idx_credito_usuario    ON "negocio"."credito_desembolsado"("id_usuario");
CREATE INDEX idx_credito_estado     ON "negocio"."credito_desembolsado"("estado_credito");
CREATE INDEX idx_banco_activo       ON "negocio"."banco"("activo");
CREATE INDEX idx_producto_banco     ON "negocio"."producto_crediticio"("id_banco");

-- ============================================================
-- DATOS DE PRUEBA
-- ============================================================

-- negocio.banco
INSERT INTO "negocio"."banco"
    ("nombre_banco", "ciudad", "contacto", "telefono", "email",
     "comision_porcentaje", "url_logo", "descripcion", "sitio_web", "estado", "activo")
VALUES
('Banco Colombiano',    'Bogota',      'Juan Gomez',   '6015551234', 'contacto@bancocol.com',    2.50, 'https://example.com/logo1.png', 'Banco con amplia cobertura nacional',  'www.bancocol.com',    'activo', true),
('Banco Metropolitano', 'Medellin',    'Maria Lopez',  '5745551234', 'contacto@bancometro.com',  2.75, 'https://example.com/logo2.png', 'Banco especializado en creditos',      'www.bancometro.com',  'activo', true),
('Banco del Oriente',   'Bucaramanga', 'Carlos Ruiz',  '7685551234', 'contacto@bancoriente.com', 2.25, 'https://example.com/logo3.png', 'Banco regional con buenos servicios', 'www.bancoriente.com', 'activo', true),
('Banco Pacifico',      'Cali',        'Pedro Diaz',   '3105551234', 'contacto@bancopac.com',    2.60, 'https://example.com/logo4.png', 'Banco del sector occidental',         'www.bancopac.com',    'activo', true);

-- negocio.producto_crediticio
INSERT INTO "negocio"."producto_crediticio"
    ("id_banco", "nombre_producto", "descripcion",
     "monto_minimo", "monto_maximo", "tasa_minima", "tasa_maxima",
     "plazo_minimo", "plazo_maximo", "requisitos", "activo")
VALUES
(1, 'Credito Personal',    'Credito para gastos personales',  1000000,    50000000,  10.00, 20.00, 12,  84,  'Cedula, comprobante ingresos',         true),
(1, 'Credito Hipotecario', 'Credito para compra de vivienda', 100000000, 1000000000,  8.00, 12.00, 120, 360, 'Cedula, avaluo, comprobante ingresos', true),
(2, 'Microcredito',        'Credito para pequenos negocios',  500000,     10000000,  15.00, 25.00,  6,  60,  'Cedula, plan de negocio',              true),
(3, 'Credito de Vehiculo', 'Financiamiento de vehiculos',     10000000,  150000000,   9.00, 18.00, 24,  84,  'Cedula, documento vehiculo',           true);

-- negocio.asesor_bancario
INSERT INTO "negocio"."asesor_bancario"
    ("id_banco", "nombre", "apellido", "email", "telefono", "especialidad", "activo")
VALUES
(1, 'Roberto',  'Sanchez',    'roberto.sanchez@bancocol.com',     '3001112222', 'Creditos Personales', true),
(1, 'Diana',    'Valenzuela', 'diana.valenzuela@bancocol.com',     '3009998888', 'Hipotecarios',        true),
(2, 'Fernando', 'Castillo',   'fernando.castillo@bancometro.com',  '3107776666', 'Microcreditos',       true),
(3, 'Claudia',  'Morales',    'claudia.morales@bancoriente.com',   '3104445555', 'Vehiculos',           true);

-- negocio.contacto_asesor
INSERT INTO "negocio"."contacto_asesor"
    ("id_asesor", "whatsapp", "email", "telefono",
     "disponible_desde", "disponible_hasta", "dias_disponibles", "activo")
VALUES
(1, '3001112222', 'roberto.sanchez@bancocol.com',     '3001112222', '08:00', '18:00', 'Lunes a Viernes', true),
(2, '3009998888', 'diana.valenzuela@bancocol.com',    '3009998888', '09:00', '17:00', 'Lunes a Viernes', true),
(3, '3107776666', 'fernando.castillo@bancometro.com', '3107776666', '08:00', '20:00', 'Lunes a Sabado',  true),
(4, '3104445555', 'claudia.morales@bancoriente.com',  '3104445555', '07:00', '19:00', 'Lunes a Viernes', true);

-- negocio.lead
INSERT INTO "negocio"."lead"
    ("id_usuario", "id_producto", "id_asesor", "tipo_credito",
     "monto_interes", "plazo_interes", "estado_lead",
     "fecha_generacion", "fecha_contacto", "observaciones", "activo")
VALUES
(1, 1, 1, 'Personal',     5000000,   36, 'aprobado',   CURRENT_TIMESTAMP - INTERVAL '60 days', CURRENT_TIMESTAMP - INTERVAL '50 days', 'Cliente solvente, aprobado sin inconvenientes', true),
(2, 2, 2, 'Hipotecario',  200000000, 240, 'en_proceso', CURRENT_TIMESTAMP - INTERVAL '30 days', CURRENT_TIMESTAMP - INTERVAL '25 days', 'Pendiente evaluacion de inmueble',              true),
(3, 3, 3, 'Microcredito', 5000000,   36, 'aprobado',   CURRENT_TIMESTAMP - INTERVAL '45 days', CURRENT_TIMESTAMP - INTERVAL '40 days', 'Negocio informal, requiere seguimiento',        true),
(4, 4, 4, 'Vehiculo',     50000000,  60, 'contactado', CURRENT_TIMESTAMP - INTERVAL '15 days', CURRENT_TIMESTAMP - INTERVAL '10 days', 'Interesado en financiamiento de camioneta',     true),
(5, 1, 1, 'Personal',     3000000,   24, 'nuevo',      CURRENT_TIMESTAMP - INTERVAL '5 days',  NULL,                                   'Lead recien generado',                          true);

-- negocio.conversacion_usuario_asesor
INSERT INTO "negocio"."conversacion_usuario_asesor"
    ("id_lead", "id_usuario", "id_asesor", "tipo_contacto",
     "asunto", "contenido", "fecha_mensaje", "activo")
VALUES
(1, 1, 1, 'email',    'Solicitud de Credito Aprobada', 'Le informamos que su solicitud fue aprobada.',         CURRENT_TIMESTAMP - INTERVAL '50 days', true),
(1, 1, 1, 'whatsapp', 'Confirmacion de Desembolso',    'El dinero sera depositado en 2 dias habiles.',         CURRENT_TIMESTAMP - INTERVAL '45 days', true),
(2, 2, 2, 'telefono', 'Solicitud de Informacion',      'Llamada para recabar informacion sobre su propiedad.', CURRENT_TIMESTAMP - INTERVAL '25 days', true),
(3, 3, 3, 'email',    'Plan de Acompanamiento',        'Iniciamos plan de acompanamiento para su negocio.',    CURRENT_TIMESTAMP - INTERVAL '40 days', true),
(4, 4, 4, 'whatsapp', 'Consulta sobre Vehiculo',       'Informacion sobre financiamiento de su vehiculo.',     CURRENT_TIMESTAMP - INTERVAL '10 days', true);

-- negocio.credito_desembolsado
INSERT INTO "negocio"."credito_desembolsado"
    ("id_lead", "id_usuario", "id_producto", "id_banco",
     "numero_credito", "monto_aprobado", "tasa_interes_final", "plazo_meses",
     "fecha_aprobacion", "fecha_desembolso", "estado_credito", "saldo_actual", "activo")
VALUES
(1, 1, 1, 1, 'CRED-2024-001',   5000000, 15.50,  36, CURRENT_DATE - INTERVAL '50 days', CURRENT_DATE - INTERVAL '48 days', 'activo',   4800000, true),
(2, 2, 2, 1, 'CRED-2024-002', 200000000, 10.00, 240, CURRENT_DATE - INTERVAL '20 days', CURRENT_DATE - INTERVAL '18 days', 'activo', 199800000, true),
(3, 3, 3, 2, 'CRED-2024-003',   5000000, 18.75,  36, CURRENT_DATE - INTERVAL '40 days', CURRENT_DATE - INTERVAL '38 days', 'activo',   4700000, true),
(4, 4, 4, 3, 'CRED-2024-004',  50000000, 12.50,  60, CURRENT_DATE - INTERVAL '5 days',  NULL,                              'activo',  50000000, true);

-- negocio.transaccion_comision
INSERT INTO "negocio"."transaccion_comision"
    ("id_credito", "id_banco", "monto_comision", "porcentaje_aplicado",
     "fecha_transaccion", "estado", "referencia_pago", "activo")
VALUES
(1, 1,  125000, 2.50, CURRENT_TIMESTAMP - INTERVAL '48 days', 'pagada',    'TRANS-001-2024', true),
(2, 1, 5000000, 2.50, CURRENT_TIMESTAMP - INTERVAL '18 days', 'pendiente', 'TRANS-002-2024', true),
(3, 2,  125000, 2.50, CURRENT_TIMESTAMP - INTERVAL '38 days', 'pagada',    'TRANS-003-2024', true),
(4, 3, 1125000, 2.25, CURRENT_TIMESTAMP - INTERVAL '5 days',  'pendiente', 'TRANS-004-2024', true);