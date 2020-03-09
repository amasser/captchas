CREATE TABLE captchas (
    id VARCHAR(64) NOT NULL PRIMARY KEY,
    answer VARCHAR(32) NOT NULL,
    created_at BIGINT NOT NULL,
    expires_in BIGINT NOT NULL,
    INDEX(expires_in)
);