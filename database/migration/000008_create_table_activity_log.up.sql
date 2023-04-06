-- Author: Risky Feryansyah Pribadi
CREATE TABLE IF NOT EXISTS activity_log (
    id varchar(36) PRIMARY KEY NOT NULL,
    action varchar(255),
    fields varchar(255),
    workspace_id varchar(36),
    project_id varchar(36),
    environment_id varchar(36),
    config_id varchar(36)
) engine=InnoDB default charset=utf8mb4 collate=utf8mb4_0900_ai_ci;