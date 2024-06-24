FROM ubuntu:22.04
RUN apt update && apt install socat -yyy
WORKDIR /app

COPY src/main.c .
COPY flag /flag

RUN apt update && apt install gcc -y
RUN gcc -z execstack -o sh3llc0d3r main.c

RUN adduser user && \
        chmod 555 sh3llc0d3r
USER user

CMD ["socat", "tcp-l:9000,reuseaddr,fork", "EXEC:'/app/sh3llc0d3r'"]

EXPOSE 9000/tcp