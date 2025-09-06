import { PrismaClient } from "@prisma/client";
import { BaseQuery, QueryHandler } from "../../../core/cqrs/query";

// ==================== REQUEST ====================
export interface GetUserRequest {
  userId: number;
}

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

// ==================== QUERY ====================
export class GetUserQuery extends BaseQuery<GetUserResponse> {
  readonly type = "GetUserQuery";

  constructor(public readonly data: GetUserRequest) {
    super();
  }
}

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
