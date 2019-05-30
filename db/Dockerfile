FROM mysql:latest

EXPOSE 3306

ENV MYSQL_ALLOW_EMPTY_PASSWORD=yes

RUN apt-get update && \
    apt-get install jq -y wget && \
    rm -rf /var/lib/apt/lists/*

COPY ./init /docker-entrypoint-initdb.d/
WORKDIR /brieefly
ADD . .