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
