import { BaseCommand, CommandHandler } from "../../../core/cqrs/command";
import { PrismaClient } from "@prisma/client";
import { prisma } from "../../../prisma/client";

// ==================== REQUEST ====================
export interface CreateUserRequest {
  name: string;
  email: string;
  vscAccount: string;
  roleId: number;
  tenantId: number;
}

// ==================== RESPONSE ====================
export interface CreateUserResponse {
  id: number;
}

// ==================== COMMAND ====================
export class CreateUserCommand extends BaseCommand<CreateUserResponse> {
  readonly type = "CreateUserCommand";

  constructor(public readonly data: CreateUserRequest) {
    super();
  }
}

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
