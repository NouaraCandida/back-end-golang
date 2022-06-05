
BEGIN;

	CREATE TABLE localidades(
		id					UUID	NOT NULL,
		id_pais				UUID	NOT NULL,
		id_regiao			UUID	NOT NULL,
		CONSTRAINT		 localidades_pk
		PRIMARY KEY		 (id),
		CONSTRAINT		 localidades_pais_regioes_ak1
		UNIQUE			 (id_pais, id_regiao),
		CONSTRAINT       pais_fk1
		FOREIGN KEY		 (id_pais)
		REFERENCES		 paises,
		CONSTRAINT       regioes_fk1
		FOREIGN KEY		 (id_regiao)
		REFERENCES		 regioes 
	);

	insert into localidades(id,id_pais, id_regiao) values('4ac94fd2-0517-4dc7-99c9-644f719e955b','988fff3e-9458-4880-80fa-a3104c43dcbe','988fff3e-9458-4880-80fa-a3104c43dcba');
	insert into localidades(id,id_pais, id_regiao) values('964ee9b0-6c30-494a-83cb-dc61c2e950bb','988fff3e-9458-4880-80fa-a3104c43dcbe','0cc22945-6805-44cc-9c38-50675245c1aa');
	insert into localidades(id,id_pais, id_regiao) values('e9ef6beb-c341-4a29-8e7d-935a76b95cf7','988fff3e-9458-4880-80fa-a3104c43dcbe', '509e0816-6c29-4075-abbc-a034649b8e5d');
	insert into localidades(id,id_pais, id_regiao) values('d7fb4466-ff99-4d1c-91bd-64c70c433c32','988fff3e-9458-4880-80fa-a3104c43dcbe', '26523661-2e74-4275-9d27-8063b52495fd');
	insert into localidades(id,id_pais, id_regiao) values('610b5271-bdd2-4838-9679-1cdfeca51f75','988fff3e-9458-4880-80fa-a3104c43dcbe', '0771ad41-f6fb-46f6-a977-7753279361f0');

COMMIT;