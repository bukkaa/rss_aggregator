-- +goose Up

insert into users (id, created_at, updated_at, name, api_key)
values 
    ('8b6e9052-7c3c-47fc-8356-6c8d3843578c', '2026-05-18 22:56:14.259617', '2026-05-18 22:56:14.259617', 'JSON Momoa', '8ee16f3f47fa000d2a5a49e33082c765404ef08023506a0d56b0c1a9f06bb000'),
    ('ac10e478-4c2f-49cb-b6c8-25f44a235440', '2026-05-18 23:05:06.825038', '2026-05-19 00:38:16.021516', 'Michael Jackson', 'cd4ef161f1a669cb4ef83c1f3bc184a9c8f90596e2db0628d0acad1d3c28a23d')
;

-- +goose Down
truncate table users;