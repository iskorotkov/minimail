FROM node:17-alpine3.14 as build
WORKDIR /src

RUN npm i -g pnpm

COPY ["package.json", "pnpm-lock.yaml", "./"]
RUN pnpm i

COPY ["./", "./"]
RUN pnpm run build

FROM node:17-alpine3.14 as run
WORKDIR /app

COPY --from=build /src/build /app/build
RUN npm i -g serve

EXPOSE 3000
ENTRYPOINT [ "serve", "-s", "build" ]
