-- Author: Risky Feryansyah Pribadi
CREATE TABLE IF NOT EXISTS user (
    id varchar(36) PRIMARY KEY NOT NULL,
    email varchar(100) NOT NULL,
    password varchar(255) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp,
    updated_by varchar(36),
    UNIQUE (email)
) engine=InnoDB default charset=utf8mb4 collate=utf8mb4_0900_ai_ci;