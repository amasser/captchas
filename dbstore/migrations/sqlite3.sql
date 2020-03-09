BEGIN;
CREATE TABLE captchas (
    id VARCHAR(64) NOT NULL PRIMARY KEY,
    answer VARCHAR(32) NOT NULL,
    created_at BIGINT NOT NULL,
    expires_in BIGINT NOT NULL
);
CREATE INDEX captchas_expires_in_idx ON captchas (expires_in);
COMMIT;