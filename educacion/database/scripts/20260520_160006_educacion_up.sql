CREATE SCHEMA IF NOT EXISTS "educacion";

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
-- ESQUEMA: educacion
-- Tablas: modulo_educativo, contenido, leccion,
--         progreso_educativo, progreso_leccion
-- ============================================================

-- educacion.modulo_educativo
CREATE TABLE "educacion"."modulo_educativo" (
    "id_modulo"           SERIAL PRIMARY KEY,
    "titulo"              VARCHAR(200) NOT NULL,
    "descripcion"         TEXT,
    "contenido"           TEXT,
    "nivel"               VARCHAR(20)  NOT NULL DEFAULT 'basico'
                              CHECK (nivel IN ('basico','intermedio','avanzado')),
    "url_thumbnail"       VARCHAR(500),
    "activo"              BOOLEAN      NOT NULL DEFAULT true,
    "fecha_creacion"      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_modulo_educativo_mod
    BEFORE UPDATE ON "educacion"."modulo_educativo"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- educacion.contenido
CREATE TABLE "educacion"."contenido" (
    "id_contenido"        SERIAL PRIMARY KEY,
    "titulo"              VARCHAR(200) NOT NULL,
    "descripcion"         TEXT,
    "duracion_minutos"    INT,
    "url_video"           VARCHAR(500),
    "activo"              BOOLEAN      NOT NULL DEFAULT true,
    "fecha_creacion"      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_contenido_mod
    BEFORE UPDATE ON "educacion"."contenido"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- educacion.leccion
CREATE TABLE "educacion"."leccion" (
    "id_leccion"          SERIAL PRIMARY KEY,
    "id_modulo"           INT          NOT NULL
                              REFERENCES "educacion"."modulo_educativo"("id_modulo") ON DELETE RESTRICT,
    "id_contenido"        INT
                              REFERENCES "educacion"."contenido"("id_contenido") ON DELETE RESTRICT,
    "titulo"              VARCHAR(200) NOT NULL,
    "descripcion"         TEXT,
    "duracion_minutos"    INT,
    "url_video"           VARCHAR(500),
    "numero_leccion"      INT,
    "activo"              BOOLEAN      NOT NULL DEFAULT true,
    "fecha_creacion"      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_leccion_mod
    BEFORE UPDATE ON "educacion"."leccion"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- educacion.progreso_educativo
CREATE TABLE "educacion"."progreso_educativo" (
    "id_progreso"             SERIAL PRIMARY KEY,
    "id_usuario"              INT       NOT NULL,
    "id_modulo"               INT       NOT NULL
                                  REFERENCES "educacion"."modulo_educativo"("id_modulo") ON DELETE RESTRICT,
    "porcentaje_completado"   INT       DEFAULT 0,
    "fecha_inicio"            TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "fecha_completado"        TIMESTAMP,
    "calificacion"            INT,
    "activo"                  BOOLEAN   NOT NULL DEFAULT true,
    "fecha_creacion"          TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE("id_usuario", "id_modulo")
);

CREATE TRIGGER trg_progreso_educativo_mod
    BEFORE UPDATE ON "educacion"."progreso_educativo"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();

-- educacion.progreso_leccion
CREATE TABLE "educacion"."progreso_leccion" (
    "id_progreso_leccion" SERIAL PRIMARY KEY,
    "id_usuario"          INT       NOT NULL,
    "id_leccion"          INT       NOT NULL
                              REFERENCES "educacion"."leccion"("id_leccion") ON DELETE RESTRICT,
    "completado"          BOOLEAN   DEFAULT false,
    "fecha_inicio"        TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "fecha_completado"    TIMESTAMP,
    "activo"              BOOLEAN   NOT NULL DEFAULT true,
    "fecha_creacion"      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE("id_usuario", "id_leccion")
);

CREATE TRIGGER trg_progreso_leccion_mod
    BEFORE UPDATE ON "educacion"."progreso_leccion"
    FOR EACH ROW EXECUTE FUNCTION fn_update_fecha_modificacion();


-- ============================================================
-- INDICES
-- ============================================================

CREATE INDEX idx_progreso_usuario ON "educacion"."progreso_educativo"("id_usuario");
