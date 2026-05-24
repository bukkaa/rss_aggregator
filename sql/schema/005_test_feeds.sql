-- +goose Up

insert into feeds (id, name, url, user_id, created_at, updated_at)
values 
    ('089d4561-b07a-445d-a0e7-a0b138fd8a11', 'Garbage Blog', 'https://garbage/index.xml', '8b6e9052-7c3c-47fc-8356-6c8d3843578c', '2026-05-23 23:26:38.636885', '2026-05-23 23:26:38.636885'),
    ('3ae4e591-6a81-483d-98f2-26efde635616', 'Third blog', 'https://th.ird/index.xml', 'ac10e478-4c2f-49cb-b6c8-25f44a235440', '2026-05-23 23:18:14.264784', '2026-05-23 23:18:14.264784'),
    ('868fa9d9-df79-4ce9-a54e-8ef5f9a81fb7', 'Second blog', 'https://sec.ond/index.xml', 'ac10e478-4c2f-49cb-b6c8-25f44a235440', '2026-05-23 23:18:02.174576', '2026-05-23 23:18:02.174576'),
    ('20859613-ecf0-4ca7-896b-3879ef620944', 'Lane''s blog', 'https://wagslane.dev/index.xml', 'ac10e478-4c2f-49cb-b6c8-25f44a235440', '2026-05-23 23:17:33.823965', '2026-05-23 23:17:33.823965')
;

-- +goose Down
truncate table feeds;