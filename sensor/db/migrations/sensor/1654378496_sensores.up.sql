BEGIN;

	CREATE TABLE sensores(
		id				UUID	NOT NULL,
		nome			CITEXT	NOT NULL,
		nome_regiao   	CITEXT    NOT NULL,
		nome_pais 		CITEXT NOT NULL
		created_at 		TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		updated_at 		TIMESTAMPTZ,

		CONSTRAINT		sensores_pk
		PRIMARY KEY		(id),
		CONSTRAINT		sensores_nome_ak1
		UNIQUE			(nome, id_localidade),
		CONSTRAINT		sensor_localidade_fk1
		FOREIGN KEY		(id_localidade)
		REFERENCES	     localidades
	);




COMMIT;