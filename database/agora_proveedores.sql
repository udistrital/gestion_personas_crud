CREATE SCHEMA "proveedores";

CREATE TABLE "proveedores"."proveedor" (
  "id" serial PRIMARY KEY,
  "tipo_proveedor" integer NOT NULL,
  "descripcion_portafolio" varchar(50) NOT NULL,
  "puntaje_evaluacion" numeric(5,2),
  "clasificacion_evaluacion" integer,
  "estado_proveedor" integer,
  "consentimiento" boolean DEFAULT (true),
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "fecha_modificacion" timestamp,
  "autor_modificacion_id" integer,
  PRIMARY KEY ("id")
);

CREATE TABLE "proveedores"."contacto" (
  "id" serial PRIMARY KEY,
  "nombre_contacto" varchar(100) NOT NULL,
  "valor_especifico" varchar(100) NOT NULL,
  "proveedor_id" integer NOT NULL,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "fecha_modificacion" timestamp,
  "autor_modificacion_id" integer,
  PRIMARY KEY ("id")
);

CREATE TABLE "proveedores"."cuenta_bancaria" (
  "id" serial PRIMARY KEY,
  "tipo_cuenta_id" integer NOT NULL,
  "numero_cuenta" varchar(50) NOT NULL,
  "entidad_bancaria_id" integer NOT NULL,
  "correo_pago" varchar(320) NOT NULL,
  "tipo_cuenta_bancaria_id" boolean NOT NULL,
  "monto_capital_autorizado" numeric(20,7),
  "proveedor_id" integer NOT NULL,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "fecha_modificacion" timestamp,
  "autor_modificacion_id" integer,
  PRIMARY KEY ("id")
);

CREATE TABLE "proveedores"."documento" (
  "id" serial PRIMARY KEY,
  "tipo_documento_id" integer NOT NULL,
  "enlace" uuid NOT NULL,
  "proveedor_id" integer NOT NULL,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "fecha_modificacion" timestamp,
  "autor_modificacion_id" integer,
  PRIMARY KEY ("id")
);

CREATE TABLE "proveedores"."proveedor_natural" (
  "id" serial PRIMARY KEY,
  "pais_nacimiento_id" integer NOT NULL,
  "cargo" varchar(150),
  "medicina_prepagada" boolean NOT NULL,
  "valor_uvt_prepagada" numeric(20,7),
  "interes_vivienda_afc" numeric(5,4),
  "dependiente_hijo_mejor_edad" boolean NOT NULL,
  "dependiente_hijo_menos23_estudiando" boolean NOT NULL,
  "dependiente_hijo_mas23_discapacitado" boolean NOT NULL,
  "dependiente_conyuge" boolean NOT NULL,
  "dependiente_padre_o_hermano" boolean NOT NULL,
  "pensionado" boolean NOT NULL,
  "valor_afc" numeric(20,7),
  "responsable_iva" boolean,
  "reside_colombia" boolean,
  "pension_voluntaria" boolean,
  "valor_pension_voluntaria" numeric(20,7),
  "experiencia_laboral" integer,
  "experiencia_profesional" integer,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "fecha_modificacion" timestamp,
  "autor_modificacion_id" integer,
  PRIMARY KEY ("id")
);

CREATE TABLE "proveedores"."perfil_proveedor" (
  "id" serial PRIMARY KEY,
  "perfil_id" integer NOT NULL,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "fecha_modificacion" timestamp,
  "autor_modificacion_id" integer,
  PRIMARY KEY ("id")
);

CREATE TABLE "proveedores"."titulo_proveedor_natural" (
  "id" serial PRIMARY KEY,
  "titulo_obtenido" varchar(150) NOT NULL,
  "nucleo_basico_conocimiento_snies_id" integer NOT NULL,
  "perfil_id" integer NOT NULL,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "fecha_modificacion" timestamp,
  "autor_modificacion_id" integer,
  PRIMARY KEY ("id")
);

CREATE TABLE "proveedores"."proveedor_juridico" (
  "id" serial PRIMARY KEY,
  "procedencia_empresa_id" integer NOT NULL,
  "ciudad_origen_id" integer,
  "pais_origen_id" integer,
  "codigo_postal" varchar(16),
  "tipo_conformacion_id" integer NOT NULL,
  "exclusividad_producto" boolean NOT NULL,
  "regimen_contributivo" integer NOT NULL,
  "pyme" boolean NOT NULL,
  "registro_mercantil" boolean NOT NULL,
  "sujeto_retencion" boolean NOT NULL,
  "agente_retenedor" boolean NOT NULL,
  "responsable_ica" boolean NOT NULL,
  "responsable_iva" boolean NOT NULL,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "fecha_modificacion" timestamp,
  "autor_modificacion_id" integer,
  PRIMARY KEY ("id")
);

