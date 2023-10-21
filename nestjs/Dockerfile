FROM node:20-slim AS development

WORKDIR /home/node/app

RUN apt update -y && \
    apt install -y procps && \
    npm install -g @nestjs/cli@10.0.0

USER node

CMD npm ci && \
    npm run start:dev && \
    tail -f /dev/null