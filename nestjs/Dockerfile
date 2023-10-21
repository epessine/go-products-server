FROM node:20-slim AS development

WORKDIR /home/node/app

RUN apt update -y && \
    apt install -y procps && \
    npm install -g @nestjs/cli@10.0.0

USER node

CMD [ "tail", "-f", "/dev/null" ]

FROM development AS build

CMD [ "npm", "ci" ]

FROM build AS serve

CMD [ "npm", "run", "start:dev" ]