FROM golang:1.18.0-alpine3.15 as builder

# Force the go compiler to use modules
ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org
# Indicates module paths that are not publicly available
ENV GOPRIVATE=bitbucket.org/sdssc

# Update OS package and install Git
RUN apk update && apk add git openssh && apk add build-base && apk add tzdata

# Setup github credential
#ADD ./resources/docker/keys/id_rsa /root/.ssh/id_rsa
#ADD ./resources/docker/keys/id_rsa.pub /root/.ssh/id_rsa.pub
#RUN chmod 600 /root/.ssh/id_rsa


# make sure your domain is accepted
#RUN touch /root/.ssh/known_hosts
#RUN ssh-keyscan bitbucket.org >> /root/.ssh/known_hosts
#RUN git config --global url."git@bitbucket.org:".insteadOf "https://bitbucket.org"

WORKDIR /go/src/bitbucket.org/sdssc/dooflix-api-v3
ADD go.mod go.mod
ADD go.sum go.sum
ADD app app
# ADD resources/cert/oauth-public.key oauth-public.key
# ADD resources/cert/oauth-private.key oauth-private.key


# Set Docker's entry point commands
#RUN cd app/ && go build -o /go/bin/app.bin


# Install Fresh for local development
#RUN go get github.com/pilu/fresh

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://proxy.golang.org
RUN go env -w GOPRIVATE=bitbucket.org/sdssc
RUN go mod vendor
RUN cd app/ && go build -o /go/bin/app.bin -mod vendor -tags "musl static_all"

# ------------------------------------------------------------------------------
# Deployment image
# ------------------------------------------------------------------------------
FROM golang:1.18.0-alpine3.15
COPY --from=builder /go/bin/app.bin /app/app.bin
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Bangkok

ADD entrypoint.sh /entrypoint.sh

RUN adduser -u 1001 -D -s /bin/sh -G ping 1001
RUN chown 1001:1001 /entrypoint.sh
RUN chown 1001:1001 /app/app.bin
RUN chmod +x /entrypoint.sh
RUN chmod +x /app/app.bin

USER 1001

EXPOSE 3000

ENTRYPOINT ["/entrypoint.sh"]