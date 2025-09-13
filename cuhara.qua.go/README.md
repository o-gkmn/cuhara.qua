.
├── cmd/
│   └── api/
│       └── main.go                  # uygulama giriş noktası (bootstrap)
├── internal/
│   ├── app/
│   │   ├── bootstrap.go             # DI, bus'ların kaydı, router wiring
│   │   └── router.go                # global router (HTTP)
│   ├── common/
│   │   ├── cqrs/
│   │   │   ├── command.go           # Command marker + base
│   │   │   ├── command_bus.go       # Komut otobüsü (sync/async strateji)
│   │   │   ├── command_handler.go   # ICommandHandler<T>
│   │   │   ├── query.go             # Query marker + base
│   │   │   ├── query_bus.go         # Sorgu otobüsü
│   │   │   └── query_handler.go     # IQueryHandler<Q, R>
│   │   ├── dto/                     # Ortak DTO’lar (pagination vb.)
│   │   ├── errors/                  # Domain + app hataları haritalama
│   │   ├── middleware/              # Logger, recovery, auth, CORS
│   │   └── validation/              # request validation yardımcıları
│   ├── infra/
│   │   ├── db/
│   │   │   ├── write_pool.go        # write DB (transactional)
│   │   │   └── read_pool.go         # read DB (replica/readonly)
│   │   ├── server.go                # http.Server ayarları (graceful)
│   │   └── config/
│   │       └── config.go            # env/config yükleme
│   ├── readmodel/                   # Okuma tarafı “projection/read-store” katmanı
│   │   ├── users/
│   │   │   ├── model.go             # Read model (UI’ya uygun)
│   │   │   └── repo_pg.go           # Read repo (sadece SELECT)
│   │   └── tenants/
│   │       ├── model.go
│   │       └── repo_pg.go
│   ├── users/                       # Bounded Context: Users
│   │   ├── domain/
│   │   │   ├── model.go             # Aggregate/Entity (User)
│   │   │   ├── repo.go              # Write-side port (interface)
│   │   │   └── events.go            # (opsiyonel) domain events
│   │   ├── application/
│   │   │   ├── command/
│   │   │   │   ├── create_user/
│   │   │   │   │   ├── command.go   # CreateUserCommand
│   │   │   │   │   └── handler.go   # CreateUserHandler (uses domain repo)
│   │   │   │   └── update_user_email/
│   │   │   │       ├── command.go
│   │   │   │       └── handler.go
│   │   │   └── query/
│   │   │       ├── get_user_by_id/
│   │   │       │   ├── query.go     # GetUserByIdQuery
│   │   │       │   └── handler.go   # handler readmodel üzerinden okur
│   │   │       └── list_users_by_tenant/
│   │   │           ├── query.go
│   │   │           └── handler.go
│   │   ├── infrastructure/
│   │   │   ├── repo_pg.go           # Write-side repo implementasyonu (PG)
│   │   │   └── projector.go         # (opsiyonel) write event → read model güncelleme
│   │   └── interface/
│   │       └── http/
│   │           ├── users_controller.go  # HTTP endpoint → bus.Dispatch
│   │           └── users_routes.go      # /api/v1/users rotaları
│   └── tenants/                     # Bounded Context: Tenants
│       ├── domain/
│       │   ├── model.go
│       │   └── repo.go
│       ├── application/
│       │   ├── command/
│       │   │   └── create_tenant/
│       │   │       ├── command.go
│       │   │       └── handler.go
│       │   └── query/
│       │       └── get_tenant_by_name/
│       │           ├── query.go
│       │           └── handler.go
│       ├── infrastructure/
│       │   └── repo_pg.go
│       └── interface/
│           └── http/
│               ├── tenants_controller.go
│               └── tenants_routes.go
├── migrations/
│   └── 001_initial.sql
└── openapi.yaml                      # (design-first ise)


datasource db {
  provider = "postgresql"
  url      = env("DATABASE_URL")
}

generator client {
  provider = "prisma-client-js"
}

model Tenant {
  id        Int        @id @default(autoincrement()) @map("id")
  name      String     @unique @map("name")
  users     User[]
  topics    Topic[]
  subTopics SubTopic[]
  posts     Post[]
  answers   Answer[]
  votes     Vote[]
  comments  Comment[]
  tags      Tag[]
  roles     Role[]
  claims    Claim[]
  createdAt DateTime   @default(now()) @map("created_at")
  updatedAt DateTime   @updatedAt @map("updated_at")

  @@map("tenants")
}

model User {
  id         Int       @id @default(autoincrement()) @map("id")
  name       String    @map("name")
  email      String    @unique @map("email")
  vscAccount String    @map("vsc_account")
  posts      Post[]
  answers    Answer[]
  votes      Vote[]
  comments   Comment[]
  roleId     Int       @map("role_id")
  role       Role      @relation(fields: [roleId], references: [id])
  claims     Claim[]   @relation("UserClaims")
  tenantId   Int       @map("tenant_id")
  tenant     Tenant    @relation(fields: [tenantId], references: [id])
  createdAt  DateTime  @default(now()) @map("created_at")
  updatedAt  DateTime  @updatedAt @map("updated_at")

  @@map("users")
}

