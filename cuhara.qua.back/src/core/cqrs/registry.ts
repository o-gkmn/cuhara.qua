import { CommandBus } from "./command-bus";
import { QueryBus } from "./query-bus";

// User handlers
import { CreateUserHandler } from "../../modules/users/commands/create-user";
import { GetUserHandler } from "../../modules/users/queries/get-user";

export class CQRSRegistry {
  private static instance: CQRSRegistry;
  private commandBus: CommandBus;
  private queryBus: QueryBus;

  private constructor() {
    this.commandBus = CommandBus.getInstance();
    this.queryBus = QueryBus.getInstance();
    this.registerHandlers();
  }

  static getInstance(): CQRSRegistry {
    if (!CQRSRegistry.instance) {
      CQRSRegistry.instance = new CQRSRegistry();
    }
    return CQRSRegistry.instance;
  }

  private registerHandlers(): void {
    // Register User Command Handlers
    this.commandBus.register("CreateUserCommand", new CreateUserHandler());

    // Register User Query Handlers
    this.queryBus.register("GetUserQuery", new GetUserHandler());
  }

  getCommandBus(): CommandBus {
    return this.commandBus;
  }

  getQueryBus(): QueryBus {
    return this.queryBus;
  }
}
