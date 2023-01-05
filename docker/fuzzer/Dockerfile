FROM docker.io/library/golang:1.16.4
RUN apt update && apt install -y python3
WORKDIR /gfuzz

# copy source files to docker
COPY patch ./patch
COPY scripts ./scripts
COPY pkg ./pkg
COPY cmd ./cmd
COPY docker/fuzzer/entrypoint.sh docker/fuzzer/entrypoint.sh
COPY go.mod go.sum Makefile VERSION ./
# build inst and fuzzer
RUN http_proxy=http://proxy-mtl.ubisoft.org:3128 https_proxy=http://proxy-mtl.ubisoft.org:3128 make tidy
RUN ls -al
RUN BUILD=docker make

# patch golang runtime in the container
RUN chmod +x ./scripts/patch.sh
RUN echo $GOROOT
RUN cat ./scripts/patch.sh | sed '1 s/\r$//' > ./scripts/patch.sh
RUN ./scripts/patch.sh

RUN chmod +x docker/fuzzer/entrypoint.sh
RUN ls /gfuzz/docker/fuzzer
RUN cat /gfuzz/docker/fuzzer/entrypoint.sh

# RUN groupadd gfgroup
# RUN useradd -r -u 1001 -g gfgroup gfuser
# RUN chown gfuser:gfgroup ./scripts/fuzz.sh && chmod +x ./scripts/fuzz.sh
# USER gfuser
# RUN chmod +x ./scripts/fuzz.sh
# ENTRYPOINT [ "scripts/fuzz.sh" ] 

# ENTRYPOINT ["/bin/sh", "/gfuzz/docker/fuzzer/entrypoint.sh" ]
CMD ["bash"]