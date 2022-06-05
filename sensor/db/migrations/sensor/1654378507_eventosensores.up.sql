BEGIN;

	CREATE TABLE eventos_sensores(
		id					UUID	NOT NULL,
		time_stamp			NUMERIC	NOT NULL,
		id_sensor			UUID	NOT NULL,
		id_localidade	    UUID	NOT NULL,
		valor				CITEXT			,
		CONSTRAINT		 eventos_sensores_pk
		PRIMARY KEY		 (id),
		CONSTRAINT		 eventos_sensores_sensor_fk1
		FOREIGN KEY		 (id_sensor)
		REFERENCES	     sensores,
		CONSTRAINT		 eventos_localidade_sensor_fk1
		FOREIGN KEY		 (id_localidade)
		REFERENCES	     localidades
		
	);


COMMIT;