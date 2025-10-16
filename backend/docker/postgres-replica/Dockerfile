FROM postgres:17

COPY ./initdb/init-replica.sh /docker-entrypoint-initdb.d/
RUN chmod +x /docker-entrypoint-initdb.d/init-replica.sh

USER root
RUN apt-get update && apt-get install -y postgresql-client-17
USER postgres