--
-- Seed script for the abstracted_self database.
--
-- Populates every table (technology_types, technologies, images, blogs,
-- blog_images, projects, project_images, experiences, daily_visits) with
-- demo/portfolio data consistent with the profile in
-- backend/internal/db/profile_data.json.
--
-- The whole script runs inside a single transaction: if any statement
-- fails, the entire seed is rolled back. Safe to re-run -- it first
-- truncates all rows (restarting sequences).
--
-- Run manually:
--   docker compose exec -T postgres psql -U postgres -d abstracted_self -f - < postgres-init/seed.sql
-- Or automatically on a fresh container (docker-entrypoint-initdb.d loads
-- .sql files in alphabetical order: init.sql then seed.sql).
--

BEGIN;

TRUNCATE TABLE blog_images, blogs, daily_visits, experiences, images,
             project_images, projects, technologies, technology_types
RESTART IDENTITY CASCADE;

-- ---------------------------------------------------------------------------
-- technology_types & technologies
-- ---------------------------------------------------------------------------

INSERT INTO technology_types (id, type_name) VALUES
    (1, 'Programming Language'),
    (2, 'Web Framework'),
    (3, 'Frontend Framework'),
    (4, 'Database'),
    (5, 'Tools & DevOps');

INSERT INTO technologies (id, type_id, name) VALUES
    (1,  1, 'Go'),
    (2,  1, 'JavaScript'),
    (3,  1, 'Python'),
    (4,  1, 'PHP'),
    (5,  1, 'TypeScript'),
    (6,  2, 'Gin'),
    (7,  2, 'Flask'),
    (8,  2, 'Laravel'),
    (9,  2, 'NestJS'),
    (10, 3, 'Vue'),
    (11, 3, 'React'),
    (12, 4, 'PostgreSQL'),
    (13, 4, 'MySQL'),
    (14, 5, 'Docker'),
    (15, 5, 'Git'),
    (16, 5, 'Redis');

-- ---------------------------------------------------------------------------
-- images
-- Takes precedence over blogs/projects so FKs to images resolve.
-- ---------------------------------------------------------------------------

INSERT INTO images (id, file_name, file_size, mime_type) VALUES
    (1, 'blog-hello-world-cover.jpg', 245008, 'image/jpeg'),
    (2, 'blog-go-cover.png',          182344, 'image/png'),
    (3, 'blog-gin-cover.png',         210112, 'image/png'),
    (4, 'blog-hello-world-content.jpg', 98211, 'image/jpeg'),
    (5, 'blog-go-content.png',       156730, 'image/png'),
    (6, 'blog-gin-content.png',      175220, 'image/png'),
    (7, 'project-abstracted-self.png',   128400, 'image/png'),
    (8, 'project-inventory.png',        112900, 'image/png'),
    (9, 'project-recipe-finder.png',     98700, 'image/png');

SELECT setval('images_id_seq', (SELECT MAX(id) FROM images));

-- ---------------------------------------------------------------------------
-- blogs & blog_images
-- ---------------------------------------------------------------------------

INSERT INTO blogs (id, title, slug, content, blog_snippet, published, cover_image_id, created_at, updated_at) VALUES
    (1, 'Hello World: Beginning My Software Engineering Journey',
     'hello-world-beginning-my-software-engineering-journey',
     $blog1$<p>This is my first post! I'm Bani, a fresh graduate in informatics who is passionate about turning ideas into impactful software solutions.</p><p>During my studies and internship I got hands-on experience building web applications with Python, JavaScript, PHP, Go and Vue. Lately I have been focusing on backend development, learning how to build robust and reliable APIs.</p><p>I'm excited to document my learnings here, from database design to deployment.</p>$blog1$,
     'My first post about why I started this blog and where my journey is heading.',
     true, 1, '2025-01-15 09:00:00', '2025-01-15 09:00:00'),
    (2, 'Why I Fell in Love with Go',
     'why-i-fell-in-love-with-go',
     $blog2$<p>Go is a language I was skeptical about at first, but after writing a couple of production-grade REST APIs with Gin, I never looked back.</p><p>Concurrency with goroutines felt approachable, the standard library is a joy to read, and compiling to a single static binary makes deployment with Docker so simple.</p><p>If you are coming from dynamic languages, Go's type system may feel verbose at first, but it pays off with compile-time safety.</p>$blog2$,
     'Notes on what made me pick Go for backend work and why I keep recommending it.',
     true, 2, '2025-02-20 10:30:00', '2025-03-01 08:15:00'),
    (3, 'From Laravel to Gin: Migrating My Web Stack',
     'from-laravel-to-gin-migrating-my-web-stack',
     $blog3$<p>I spent most of my internship writing Laravel and Vue, so learning Gin felt like a big change. This post collects the notes I wished I had found earlier.</p><p>The biggest mental shift was moving from an expressive ORM to writing raw SQL. Writing queries by hand taught me more about PostgreSQL in a month than a year of Eloquent.</p><p>I cover middleware, JWT auth, graceful shutdown, and how to structure a small Go project so it stays maintainable as it grows.</p>$blog3$,
     'A practical write-up about moving from a PHP framework to the Go/Gin ecosystem.',
     true, 3, '2025-04-05 14:00:00', '2025-04-10 18:45:00');

