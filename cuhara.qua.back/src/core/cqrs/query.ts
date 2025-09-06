export interface Query<TResponse = void> {
  readonly type: string;
}

export interface QueryHandler<
  TQuery extends Query<TResponse>,
  TResponse = void
> {
  handle(query: TQuery): Promise<TResponse>;
}

export abstract class BaseQuery<TResponse = void> implements Query<TResponse> {
  abstract readonly type: string;
}
