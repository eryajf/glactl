FROM docker.cnb.cool/znb/images/alpine

LABEL maintainer=eryajf

ENV TZ=Asia/Shanghai
ENV BINARY_NAME=glactl

ARG TARGETOS
ARG TARGETARCH

COPY bin/${BINARY_NAME}_${TARGETOS}_${TARGETARCH} /usr/local/bin/${BINARY_NAME}

RUN chmod +x /usr/local/bin/${BINARY_NAME}