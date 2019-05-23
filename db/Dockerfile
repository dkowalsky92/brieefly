FROM mysql:latest

ADD ./sql-scripts /docker-entrypoint-initdb.d
RUN mkdir /secrets
COPY ../secrets/ /secrets
RUN apt-get update && \
    apt-get install jq bash -y wget && \
    rm -rf /var/lib/apt/lists/*

RUN cd /run ; ls -a

RUN mkdir /db
WORKDIR /db
ADD . .

RUN chmod +x ./mysql-create.sh ; \
    chmod +x ./get-json-value.sh

SHELL ["/bin/bash", "-c"]

RUN . "./mysql-create.sh <(get-json-value.sh ../run/secrets/config_local.json .database.name) \
    <(get-json-value.sh ../run/secrets/config_local.json" ".database.user) \
    <(get-json-value.sh ../run/secrets/config_local.json" ".database.password)"

