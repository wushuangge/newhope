## docker-compose
cd taskdash/docker

#### 启动容器
sudo docker-compose -f docker-compose.yml start

#### 停止容器运行
sudo docker-compose -f docker-compose.yml stop

#### 重启容器
sudo docker-compose -f docker-compose.yml restart

#### 重新构建容器
sudo docker-compose -f docker-compose.yml up -d

#### 删除容器
sudo docker-compose -f docker-compose.yml down

#### 查看容器
sudo docker-compose -f docker-compose.yml ps

## docker
cd taskdash

#### docker镜像构建
sudo docker build . -t docker_taskdash -f docker/Dockerfile

#### docker镜像移除 
sudo docker rmi -f IMAGE_ID

#### 以交互模式启动容器
sudo docker run -it -p 8080:8080 CONTAINER_NAME

#### 以附加进程方式启动容器
sudo docker run -d -p 8080:8080 CONTAINER_NAME

#### 停止一个正在运行的容器
sudo docker stop CONTAINER_ID

sudo docker kill CONTAINER_ID

#### 重启一个容器
sudo docker ps

sudo docker restart CONTAINER_ID

## docker 查看日志
    $ docker logs [OPTIONS] CONTAINER
      Options:
            --details        显示更多的信息
        -f, --follow         跟踪实时日志
            --since string   显示自某个timestamp之后的日志，或相对时间，如42m（即42分钟）
            --tail string    从日志末尾显示多少行日志， 默认是all
        -t, --timestamps     显示时间戳
            --until string   显示自某个timestamp之前的日志，或相对时间，如42m（即42分钟）
            
sudo docker logs --since 30m CONTAINER_ID

sudo docker logs -f CONTAINER_ID

