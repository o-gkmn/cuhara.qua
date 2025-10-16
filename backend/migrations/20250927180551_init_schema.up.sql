
-- +migrate Up

CREATE TABLE tenants (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

-- Role table
CREATE TABLE roles (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255) UNIQUE NOT NULL,
    tenant_id BIGINT NOT NULL REFERENCES tenants(id),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

-- Claim table
CREATE TABLE claims (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    tenant_id BIGINT NOT NULL REFERENCES tenants(id),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

-- User table
CREATE TABLE users (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    vsc_account VARCHAR(255),
    role_id BIGINT NOT NULL REFERENCES roles(id),
    tenant_id BIGINT NOT NULL REFERENCES tenants(id),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

-- Topic table
CREATE TABLE topics (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255) NOT NULL,
    tenant_id BIGINT NOT NULL REFERENCES tenants(id),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

-- SubTopic table
CREATE TABLE sub_topics (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255) NOT NULL,
    topic_id BIGINT NOT NULL REFERENCES topics(id),
    tenant_id BIGINT NOT NULL REFERENCES tenants(id),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

-- Post table
CREATE TABLE posts (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    creator_id BIGINT NOT NULL REFERENCES users(id),
    subtopic_id BIGINT NOT NULL REFERENCES sub_topics(id),
    tenant_id BIGINT NOT NULL REFERENCES tenants(id),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

-- Answer table
CREATE TABLE answers (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    body TEXT NOT NULL,
    is_accepted BOOLEAN DEFAULT false,
    is_first_reply BOOLEAN DEFAULT false,
    creator_id BIGINT NOT NULL REFERENCES users(id),
    post_id BIGINT NOT NULL REFERENCES posts(id),
    tenant_id BIGINT NOT NULL REFERENCES tenants(id),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

-- Vote table
CREATE TABLE votes (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    voter_id BIGINT NOT NULL REFERENCES users(id),
    answer_id BIGINT NOT NULL REFERENCES answers(id),
    is_owner_vote BOOLEAN DEFAULT false,
    tenant_id BIGINT NOT NULL REFERENCES tenants(id),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

-- Comment table
CREATE TABLE comments (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    body TEXT NOT NULL,
    sender_id BIGINT NOT NULL REFERENCES users(id),
    answer_id BIGINT NOT NULL REFERENCES answers(id),
    tenant_id BIGINT NOT NULL REFERENCES tenants(id),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

-- Tag table
CREATE TABLE tags (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255) UNIQUE NOT NULL,
    tenant_id BIGINT NOT NULL REFERENCES tenants(id),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

-- PostTags join table
CREATE TABLE post_tags (
    post_id BIGINT NOT NULL REFERENCES posts(id),
    tag_id BIGINT NOT NULL REFERENCES tags(id),
    PRIMARY KEY(post_id, tag_id)
);

-- RoleClaims join table
CREATE TABLE role_claims (
    role_id BIGINT NOT NULL REFERENCES roles(id),
    claim_id BIGINT NOT NULL REFERENCES claims(id),
    PRIMARY KEY(role_id, claim_id)
);

-- UserClaims join table
CREATE TABLE user_claims (
    user_id BIGINT NOT NULL REFERENCES users(id),
    claim_id BIGINT NOT NULL REFERENCES claims(id),
    PRIMARY KEY(user_id, claim_id)
);