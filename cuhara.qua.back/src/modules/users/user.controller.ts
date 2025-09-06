import { Request, Response } from "express";
import { CommandBus } from "../../core/cqrs/command-bus";
import { QueryBus } from "../../core/cqrs/query-bus";
import { CreateUserCommand } from "./commands/create-user";
import { GetUserQuery } from "./queries/get-user";

export class UserController {
  private commandBus: CommandBus;
  private queryBus: QueryBus;

  constructor() {
    this.commandBus = CommandBus.getInstance();
    this.queryBus = QueryBus.getInstance();
  }

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

}