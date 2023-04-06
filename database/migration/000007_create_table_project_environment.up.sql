-- Author: Risky Feryansyah Pribadi
CREATE TABLE IF NOT EXISTS project_environment (
    id varchar(36) PRIMARY KEY NOT NULL,
    workspace_id varchar(36),
    project_id varchar(36),
    environment_id varchar(36),
    config_id varchar(36),
    CONSTRAINT UC_project_environment UNIQUE (workspace_id, project_id, environment_id, config_id)
) engine=InnoDB default charset=utf8mb4 collate=utf8mb4_0900_ai_ci;