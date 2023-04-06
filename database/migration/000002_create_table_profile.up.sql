-- Author: Risky Feryansyah Pribadi
CREATE TABLE IF NOT EXISTS profile (
    id varchar(36) PRIMARY KEY NOT NULL,
    full_name varchar(255) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp,
    updated_by varchar(36)
) engine=InnoDB default charset=utf8mb4 collate=utf8mb4_0900_ai_ci;