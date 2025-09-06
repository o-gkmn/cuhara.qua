import { Command, CommandHandler } from "./command";

export class CommandBus {
  private static instance: CommandBus;
  private handlers = new Map<string, CommandHandler<any, any>>();

  private constructor() {}

  static getInstance(): CommandBus {
    if (!CommandBus.instance) {
      CommandBus.instance = new CommandBus();
    }
    return CommandBus.instance;
  }

  register<TCommand extends Command<TResponse>, TResponse>(
    commandName: string,
    handler: CommandHandler<TCommand, TResponse>
  ) {
    this.handlers.set(commandName, handler);
  }

  async execute<TCommand extends Command<TResponse>, TResponse>(
    command: TCommand
  ): Promise<TResponse> {
    const handler = this.handlers.get(command.type);
    if (!handler) {
      throw new Error(
        `Handler not found for command: ${command.type}`
      );
    }
    return handler.handle(command);
  }
}