model Role {
  id        Int      @id @default(autoincrement()) @map("id")
  name      String   @unique @map("name")
  users     User[]
  claims    Claim[]  @relation("RoleClaims")
  tenantId  Int      @map("tenant_id")
  tenant    Tenant   @relation(fields: [tenantId], references: [id])
  createdAt DateTime @default(now()) @map("created_at")
  updatedAt DateTime @updatedAt @map("updated_at")

  @@map("roles")
}

model Claim {
  id          Int      @id @default(autoincrement()) @map("id")
  name        String   @map("name")
  description String?  @map("description")
  roles       Role[]   @relation("RoleClaims")
  users       User[]   @relation("UserClaims")
  tenantId    Int      @map("tenant_id")
  tenant      Tenant   @relation(fields: [tenantId], references: [id])
  createdAt   DateTime @default(now()) @map("created_at")
  updatedAt   DateTime @updatedAt @map("updated_at")

  @@map("claims")
}

model Topic {
  id        Int        @id @default(autoincrement()) @map("id")
  name      String     @map("name")
  subTopics SubTopic[]
  tenantId  Int        @map("tenant_id")
  tenant    Tenant     @relation(fields: [tenantId], references: [id])
  createdAt DateTime   @default(now()) @map("created_at")
  updatedAt DateTime   @updatedAt @map("updated_at")

  @@map("topics")
}

model SubTopic {
  id        Int      @id @default(autoincrement()) @map("id")
  name      String   @map("name")
  topicId   Int      @map("topic_id")
  topic     Topic    @relation(fields: [topicId], references: [id])
  posts     Post[]
  tenantId  Int      @map("tenant_id")
  tenant    Tenant   @relation(fields: [tenantId], references: [id])
  createdAt DateTime @default(now()) @map("created_at")
  updatedAt DateTime @updatedAt @map("updated_at")

  @@map("sub_topics")
}

model Post {
  id         Int      @id @default(autoincrement()) @map("id")
  answers    Answer[]
  tags       Tag[]    @relation("PostTags")
  creatorId  Int      @map("creator_id")
  creator    User     @relation(fields: [creatorId], references: [id])
  subtopicId Int      @map("subtopic_id")
  subtopic   SubTopic @relation(fields: [subtopicId], references: [id])
  tenantId   Int      @map("tenant_id")
  tenant     Tenant   @relation(fields: [tenantId], references: [id])
  createdAt  DateTime @default(now()) @map("created_at")
  updatedAt  DateTime @updatedAt @map("updated_at")

  @@map("posts")
}

model Answer {
  id           Int       @id @default(autoincrement()) @map("id")
  body         String    @map("body")
  isAccepted   Boolean   @default(false) @map("is_accepted")
  isFirstReply Boolean   @default(false) @map("is_first_reply")
  creatorId    Int       @map("creator_id")
  creator      User      @relation(fields: [creatorId], references: [id])
  postId       Int       @map("post_id")
  post         Post      @relation(fields: [postId], references: [id])
  votes        Vote[]
  comments     Comment[]
  tenantId     Int       @map("tenant_id")
  tenant       Tenant    @relation(fields: [tenantId], references: [id])
  createdAt    DateTime  @default(now()) @map("created_at")
  updatedAt    DateTime  @updatedAt @map("updated_at")

  @@map("answers")
}

model Vote {
  id          Int      @id @default(autoincrement()) @map("id")
  voterId     Int      @map("voter_id")
  answerId    Int      @map("answer_id")
  isOwnerVote Boolean  @default(false) @map("is_owner_vote")
  voter       User     @relation(fields: [voterId], references: [id])
  answer      Answer   @relation(fields: [answerId], references: [id])
  tenantId    Int      @map("tenant_id")
  tenant      Tenant   @relation(fields: [tenantId], references: [id])
  createdAt   DateTime @default(now()) @map("created_at")
  updatedAt   DateTime @updatedAt @map("updated_at")

  @@map("votes")
}

model Comment {
  id        Int      @id @default(autoincrement()) @map("id")
  body      String   @map("body")
  senderId  Int      @map("sender_id")
  answerId  Int      @map("answer_id")
  sender    User     @relation(fields: [senderId], references: [id])
  answer    Answer   @relation(fields: [answerId], references: [id])
  tenantId  Int      @map("tenant_id")
  tenant    Tenant   @relation(fields: [tenantId], references: [id])
  createdAt DateTime @default(now()) @map("created_at")
  updatedAt DateTime @updatedAt @map("updated_at")

  @@map("comments")
}

model Tag {
  id        Int      @id @default(autoincrement()) @map("id")
  name      String   @unique @map("name")
  posts     Post[]   @relation("PostTags")
  tenantId  Int      @map("tenant_id")
  tenant    Tenant   @relation(fields: [tenantId], references: [id])
  createdAt DateTime @default(now()) @map("created_at")
  updatedAt DateTime @updatedAt @map("updated_at")

  @@map("tags")
}
