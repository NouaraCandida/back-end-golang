BEGIN;

	CREATE TABLE regioes(
		id				UUID	NOT NULL,
		nome			CITEXT	NOT NULL,
		CONSTRAINT		regioes_pk
		PRIMARY KEY		(id),
		CONSTRAINT		regioes_nome_ak1
		UNIQUE			(nome)
	);

	insert into regioes(id,nome) values('988fff3e-9458-4880-80fa-a3104c43dcba','Sul');
	insert into regioes(id,nome) values('0cc22945-6805-44cc-9c38-50675245c1aa','Sudeste');
	insert into regioes(id,nome) values('509e0816-6c29-4075-abbc-a034649b8e5d','Centro Oeste');
	insert into regioes(id,nome) values('26523661-2e74-4275-9d27-8063b52495fd','Nordeste');
	insert into regioes(id,nome) values('0771ad41-f6fb-46f6-a977-7753279361f0','Norte');

COMMIT;