--
-- Rollback script for postgres-init/seed.sql.
--
-- Removes ALL seeded data from every table (technology_types, technologies,
-- images, blogs, blog_images, projects, project_images, experiences,
-- daily_visits), leaving the database empty but with the schema intact.
-- Sequences are restarted, so a subsequent seed run starts at id 1 again.
--
-- The whole script runs inside a single transaction.
--
-- This file lives OUTSIDE postgres-init/ on purpose: docker-entrypoint-
-- initdb.d executes every *.sql there on a fresh container, which would
-- wipe the seed immediately after it is written.
--
-- Run manually:
--   docker compose exec -T postgres psql -U postgres -d abstracted_self -f - < seed_rollback.sql
--

BEGIN;

TRUNCATE TABLE blog_images, blogs, daily_visits, experiences, images,
             project_images, projects, technologies, technology_types
RESTART IDENTITY CASCADE;

COMMIT;