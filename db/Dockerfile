FROM postgres:13-alpine

COPY init.sql /docker-entrypoint-initdb.d/

ENV POSTGRES_USER=kennytest
ENV POSTGRES_PASSWORD=password
ENV POSTGRES_DB=test_db

EXPOSE 5432