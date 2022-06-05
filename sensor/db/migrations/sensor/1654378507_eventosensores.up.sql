BEGIN;

	CREATE TABLE eventos_sensores(
		id					UUID			NOT NULL,
		id_sensor			UUID			NOT NULL,
		valor				CITEXT,
		created_at 							TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		updated_at 							TIMESTAMPTZ,

		CONSTRAINT		 eventos_sensores_pk
		PRIMARY KEY		 (id),
		CONSTRAINT		 eventos_sensores_sensor_fk1
		FOREIGN KEY		 (id_sensor)
		REFERENCES	     sensores,
		CONSTRAINT		 eventos_localidade_sensor_fk1

		
	);


COMMIT;