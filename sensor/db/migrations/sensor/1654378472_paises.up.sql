BEGIN;

	CREATE TABLE paises(
		id				UUID	NOT NULL,
		nome			CITEXT	NOT NULL,
		CONSTRAINT		pais_pk
		PRIMARY KEY		(id),
		CONSTRAINT		pais_nome_ak1
		UNIQUE			(nome)
	);

	insert into paises(id,nome)
	values('988fff3e-9458-4880-80fa-a3104c43dcbe','Brasil');

COMMIT;