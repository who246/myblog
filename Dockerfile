FROM ubuntu:trusty
MAINTAINER alaudadoc alaudadoc@alauda.cn

RUN apt-get update && apt-get install -y Go
EXPOSE 80
COPY myblog /
CMD ["nohup", "./myblog&"]
