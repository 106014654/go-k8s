#构建本地镜像
sudo docker build -t httpquery:v1 .
Sending build context to Docker daemon  7.469MB
Step 1/4 : FROM ubuntu
 ---> 216c552ea5ba
Step 2/4 : RUN  mkdir -p /dockerstudy/httpserver
 ---> Using cache
 ---> b2908472b941
Step 3/4 : COPY go /dockerstudy/httpserver/
 ---> Using cache
 ---> c571646ae6e0
Step 4/4 : ENTRYPOINT "./dockerstudy/httpserver/go"
 ---> Using cache
 ---> 5e509b4fe62e
Successfully built 5e509b4fe62e
Successfully tagged httpquery:v1


ubuntu@VM-0-114-ubuntu:~/go$ sudo docker images
REPOSITORY        TAG         IMAGE ID       CREATED          SIZE
ccjj1091/go-k8s   httpquery   5e509b4fe62e   36 minutes ago   85.3MB
httpquery         latest      5e509b4fe62e   36 minutes ago   85.3MB
httpquery         v1          5e509b4fe62e   36 minutes ago   85.3MB
ubuntu            latest      216c552ea5ba   11 days ago      77.8MB

#编写 Dockerfile 将模块二作业编写的 httpserver 容器化
#通过 docker 命令本地启动 httpserver
sudo docker run -p 80:8080 httpquery:v1



#将镜像推送至 docker 官方镜像仓库

sudo docker tag 5e509b4fe62e ccjj1091/go-k8s:httpquery
sudo docker push ccjj1091/go-k8s:httpquery





#通过 nsenter 进入容器查看 IP 配置
sudo docker inspect 8a7eb0507208 | grep IPAddress

            "SecondaryIPAddresses": null,
            "IPAddress": "172.17.0.2",
                    "IPAddress": "172.17.0.2",