CREATE TABLE "proveedores"."actividad_economica_proveedor" (
  "id" serial PRIMARY KEY,
  "proveedor_id" integer NOT NULL,
  "actividad_economica_id" integer NOT NULL,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "fecha_modificacion" timestamp,
  "autor_modificacion_id" integer,
  PRIMARY KEY ("id")
);

CREATE TABLE "proveedores"."representacion" (
  "id" serial PRIMARY KEY,
  "proveedor_id" integer NOT NULL,
  "representante_id" integer NOT NULL,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "fecha_modificacion" timestamp,
  "autor_modificacion_id" integer,
  PRIMARY KEY ("id")
);

CREATE TABLE "proveedores"."persona_juridica_tipo_entidad" (
  "id" serial PRIMARY KEY,
  "nombre_perfil_entidad" varchar(50),
  "tipo_entidad_id" integer NOT NULL,
  "persona_juridica_id" varchar(50) NOT NULL,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "fecha_modificacion" timestamp,
  "autor_modificacion_id" integer,
  PRIMARY KEY ("id")
);

CREATE TABLE "proveedores"."tipo_conformacion" (
  "id" serial PRIMARY KEY,
  "nombre" varchar(50) NOT NULL,
  "descripcion" varchar(250),
  "codigo_abreviacion" varchar(20),
  "activo" boolean NOT NULL DEFAULT (true),
  "numero_orden" numeric(5,2),
  PRIMARY KEY ("id")
);

CREATE INDEX "idx_proveedor_tipo_proveedor" ON "proveedores"."proveedor" ("tipo_proveedor");

CREATE INDEX "idx_proveedor_estado_proveedor" ON "proveedores"."proveedor" ("estado_proveedor");

CREATE INDEX "idx_proveedor_activo" ON "proveedores"."proveedor" ("activo");

CREATE INDEX "idx_proveedor_fecha_creacion" ON "proveedores"."proveedor" ("fecha_creacion");

CREATE INDEX "idx_proveedor_fecha_modificacion" ON "proveedores"."proveedor" ("fecha_modificacion");

CREATE INDEX "idx_proveedor_autor_modificacion_id" ON "proveedores"."proveedor" ("autor_modificacion_id");

CREATE INDEX "idx_contacto_proveedor_id" ON "proveedores"."contacto" ("proveedor_id");

CREATE INDEX "idx_contacto_activo" ON "proveedores"."contacto" ("activo");

CREATE INDEX "idx_contacto_fecha_creacion" ON "proveedores"."contacto" ("fecha_creacion");

CREATE INDEX "idx_contacto_fecha_modificacion" ON "proveedores"."contacto" ("fecha_modificacion");

CREATE INDEX "idx_contacto_autor_modificacion_id" ON "proveedores"."contacto" ("autor_modificacion_id");

CREATE INDEX "idx_cuenta_bancaria_tipo_cuenta_id" ON "proveedores"."cuenta_bancaria" ("tipo_cuenta_id");

CREATE INDEX "idx_cuenta_bancaria_entidad_bancaria_id" ON "proveedores"."cuenta_bancaria" ("entidad_bancaria_id");

CREATE INDEX "idx_cuenta_bancaria_tipo_cuenta_bancaria_id" ON "proveedores"."cuenta_bancaria" ("tipo_cuenta_bancaria_id");

CREATE INDEX "idx_cuenta_bancaria_proveedor_id" ON "proveedores"."cuenta_bancaria" ("proveedor_id");

CREATE INDEX "idx_cuenta_bancaria_activo" ON "proveedores"."cuenta_bancaria" ("activo");

CREATE INDEX "idx_cuenta_bancaria_fecha_creacion" ON "proveedores"."cuenta_bancaria" ("fecha_creacion");

CREATE INDEX "idx_cuenta_bancaria_fecha_modificacion" ON "proveedores"."cuenta_bancaria" ("fecha_modificacion");

CREATE INDEX "idx_cuenta_bancaria_autor_modificacion_id" ON "proveedores"."cuenta_bancaria" ("autor_modificacion_id");

CREATE INDEX "idx_documento_tipo_documento_id" ON "proveedores"."documento" ("tipo_documento_id");

CREATE INDEX "idx_documento_proveedor_id" ON "proveedores"."documento" ("proveedor_id");

CREATE INDEX "idx_documento_activo" ON "proveedores"."documento" ("activo");

CREATE INDEX "idx_documento_fecha_creacion" ON "proveedores"."documento" ("fecha_creacion");

