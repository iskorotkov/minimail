FROM node:17-alpine3.14 as build
WORKDIR /src

RUN npm i -g pnpm

COPY ["client-spa/package.json", "client-spa/pnpm-lock.yaml", "./"]
RUN pnpm i

COPY ["./client-spa", "./"]
RUN REACT_APP_API_URL=http://localhost:8080/api pnpm run build

FROM node:17-alpine3.14 as run
WORKDIR /app

COPY --from=build /src/build /app/build
RUN npm i -g serve

EXPOSE 3000
ENTRYPOINT [ "serve", "-s", "build" ]
