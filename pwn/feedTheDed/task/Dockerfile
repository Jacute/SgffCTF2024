FROM gcc:11.4 as builder
COPY src /usr/src/
WORKDIR /usr/src/
RUN gcc main.c -m64 -fno-stack-protector -no-pie -g -o feedTheDed

FROM debian:12-slim
RUN apt update && apt install socat -yyy
WORKDIR /app
COPY --from=builder /usr/src/feedTheDed .
RUN adduser user && \
        chmod 555 feedTheDed

USER user
CMD ["socat", "tcp-l:9002,reuseaddr,fork", "EXEC:'/app/feedTheDed'"]

EXPOSE 9002/tcp