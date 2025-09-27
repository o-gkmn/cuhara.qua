-- +migrate Up
-- Change all primary and foreign keys from INTEGER to BIGINT

-- First, drop all foreign key constraints
ALTER TABLE users DROP CONSTRAINT IF EXISTS users_role_id_fkey;
ALTER TABLE users DROP CONSTRAINT IF EXISTS users_tenant_id_fkey;
ALTER TABLE roles DROP CONSTRAINT IF EXISTS roles_tenant_id_fkey;
ALTER TABLE posts DROP CONSTRAINT IF EXISTS posts_creator_id_fkey;
ALTER TABLE posts DROP CONSTRAINT IF EXISTS posts_subtopic_id_fkey;
ALTER TABLE posts DROP CONSTRAINT IF EXISTS posts_tenant_id_fkey;
ALTER TABLE comments DROP CONSTRAINT IF EXISTS comments_sender_id_fkey;
ALTER TABLE comments DROP CONSTRAINT IF EXISTS comments_answer_id_fkey;
ALTER TABLE comments DROP CONSTRAINT IF EXISTS comments_tenant_id_fkey;
ALTER TABLE answers DROP CONSTRAINT IF EXISTS answers_creator_id_fkey;
ALTER TABLE answers DROP CONSTRAINT IF EXISTS answers_post_id_fkey;
ALTER TABLE answers DROP CONSTRAINT IF EXISTS answers_tenant_id_fkey;
ALTER TABLE claims DROP CONSTRAINT IF EXISTS claims_tenant_id_fkey;
ALTER TABLE topics DROP CONSTRAINT IF EXISTS topics_tenant_id_fkey;
ALTER TABLE sub_topics DROP CONSTRAINT IF EXISTS sub_topics_topic_id_fkey;
ALTER TABLE sub_topics DROP CONSTRAINT IF EXISTS sub_topics_tenant_id_fkey;
ALTER TABLE tags DROP CONSTRAINT IF EXISTS tags_tenant_id_fkey;
ALTER TABLE votes DROP CONSTRAINT IF EXISTS votes_voter_id_fkey;
ALTER TABLE votes DROP CONSTRAINT IF EXISTS votes_answer_id_fkey;
ALTER TABLE votes DROP CONSTRAINT IF EXISTS votes_tenant_id_fkey;

-- Change primary keys to BIGINT
ALTER TABLE tenants ALTER COLUMN id TYPE BIGINT;
ALTER TABLE roles ALTER COLUMN id TYPE BIGINT;
ALTER TABLE users ALTER COLUMN id TYPE BIGINT;
ALTER TABLE topics ALTER COLUMN id TYPE BIGINT;
ALTER TABLE sub_topics ALTER COLUMN id TYPE BIGINT;
ALTER TABLE posts ALTER COLUMN id TYPE BIGINT;
ALTER TABLE answers ALTER COLUMN id TYPE BIGINT;
ALTER TABLE comments ALTER COLUMN id TYPE BIGINT;
ALTER TABLE tags ALTER COLUMN id TYPE BIGINT;
ALTER TABLE votes ALTER COLUMN id TYPE BIGINT;
ALTER TABLE claims ALTER COLUMN id TYPE BIGINT;

-- Change foreign keys to BIGINT
ALTER TABLE users ALTER COLUMN role_id TYPE BIGINT;
ALTER TABLE users ALTER COLUMN tenant_id TYPE BIGINT;
ALTER TABLE roles ALTER COLUMN tenant_id TYPE BIGINT;
ALTER TABLE topics ALTER COLUMN tenant_id TYPE BIGINT;
ALTER TABLE sub_topics ALTER COLUMN topic_id TYPE BIGINT;
ALTER TABLE sub_topics ALTER COLUMN tenant_id TYPE BIGINT;
ALTER TABLE posts ALTER COLUMN creator_id TYPE BIGINT;
ALTER TABLE posts ALTER COLUMN subtopic_id TYPE BIGINT;
ALTER TABLE posts ALTER COLUMN tenant_id TYPE BIGINT;
ALTER TABLE answers ALTER COLUMN creator_id TYPE BIGINT;
ALTER TABLE answers ALTER COLUMN post_id TYPE BIGINT;
ALTER TABLE answers ALTER COLUMN tenant_id TYPE BIGINT;
ALTER TABLE comments ALTER COLUMN sender_id TYPE BIGINT;
ALTER TABLE comments ALTER COLUMN answer_id TYPE BIGINT;
ALTER TABLE comments ALTER COLUMN tenant_id TYPE BIGINT;
ALTER TABLE tags ALTER COLUMN tenant_id TYPE BIGINT;
ALTER TABLE votes ALTER COLUMN voter_id TYPE BIGINT;
ALTER TABLE votes ALTER COLUMN answer_id TYPE BIGINT;
ALTER TABLE votes ALTER COLUMN tenant_id TYPE BIGINT;
ALTER TABLE claims ALTER COLUMN tenant_id TYPE BIGINT;

-- Recreate foreign key constraints
ALTER TABLE users ADD CONSTRAINT users_role_id_fkey FOREIGN KEY (role_id) REFERENCES roles(id);
ALTER TABLE users ADD CONSTRAINT users_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES tenants(id);
ALTER TABLE roles ADD CONSTRAINT roles_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES tenants(id);
ALTER TABLE posts ADD CONSTRAINT posts_creator_id_fkey FOREIGN KEY (creator_id) REFERENCES users(id);
ALTER TABLE posts ADD CONSTRAINT posts_subtopic_id_fkey FOREIGN KEY (subtopic_id) REFERENCES sub_topics(id);
ALTER TABLE posts ADD CONSTRAINT posts_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES tenants(id);
ALTER TABLE comments ADD CONSTRAINT comments_sender_id_fkey FOREIGN KEY (sender_id) REFERENCES users(id);
ALTER TABLE comments ADD CONSTRAINT comments_answer_id_fkey FOREIGN KEY (answer_id) REFERENCES answers(id);
ALTER TABLE comments ADD CONSTRAINT comments_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES tenants(id);
ALTER TABLE answers ADD CONSTRAINT answers_creator_id_fkey FOREIGN KEY (creator_id) REFERENCES users(id);
ALTER TABLE answers ADD CONSTRAINT answers_post_id_fkey FOREIGN KEY (post_id) REFERENCES posts(id);
ALTER TABLE answers ADD CONSTRAINT answers_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES tenants(id);
ALTER TABLE claims ADD CONSTRAINT claims_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES tenants(id);
ALTER TABLE topics ADD CONSTRAINT topics_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES tenants(id);
ALTER TABLE sub_topics ADD CONSTRAINT sub_topics_topic_id_fkey FOREIGN KEY (topic_id) REFERENCES topics(id);
ALTER TABLE sub_topics ADD CONSTRAINT sub_topics_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES tenants(id);
ALTER TABLE tags ADD CONSTRAINT tags_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES tenants(id);
ALTER TABLE votes ADD CONSTRAINT votes_voter_id_fkey FOREIGN KEY (voter_id) REFERENCES users(id);
ALTER TABLE votes ADD CONSTRAINT votes_answer_id_fkey FOREIGN KEY (answer_id) REFERENCES answers(id);
ALTER TABLE votes ADD CONSTRAINT votes_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES tenants(id);