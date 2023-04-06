-- Author: Risky Feryansyah Pribadi
CREATE TABLE IF NOT EXISTS workspace (
    id varchar(36) PRIMARY KEY NOT NULL,
    name varchar(255) NOT NULL,
    user_id varchar(36) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp,
    updated_by varchar(36)
) engine=InnoDB default charset=utf8mb4 collate=utf8mb4_0900_ai_ci;