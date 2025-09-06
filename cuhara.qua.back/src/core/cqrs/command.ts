export interface Command<TResponse = void> {
  readonly type: string;
}

export interface CommandHandler<
  TCommand extends Command<TResponse>,
  TResponse = void
> {
  handle(command: TCommand): Promise<TResponse>;
}

export abstract class BaseCommand<TResponse = void> implements Command<TResponse> {
  abstract readonly type: string;
}
