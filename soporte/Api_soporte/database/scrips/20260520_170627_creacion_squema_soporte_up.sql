
CREATE SCHEMA IF NOT EXISTS "soporte";


CREATE OR REPLACE FUNCTION fn_update_fecha_modificacion()
RETURNS TRIGGER AS $$
BEGIN
    NEW.fecha_modificacion = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


-- TABLA: soporte.estado_pqr
CREATE TABLE "soporte"."estado_pqr" (
    "id_estado"           SERIAL PRIMARY KEY,
    "nombre"              VARCHAR(50)  NOT NULL,
    "descripcion"         VARCHAR(255),
    "activo"              BOOLEAN      NOT NULL DEFAULT true,
    "fecha_creacion"      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_estado_pqr_mod
    BEFORE UPDATE ON "soporte"."estado_pqr"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();



-- TABLA: soporte.pqr
CREATE TABLE "soporte"."pqr" (
    "id_pqr"              SERIAL PRIMARY KEY,
    "id_usuario"          INT       NOT NULL,
    "descripcion"         TEXT      NOT NULL,
    "id_estado"           INT       NOT NULL
                              REFERENCES "soporte"."estado_pqr"("id_estado") ON DELETE RESTRICT,
    "activo"              BOOLEAN   NOT NULL DEFAULT true,
    "fecha_creacion"      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_pqr_mod
    BEFORE UPDATE ON "soporte"."pqr"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();



-- TABLA: soporte.adjunto
CREATE TABLE "soporte"."adjunto" (
    "id_adjunto"          SERIAL PRIMARY KEY,
    "id_pqr"              INT         NOT NULL
                              REFERENCES "soporte"."pqr"("id_pqr") ON DELETE RESTRICT,
    "nombre_archivo"      VARCHAR(255) NOT NULL,
    "ruta_archivo"        VARCHAR(500),
    "tipo_mime"           VARCHAR(100),
    "tamano_bytes"        INT,
    "activo"              BOOLEAN      NOT NULL DEFAULT true,
    "fecha_creacion"      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_adjunto_mod
    BEFORE UPDATE ON "soporte"."adjunto"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();



-- TABLA: soporte.registro_actividad
CREATE TABLE "soporte"."registro_actividad" (
    "id_registro"         SERIAL PRIMARY KEY,
    "id_usuario"          INT         NOT NULL,
    "tipo_actividad"      VARCHAR(50) NOT NULL,
    "descripcion"         TEXT,
    "entidad_afectada"    VARCHAR(100),
    "fecha_actividad"     TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
    "fecha_creacion"      TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP
);



-- ÍNDICES
CREATE INDEX idx_pqr_usuario        ON "soporte"."pqr"("id_usuario");
CREATE INDEX idx_adjunto_pqr        ON "soporte"."adjunto"("id_pqr");
CREATE INDEX idx_registro_usuario   ON "soporte"."registro_actividad"("id_usuario");