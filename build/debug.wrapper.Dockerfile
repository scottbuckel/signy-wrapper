FROM notary-binary as wrapper-binary

# build notary-wrapper binary
ARG NOTARY_WRAPPER_BRANCH
ARG NOTARY_WRAPPER_PKG

RUN git clone -b $NOTARY_WRAPPER_BRANCH https://${NOTARY_WRAPPER_PKG}.git /go/src/${NOTARY_WRAPPER_PKG}

WORKDIR /go/src/${NOTARY_WRAPPER_PKG}

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags "-w -s -X ${NOTARY_WRAPPER_PKG}/version.GitCommit=`git rev-parse --short HEAD` -X ${NOTARY_WRAPPER_PKG}/version.NotaryWrapperVersion=`cat NOTARY_WRAPPER_VERSION`" -o /go/bin/signy-wrapper ${NOTARY_WRAPPER_PKG}/cmd

#cd ~/Work/signy-wrapper/ && rm -rf bin/ && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s -X github.com/scottbuckel/signy-wrapper/version.GitCommit=`git rev-parse --short HEAD` -X github.com/scottbuckel/signy-wrapper/version.NotaryWrapperVersion=`cat NOTARY_WRAPPER_VERSION`" -o ./bin/signy-wrapper github.com/scottbuckel/signy-wrapper/cmd
#cd ~/Work/signy-wrapper/ && rm -f bin/signy-wrapper && go build -ldflags "-w -s -X github.com/scottbuckel/signy-wrapper/version.GitCommit=`git rev-parse --short HEAD` -X github.com/scottbuckel/signy-wrapper/version.NotaryWrapperVersion=`cat NOTARY_WRAPPER_VERSION`" -o ./bin/signy-wrapper github.com/scottbuckel/signy-wrapper/cmd && cd bin && ./signy-wrapper

FROM golang:1.14.4-alpine3.12

COPY --from=notary-binary /.notary /.notary
COPY --from=notary-binary /user/group /user/passwd /etc/
COPY --from=notary-binary /go/bin/notary /notary/notary
COPY --from=wrapper-binary /go/bin/signy-wrapper /notary/signy-wrapper
COPY --from=wrapper-binary /etc/ssl /etc/ssl


## start local only (docker run ...)
# use secret for k8s
COPY notary-wrapper.crt  /etc/certs/notary/notary-wrapper.crt

COPY notary-wrapper.key  /etc/certs/notary/notary-wrapper.key
RUN chown -R notary:notary /etc/certs/notary
RUN chown -R notary:notary /notary
RUN chown -R notary:notary /.notary

#docker run sebbyii/signy-wrapper:0.0.1
## end local only


ENV NOTARY_PORT "4445"
ENV NOTARY_CERT_PATH "/etc/certs/notary"
ENV NOTARY_ROOT_CA "root-ca.crt"
ENV NOTARY_CLI_PATH "/notary/notary"

USER notary:notary

EXPOSE 4445

WORKDIR /notary

ENTRYPOINT [ "/notary/signy-wrapper" ]
