BEGIN;

	CREATE TABLE sensores(
		id				UUID	NOT NULL,
		nome			CITEXT	NOT NULL,
		id_localidade   UUID    NOT NULL,
		created_at 		timestamptz NOT NULL DEFAULT NOW(),
		updated_at 		timestamptz,

		CONSTRAINT		sensores_pk
		PRIMARY KEY		(id),
		CONSTRAINT		sensores_nome_ak1
		UNIQUE			(nome, id_localidade),
		CONSTRAINT		sensor_localidade_fk1
		FOREIGN KEY		(id_localidade)
		REFERENCES	     localidades
	);

	

	insert into sensores(id,nome, id_localidade) values('988fff3e-9458-4880-80fa-a3104c43dcba','sensor001','4ac94fd2-0517-4dc7-99c9-644f719e955b');
	insert into sensores(id,nome, id_localidade)  values('0cc22945-6805-44cc-9c38-50675245c1aa','sensor002','4ac94fd2-0517-4dc7-99c9-644f719e955b');
	insert into sensores(id,nome, id_localidade)  values('509e0816-6c29-4075-abbc-a034649b8e5d','sensor003', '4ac94fd2-0517-4dc7-99c9-644f719e955b');
	insert into sensores(id,nome, id_localidade)  values('26523661-2e74-4275-9d27-8063b52495fd','sensor004', '4ac94fd2-0517-4dc7-99c9-644f719e955b');
	insert into sensores(id,nome, id_localidade)  values('0771ad41-f6fb-46f6-a977-7753279361f0','sensor005', '4ac94fd2-0517-4dc7-99c9-644f719e955b');

	insert into sensores(id,nome, id_localidade) values('06cc1332-fc36-493b-97f8-d7c1496f34d7','sensor001','964ee9b0-6c30-494a-83cb-dc61c2e950bb');
	insert into sensores(id,nome, id_localidade)  values('18e239ca-3fad-4826-ba26-6299df47f936','sensor002','964ee9b0-6c30-494a-83cb-dc61c2e950bb');
	insert into sensores(id,nome, id_localidade)  values('ce8ad725-95ed-4dff-ab4e-1d60b0051170','sensor003', '964ee9b0-6c30-494a-83cb-dc61c2e950bb');
	insert into sensores(id,nome, id_localidade)  values('705bdd11-9733-4a9d-b192-1724befcccd9','sensor004', '964ee9b0-6c30-494a-83cb-dc61c2e950bb');
	insert into sensores(id,nome, id_localidade)  values('476582e0-f923-44a7-9b61-6e489a91e073','sensor005', '964ee9b0-6c30-494a-83cb-dc61c2e950bb');

    insert into sensores(id,nome, id_localidade) values('becf1cf9-9bfb-42b9-b3a2-29156714d84a','sensor001','e9ef6beb-c341-4a29-8e7d-935a76b95cf7');
	insert into sensores(id,nome, id_localidade)  values('0ea3340c-5a3d-4142-87dd-6e37ea90627b','sensor002','e9ef6beb-c341-4a29-8e7d-935a76b95cf7');
	insert into sensores(id,nome, id_localidade)  values('ee81e5fa-20a9-4a9d-b803-96d80f3f6f56','sensor003', 'e9ef6beb-c341-4a29-8e7d-935a76b95cf7');
	insert into sensores(id,nome, id_localidade)  values('7c86fb6e-a365-4b84-996f-741818f30309','sensor004', 'e9ef6beb-c341-4a29-8e7d-935a76b95cf7');
	insert into sensores(id,nome, id_localidade)  values('2f67ecd1-ea32-4363-b545-da4c23018d36','sensor005', 'e9ef6beb-c341-4a29-8e7d-935a76b95cf7');

    insert into sensores(id,nome, id_localidade) values('962a2bd7-1863-426b-a8ad-df074959380e','sensor001','d7fb4466-ff99-4d1c-91bd-64c70c433c32');
	insert into sensores(id,nome, id_localidade)  values('b4586927-7147-4c62-a548-792fb5fce5d5','sensor002','d7fb4466-ff99-4d1c-91bd-64c70c433c32');
	insert into sensores(id,nome, id_localidade)  values('15607a6b-9340-4e73-bbca-317f1b41f0e3','sensor003', 'd7fb4466-ff99-4d1c-91bd-64c70c433c32');
	insert into sensores(id,nome, id_localidade)  values('7a1a7d6d-82c3-45fd-8117-64b9184e6826','sensor004', 'd7fb4466-ff99-4d1c-91bd-64c70c433c32');
	insert into sensores(id,nome, id_localidade)  values('de9d8be5-125d-4eca-815a-116680f2bfa7','sensor005', 'd7fb4466-ff99-4d1c-91bd-64c70c433c32');

	insert into sensores(id,nome, id_localidade) values('06e3abff-bdf8-4085-99fc-50833ba1b6f6','sensor001','610b5271-bdd2-4838-9679-1cdfeca51f75');
	insert into sensores(id,nome, id_localidade)  values('2621492f-489e-408d-a87d-d250363e3620','sensor002','610b5271-bdd2-4838-9679-1cdfeca51f75');
	insert into sensores(id,nome, id_localidade)  values('f5428658-fae9-4908-ac16-8a1a38e6f562','sensor003', '610b5271-bdd2-4838-9679-1cdfeca51f75');
	insert into sensores(id,nome, id_localidade)  values('6bce1b2f-c335-49bb-9222-8adc8293f948','sensor004', '610b5271-bdd2-4838-9679-1cdfeca51f75');
	insert into sensores(id,nome, id_localidade)  values('671447e4-f098-48a3-a5ea-b4c4a382a825','sensor005', '610b5271-bdd2-4838-9679-1cdfeca51f75');

COMMIT;