INSERT INTO blog_images (blog_id, image_id) VALUES
    (1, 4),
    (2, 5),
    (3, 6);

SELECT setval('blogs_id_seq', (SELECT MAX(id) FROM blogs));

-- ---------------------------------------------------------------------------
-- projects & project_images
-- ---------------------------------------------------------------------------

INSERT INTO projects (id, name, description, tech_stack, source_url, project_url, start_date, end_date) VALUES
    (1, 'abstracted_self',
     $p1$My personal portfolio and blog platform. A Vue 3 + Tailwind frontend served behind nginx, a Go/Gin REST API, and a PostgreSQL database all wired together with Docker Compose.$p1$,
     ARRAY['Go', 'Gin', 'PostgreSQL', 'Vue 3', 'Tailwind CSS', 'Docker'],
     ARRAY['https://github.com/husni-robani/abstracted_self'],
     NULL,
     '2025-01-10 00:00:00', '2025-04-30 00:00:00'),
    (2, 'Warehouse Inventory System',
     $p2$An internal inventory management dashboard built with Laravel and Vue. Features role-based access control, stock tracking, barcode scanning, and report generation.$p2$,
     ARRAY['Laravel', 'Vue', 'MySQL', 'Redis', 'Docker'],
     ARRAY['https://github.com/husni-robani/warehouse-inventory'],
     NULL,
     '2024-08-01 00:00:00', '2024-12-15 00:00:00'),
    (3, 'Recipe Finder',
     $p3$A recipe search web app with Flask on the backend and React on the frontend. Scrapes ingredients, maps them to recipes, and serves them through a caching layer.$p3$,
     ARRAY['Python', 'Flask', 'React', 'PostgreSQL', 'Redis'],
     ARRAY['https://github.com/husni-robani/recipe-finder'],
     'https://recipe-finder.example.com',
     '2023-04-01 00:00:00', '2023-09-30 00:00:00');

INSERT INTO project_images (project_id, image_id) VALUES
    (1, 7),
    (2, 8),
    (3, 9);

SELECT setval('projects_id_seq', (SELECT MAX(id) FROM projects));

-- ---------------------------------------------------------------------------
-- experiences
-- ---------------------------------------------------------------------------

INSERT INTO experiences (id, job_title, company_name, work_place, start_date, end_date, description, accomplishments, tech_stack) VALUES
    (1, 'Backend Engineer', 'PT Digital Nusantara', 'Jakarta, Indonesia',
     '2025-02-01', NULL,
     $e1$Build and maintain REST APIs for customer-facing services, focusing on performance and code quality.$e1$,
     ARRAY[$a1$Designed the API and database schema powering 3 microservices$a1$, $a2$Cut paginated query latency by 40% with proper indexing and caching$a2$],
     ARRAY['Go', 'Gin', 'PostgreSQL', 'Redis', 'Docker']),
    (2, 'Backend Engineer Intern', 'PT Maju Bersama', 'Bandung, Indonesia',
     '2024-06-01', '2024-12-31',
     $e2$Interned on an internal LMS platform, shipped features end-to-end and learned production debugging workflows.$e2$,
     ARRAY[$a3$Shipped the module progress tracking feature used by 500+ users$a3$, $a4$Wrote integration tests that became the team's baseline for new endpoints$a4$],
     ARRAY['Laravel', 'Vue', 'MySQL', 'Git']),
    (3, 'Freelance Web Developer', 'Self-employed', 'Remote',
     '2023-03-01', '2024-05-31',
     $e3$Built websites and small applications for local clients end-to-end, from requirement gathering to deployment.$e3$,
     ARRAY[$a5$Delivered 6 client projects on schedule with 95%+ satisfaction feedback$a5$],
     ARRAY['Python', 'Flask', 'JavaScript', 'Vue']);

SELECT setval('experiences_id_seq', (SELECT MAX(id) FROM experiences));

-- ---------------------------------------------------------------------------
-- daily_visits
-- ---------------------------------------------------------------------------

INSERT INTO daily_visits (uuid, visit_date, ip, device) VALUES
    ('6b6f8c90-6a1c-4f3e-9c0d-2e2f3a4b5c6d', '2025-08-01', '103.10.20.30', 'Desktop'),
    ('7c7f9d01-7b2d-4f4e-ad1e-3f3f4b5c6d7e', '2025-08-02', '114.20.30.40', 'Mobile'),
    ('8d8fae12-8c3e-4f5f-be2f-4f4f5c6d7e8f', '2025-08-03', '125.30.40.50', 'Tablet'),
    ('9e9fbf23-9d4f-4f60-cf30-5f5f6d7e8f90', '2025-08-03', '136.40.50.60', 'Desktop'),
    ('afafc034-a5f4-4f71-d041-6a6b7e8f9a0b', '2025-08-04', '147.50.60.70', 'Mobile');

COMMIT;