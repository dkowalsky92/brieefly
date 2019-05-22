FROM mysql:latest

RUN apt-get update && \
    apt-get install jq bash -y wget && \
    rm -rf /var/lib/apt/lists/*

RUN mkdir /db
WORKDIR /db
ADD . .

RUN chmod +x ./mysql-create.sh ; \
    chmod +x ./get-json-value.sh

SHELL ["/bin/bash", "-c"]

RUN . "./mysql-create.sh <(get-json-value.sh ../run/secrets/config-local.json .database.name) \
    <(get-json-value.sh ../run/secrets/config-local.json" ".database.user) \
    <(get-json-value.sh ../run/secrets/config-local.json" ".database.password)" 

