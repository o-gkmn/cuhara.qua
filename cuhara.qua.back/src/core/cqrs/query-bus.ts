import { Query, QueryHandler } from "./query";

export class QueryBus {
  private static instance: QueryBus;
  private handlers = new Map<string, QueryHandler<any, any>>();

  private constructor() {}

  static getInstance(): QueryBus {
    if (!QueryBus.instance) {
      QueryBus.instance = new QueryBus();
    }
    return QueryBus.instance;
  }

  register<TQuery extends Query<TResponse>, TResponse>(
    queryName: string,
    handler: QueryHandler<TQuery, TResponse>
  ) {
    this.handlers.set(queryName, handler);
  }

  async execute<TQuery extends Query<TResponse>, TResponse>(
    query: TQuery
  ): Promise<TResponse> {
    const handler = this.handlers.get(query.type);

    if (!handler) {
      throw new Error(`Handler not found for query: ${query.type}`);
    }
    return handler.handle(query);
  }
}