CREATE INDEX "idx_documento_fecha_modificacion" ON "proveedores"."documento" ("fecha_modificacion");

CREATE INDEX "idx_documento_autor_modificacion_id" ON "proveedores"."documento" ("autor_modificacion_id");

CREATE INDEX "idx_proveedor_natural_pais_nacimiento_id" ON "proveedores"."proveedor_natural" ("pais_nacimiento_id");

CREATE INDEX "idx_proveedor_natural_activo" ON "proveedores"."proveedor_natural" ("activo");

CREATE INDEX "idx_proveedor_natural_fecha_creacion" ON "proveedores"."proveedor_natural" ("fecha_creacion");

CREATE INDEX "idx_proveedor_natural_fecha_modificacion" ON "proveedores"."proveedor_natural" ("fecha_modificacion");

CREATE INDEX "idx_proveedor_natural_autor_modificacion_id" ON "proveedores"."proveedor_natural" ("autor_modificacion_id");

CREATE INDEX "idx_perfil_proveedor_perfil_id" ON "proveedores"."perfil_proveedor" ("perfil_id");

CREATE INDEX "idx_perfil_proveedor_activo" ON "proveedores"."perfil_proveedor" ("activo");

CREATE INDEX "idx_perfil_proveedor_fecha_creacion" ON "proveedores"."perfil_proveedor" ("fecha_creacion");

CREATE INDEX "idx_perfil_proveedor_fecha_modificacion" ON "proveedores"."perfil_proveedor" ("fecha_modificacion");

CREATE INDEX "idx_perfil_proveedor_autor_modificacion_id" ON "proveedores"."perfil_proveedor" ("autor_modificacion_id");

CREATE INDEX "idx_titulo_proveedor_natural_nucleo_basico_conocimiento_snies_id" ON "proveedores"."titulo_proveedor_natural" ("nucleo_basico_conocimiento_snies_id");

CREATE INDEX "idx_titulo_proveedor_natural_perfil_id" ON "proveedores"."titulo_proveedor_natural" ("perfil_id");

CREATE INDEX "idx_titulo_proveedor_natural_activo" ON "proveedores"."titulo_proveedor_natural" ("activo");

CREATE INDEX "idx_titulo_proveedor_natural_fecha_creacion" ON "proveedores"."titulo_proveedor_natural" ("fecha_creacion");

CREATE INDEX "idx_titulo_proveedor_natural_fecha_modificacion" ON "proveedores"."titulo_proveedor_natural" ("fecha_modificacion");

CREATE INDEX "idx_titulo_proveedor_natural_autor_modificacion_id" ON "proveedores"."titulo_proveedor_natural" ("autor_modificacion_id");

CREATE INDEX "idx_proveedor_juridico_procedencia_empresa_id" ON "proveedores"."proveedor_juridico" ("procedencia_empresa_id");

CREATE INDEX "idx_proveedor_juridico_ciudad_origen_id" ON "proveedores"."proveedor_juridico" ("ciudad_origen_id");

CREATE INDEX "idx_proveedor_juridico_pais_origen_id" ON "proveedores"."proveedor_juridico" ("pais_origen_id");

CREATE INDEX "idx_proveedor_juridico_tipo_conformacion_id" ON "proveedores"."proveedor_juridico" ("tipo_conformacion_id");

CREATE INDEX "idx_proveedor_juridico_regimen_contributivo" ON "proveedores"."proveedor_juridico" ("regimen_contributivo");

CREATE INDEX "idx_proveedor_juridico_activo" ON "proveedores"."proveedor_juridico" ("activo");

CREATE INDEX "idx_proveedor_juridico_fecha_creacion" ON "proveedores"."proveedor_juridico" ("fecha_creacion");

CREATE INDEX "idx_proveedor_juridico_fecha_modificacion" ON "proveedores"."proveedor_juridico" ("fecha_modificacion");

CREATE INDEX "idx_proveedor_juridico_autor_modificacion_id" ON "proveedores"."proveedor_juridico" ("autor_modificacion_id");

CREATE INDEX "idx_actividad_economica_proveedor_proveedor_id" ON "proveedores"."actividad_economica_proveedor" ("proveedor_id");

CREATE INDEX "idx_actividad_economica_proveedor_actividad_economica_id" ON "proveedores"."actividad_economica_proveedor" ("actividad_economica_id");

CREATE INDEX "idx_actividad_economica_proveedor_activo" ON "proveedores"."actividad_economica_proveedor" ("activo");

CREATE INDEX "idx_actividad_economica_proveedor_fecha_creacion" ON "proveedores"."actividad_economica_proveedor" ("fecha_creacion");

CREATE INDEX "idx_actividad_economica_proveedor_fecha_modificacion" ON "proveedores"."actividad_economica_proveedor" ("fecha_modificacion");

