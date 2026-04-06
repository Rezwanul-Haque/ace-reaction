ARG GO_VERSION=1.25

# ---- Stage 1: Build React client ----
FROM node:24-alpine AS client-builder

WORKDIR /app
COPY client/package.json client/package-lock.json ./
RUN npm ci
COPY client/ ./
RUN npm run build

# ---- Stage 2: Build Go server ----
FROM golang:${GO_VERSION}-alpine AS server-builder

RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

RUN apk add --no-cache ca-certificates git

WORKDIR /src
COPY server/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix 'static' -o /server .

# ---- Stage 3: Final scratch image ----
FROM scratch AS final

COPY --from=server-builder /user/group /user/passwd /etc/
COPY --from=server-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=server-builder /server /server
COPY --from=client-builder /app/dist /public

USER nobody:nobody

EXPOSE 8080

ENTRYPOINT ["/server"]
