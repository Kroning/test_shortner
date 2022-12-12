create database shortner ENCODING 'UTF8' LC_COLLATE = 'ru_RU.UTF-8' LC_CTYPE = 'ru_RU.UTF-8' TEMPLATE=template0;
CREATE USER shortner WITH PASSWORD '';
GRANT ALL PRIVILEGES ON DATABASE shortner TO shortner;
GRANT SELECT ON ALL TABLES IN SCHEMA INFORMATION_SCHEMA TO shortner;
\c shortner postgres
GRANT ALL ON SCHEMA public TO shortner;

CREATE TABLE links (
  id              SERIAL PRIMARY KEY,
  alias           VARCHAR(10),
  url             text,
  created_at      timestamp,
  deleted_at      timestamp
);
CREATE INDEX links_alias ON links(alias);
