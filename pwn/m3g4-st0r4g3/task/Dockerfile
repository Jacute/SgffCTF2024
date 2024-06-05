FROM ubuntu:22.04
RUN apt update && apt install socat -yyy
WORKDIR /app

COPY src/main.c .

RUN apt update && apt install gcc -y
RUN gcc main.c -m64 -fno-stack-protector -no-pie -g -o m3g4-st0r4g3

RUN adduser user && \
        chmod 555 m3g4-st0r4g3
USER user

CMD ["socat", "tcp-l:9000,reuseaddr,fork", "EXEC:'/app/m3g4-st0r4g3'"]

EXPOSE 9000/tcp