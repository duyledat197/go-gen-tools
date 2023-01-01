FROM postgres:13-alpine

RUN \
    apk --no-cache add \
    build-base \
    clang-dev \
    llvm15 \
    lz4-dev \
    msgpack-c-dev \
    zstd-dev

ENV PGROONGA_VERSION=2.4.2 \
    GROONGA_VERSION=12.1.0

WORKDIR /
COPY ./groonga-build.sh /
RUN \
    /groonga-build.sh ${PGROONGA_VERSION} ${GROONGA_VERSION} && \
    rm -f groonga-build.sh

RUN \
    apk del \
    build-base \
    clang \
    clang-dev \
    llvm15

EXPOSE 5432
CMD ["postgres"]