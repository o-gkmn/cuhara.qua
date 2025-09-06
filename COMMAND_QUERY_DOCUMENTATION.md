# Command ve Query Oluşturma Rehberi

Bu dokümantasyon, CQRS (Command Query Responsibility Segregation) pattern'ini kullanarak command ve query'lerin nasıl oluşturulduğunu adım adım açıklamaktadır.

## İçindekiler

1. [CQRS Pattern Genel Bakış](#cqrs-pattern-genel-bakış)
2. [Command Oluşturma Adımları](#command-oluşturma-adımları)
3. [Query Oluşturma Adımları](#query-oluşturma-adımları)
4. [Handler Kayıt İşlemi](#handler-kayıt-işlemi)
5. [Controller'da Kullanım](#controllerda-kullanım)
6. [Örnekler](#örnekler)

## CQRS Pattern Genel Bakış

CQRS pattern'i, uygulamada veri okuma (Query) ve veri yazma (Command) işlemlerini ayırarak daha temiz ve ölçeklenebilir bir mimari sağlar.

### Temel Bileşenler:
- **Command**: Veri değiştiren işlemler (Create, Update, Delete)
- **Query**: Veri okuyan işlemler (Get, List, Search)
- **CommandBus**: Command'ları yöneten merkezi sistem
- **QueryBus**: Query'leri yöneten merkezi sistem
- **Handler**: Command/Query'leri işleyen sınıflar

## Command Oluşturma Adımları

### 1. Request Interface Tanımlama

```typescript
// ==================== REQUEST ====================
export interface CreateUserRequest {
  name: string;
  email: string;
  vscAccount: string;
  roleId: number;
  tenantId: number;
}
```

### 2. Response Interface Tanımlama

```typescript
// ==================== RESPONSE ====================
export interface CreateUserResponse {
  id: number;
}
```

### 3. Command Sınıfı Oluşturma

```typescript
// ==================== COMMAND ====================
export class CreateUserCommand extends BaseCommand<CreateUserResponse> {
  readonly type = "CreateUserCommand";

  constructor(public readonly data: CreateUserRequest) {
    super();
  }
}
```

**Önemli Noktalar:**
- `BaseCommand` sınıfından extend edilir
- `type` property'si benzersiz olmalıdır
- Generic type olarak response interface'i belirtilir
- Constructor'da request data'sı alınır

### 4. Command Handler Oluşturma

```typescript
// ==================== HANDLER ====================
export class CreateUserHandler
  implements CommandHandler<CreateUserCommand, CreateUserResponse>
{
  private prisma: PrismaClient;

  constructor() {
    this.prisma = new PrismaClient();
  }

  async handle(command: CreateUserCommand): Promise<CreateUserResponse> {
    try {
      const user = await this.prisma.user.create({
        data: {
          name: command.data.name,
          email: command.data.email,
          vscAccount: command.data.vscAccount,
          roleId: command.data.roleId,
          tenantId: command.data.tenantId,
        },
      });

      return {
        id: user.id,
      };
    } catch (error) {
      throw new Error(`Failed to create user: ${error}`);
    }
  }
}
```

**Önemli Noktalar:**
- `CommandHandler` interface'ini implement eder
- `handle` method'u async olmalıdır
- Prisma client kullanarak veritabanı işlemleri yapılır
- Error handling eklenir
- Response interface'ine uygun data döndürülür

## Query Oluşturma Adımları

### 1. Request Interface Tanımlama

```typescript
// ==================== REQUEST ====================
export interface GetUserRequest {
  userId: number;
}
```

### 2. Response Interface Tanımlama

```typescript
// ==================== RESPONSE ====================
export interface GetUserResponse {
  id: number;
  name: string;
  email: string;
  vscAccount: string;
  roleId: number;
  tenantId: number;
  createdAt: Date;
  updatedAt: Date;
}
```

### 3. Query Sınıfı Oluşturma

```typescript
// ==================== QUERY ====================
export class GetUserQuery extends BaseQuery<GetUserResponse> {
  readonly type = "GetUserQuery";

  constructor(public readonly data: GetUserRequest) {
    super();
  }
}
```

**Önemli Noktalar:**
- `BaseQuery` sınıfından extend edilir
- `type` property'si benzersiz olmalıdır
- Generic type olarak response interface'i belirtilir

### 4. Query Handler Oluşturma

```typescript
// ==================== HANDLER ====================
export class GetUserHandler
  implements QueryHandler<GetUserQuery, GetUserResponse>
{
  private prisma: PrismaClient;

  constructor() {
    this.prisma = new PrismaClient();
  }

  async handle(query: GetUserQuery): Promise<GetUserResponse> {
    try {
      const user = await this.prisma.user.findUnique({
        where: { id: query.data.userId },
      });

      if (!user) {
        throw new Error(`User with id ${query.data.userId} not found`);
      }

      return {
        id: user.id,
        name: user.name,
        email: user.email,
        vscAccount: user.vscAccount,
        roleId: user.roleId,
        tenantId: user.tenantId,
        createdAt: user.createdAt,
        updatedAt: user.updatedAt,
      };
    } catch (error) {
      throw new Error(`Failed to get user: ${error}`);
    }
  }
}
```

## Handler Kayıt İşlemi

Handler'lar `CQRSRegistry` sınıfında kayıt edilmelidir:

```typescript
// src/core/cqrs/registry.ts
private registerHandlers(): void {
  // Register User Command Handlers
  this.commandBus.register("CreateUserCommand", new CreateUserHandler());

  // Register User Query Handlers
  this.queryBus.register("GetUserQuery", new GetUserHandler());
}
```

**Önemli Noktalar:**
- Command type'ı ile handler eşleştirilir
- Query type'ı ile handler eşleştirilir
- Handler instance'ları oluşturulur

## Controller'da Kullanım

### Command Kullanımı

```typescript
async createUser(req: Request, res: Response): Promise<void> {
  try {
    const { name, email, vscAccount, roleId, tenantId } = req.body;

    const command = new CreateUserCommand({
      name,
      email,
      vscAccount,
      roleId,
      tenantId,
    });

    const result = await this.commandBus.execute(command);

    res.status(201).json({
      success: true,
      data: result,
      message: "User created successfully",
    });
  } catch (error) {
    res.status(400).json({
      success: false,
      error: error instanceof Error ? error.message : "Unknown error",
    });
  }
}
```

### Query Kullanımı

```typescript
async getUser(req: Request, res: Response): Promise<void> {
  try {
    const { id } = req.params;
    const userId = parseInt(id);

    if (isNaN(userId)) {
      res.status(400).json({
        success: false,
        error: "Invalid user ID",
      });
      return;
    }

    const query = new GetUserQuery({ userId });
    const result = await this.queryBus.execute(query);

    res.status(200).json({
      success: true,
      data: result,
    });
  } catch (error) {
    res.status(404).json({
      success: false,
      error: error instanceof Error ? error.message : "User not found",
    });
  }
}
```

## Örnekler

### Yeni Command Oluşturma Örneği

Yeni bir `UpdateUserCommand` oluşturmak için:

1. **Request Interface:**
```typescript
export interface UpdateUserRequest {
  id: number;
  name?: string;
  email?: string;
  vscAccount?: string;
  roleId?: number;
}
```

2. **Response Interface:**
```typescript
export interface UpdateUserResponse {
  id: number;
  name: string;
  email: string;
  vscAccount: string;
  roleId: number;
  updatedAt: Date;
}
```

3. **Command Sınıfı:**
```typescript
export class UpdateUserCommand extends BaseCommand<UpdateUserResponse> {
  readonly type = "UpdateUserCommand";

  constructor(public readonly data: UpdateUserRequest) {
    super();
  }
}
```

4. **Handler:**
```typescript
export class UpdateUserHandler
  implements CommandHandler<UpdateUserCommand, UpdateUserResponse>
{
  private prisma: PrismaClient;

  constructor() {
    this.prisma = new PrismaClient();
  }

  async handle(command: UpdateUserCommand): Promise<UpdateUserResponse> {
    try {
      const user = await this.prisma.user.update({
        where: { id: command.data.id },
        data: {
          ...(command.data.name && { name: command.data.name }),
          ...(command.data.email && { email: command.data.email }),
          ...(command.data.vscAccount && { vscAccount: command.data.vscAccount }),
          ...(command.data.roleId && { roleId: command.data.roleId }),
        },
      });

      return {
        id: user.id,
        name: user.name,
        email: user.email,
        vscAccount: user.vscAccount,
        roleId: user.roleId,
        updatedAt: user.updatedAt,
      };
    } catch (error) {
      throw new Error(`Failed to update user: ${error}`);
    }
  }
}
```

5. **Registry'de Kayıt:**
```typescript
this.commandBus.register("UpdateUserCommand", new UpdateUserHandler());
```

## Best Practices

1. **Naming Convention:**
   - Command: `{Action}{Entity}Command` (örn: `CreateUserCommand`)
   - Query: `{Action}{Entity}Query` (örn: `GetUserQuery`)
   - Handler: `{Action}{Entity}Handler` (örn: `CreateUserHandler`)

2. **Error Handling:**
   - Tüm handler'larda try-catch kullanın
   - Anlamlı error mesajları döndürün
   - Error'ları log'layın

3. **Type Safety:**
   - Request ve Response interface'lerini her zaman tanımlayın
   - Generic type'ları doğru kullanın
   - TypeScript'in type checking özelliklerinden faydalanın

4. **Validation:**
   - Request data'sını validate edin
   - Business rule'ları handler'da kontrol edin

5. **Performance:**
   - Gereksiz database query'lerinden kaçının
   - Pagination kullanın (list query'lerde)
   - Caching stratejileri uygulayın

Bu dokümantasyon, CQRS pattern'ini kullanarak command ve query oluşturma sürecini kapsamlı bir şekilde açıklamaktadır. Her adım detaylı örneklerle desteklenmiştir.
