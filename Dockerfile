FROM golang:1.17.8-alpine AS builder

# 环境
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"
#移动到工作目录
WORKDIR  /workspace

COPY .  .
WORKDIR  /workspace/cmd/bkAdmin
RUN go build -o bkAdmin main.go

FROM alpine as final

#设置时区
RUN apk add --no-cache tzdata
ENV TZ="Asia/Shanghai"

# 移动到新存放二进制文件的目录上
WORKDIR /opt/app

#从原目录复制到当前目录
COPY --from=builder  /workspace/cmd/bkAdmin .

# 在容器工作目录下的/opt/app创建一个目录为configs
RUN mkdir configs

#将配置文件复制到
COPY --from=builder /workspace/configs ./configs
COPY --from=builder /workspace/logs ./logs

EXPOSE 8070
CMD ["/opt/app/bkAdmin", "run","-c", "configs/config.yaml"]


# docker build -t bkadmin:v1 .

#docker run -d --name bkadmin  \
#-v     /home/ipfs/red/workspace/data/bkAdmin/config.yaml:/opt/app/configs/config.yaml   \
#-v    /home/ipfs/red/workspace/data/bkAdmin/logs:/opt/app/logs   \
#-p   8070:8070  \
#bkadmin:v1