CREATE INDEX "idx_actividad_economica_proveedor_autor_modificacion_id" ON "proveedores"."actividad_economica_proveedor" ("autor_modificacion_id");

CREATE INDEX "idx_representacion_proveedor_id" ON "proveedores"."representacion" ("proveedor_id");

CREATE INDEX "idx_representacion_representante_id" ON "proveedores"."representacion" ("representante_id");

CREATE INDEX "idx_representacion_activo" ON "proveedores"."representacion" ("activo");

CREATE INDEX "idx_representacion_fecha_creacion" ON "proveedores"."representacion" ("fecha_creacion");

CREATE INDEX "idx_representacion_fecha_modificacion" ON "proveedores"."representacion" ("fecha_modificacion");

CREATE INDEX "idx_representacion_autor_modificacion_id" ON "proveedores"."representacion" ("autor_modificacion_id");

CREATE INDEX "idx_persona_juridica_tipo_entidad_tipo_entidad_id" ON "proveedores"."persona_juridica_tipo_entidad" ("tipo_entidad_id");

CREATE INDEX "idx_persona_juridica_tipo_entidad_persona_juridica_id" ON "proveedores"."persona_juridica_tipo_entidad" ("persona_juridica_id");

CREATE INDEX "idx_persona_juridica_tipo_entidad_activo" ON "proveedores"."persona_juridica_tipo_entidad" ("activo");

CREATE INDEX "idx_persona_juridica_tipo_entidad_fecha_creacion" ON "proveedores"."persona_juridica_tipo_entidad" ("fecha_creacion");

CREATE INDEX "idx_persona_juridica_tipo_entidad_fecha_modificacion" ON "proveedores"."persona_juridica_tipo_entidad" ("fecha_modificacion");

CREATE INDEX "idx_persona_juridica_tipo_entidad_autor_modificacion_id" ON "proveedores"."persona_juridica_tipo_entidad" ("autor_modificacion_id");

CREATE INDEX "idx_tipo_conformacion_activo" ON "proveedores"."tipo_conformacion" ("activo");

COMMENT ON COLUMN "proveedores"."proveedor"."autor_modificacion_id" IS 'Posible eferencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."contacto"."autor_modificacion_id" IS 'Posible eferencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."cuenta_bancaria"."autor_modificacion_id" IS 'Posible eferencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."documento"."autor_modificacion_id" IS 'Posible eferencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."proveedor_natural"."autor_modificacion_id" IS 'Posible eferencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."perfil_proveedor"."autor_modificacion_id" IS 'Posible eferencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."titulo_proveedor_natural"."autor_modificacion_id" IS 'Posible eferencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."proveedor_juridico"."autor_modificacion_id" IS 'Posible eferencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."actividad_economica_proveedor"."autor_modificacion_id" IS 'Posible eferencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."representacion"."autor_modificacion_id" IS 'Posible eferencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."persona_juridica_tipo_entidad"."autor_modificacion_id" IS 'Posible eferencia externa a Terceros CRUD';

COMMENT ON TABLE "proveedores"."tipo_conformacion" IS 'Posible procedencia de Parametros CRUD.';

ALTER TABLE "proveedores"."proveedor_juridico" ADD CONSTRAINT "fk_proveedor_juridico_tipo_conformacion" FOREIGN KEY ("tipo_conformacion_id") REFERENCES "proveedores"."tipo_conformacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."actividad_economica_proveedor" ADD CONSTRAINT "fk_actividad_economica_proveedor_proveedor" FOREIGN KEY ("proveedor_id") REFERENCES "proveedores"."proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."representacion" ADD CONSTRAINT "fk_representacion_proveedor" FOREIGN KEY ("proveedor_id") REFERENCES "proveedores"."proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."proveedor_natural" ADD CONSTRAINT "fk_proveedor_natural_proveedor" FOREIGN KEY ("id") REFERENCES "proveedores"."proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."proveedor_juridico" ADD CONSTRAINT "fk_proveedor_juridico_proveedor" FOREIGN KEY ("id") REFERENCES "proveedores"."proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."contacto" ADD CONSTRAINT "fk_contacto_proveedor" FOREIGN KEY ("proveedor_id") REFERENCES "proveedores"."proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."cuenta_bancaria" ADD CONSTRAINT "fk_cuenta_bancaria_proveedor" FOREIGN KEY ("proveedor_id") REFERENCES "proveedores"."proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."documento" ADD CONSTRAINT "fk_documento_proveedor" FOREIGN KEY ("proveedor_id") REFERENCES "proveedores"."proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;
