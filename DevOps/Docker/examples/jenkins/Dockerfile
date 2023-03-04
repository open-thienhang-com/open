FROM jenkins/jenkins
USER root
# máy MT chạy ko đc 2 lệnh này nên comment rồi ai chạy thử thì mở chạy ạ
# RUN apt-get update && apt-get install -g docker-ce-cli
# RUN apt-get update && apt-get install -g lsb-release
RUN echo "deb [arch=$(dpkg --print-architecture) \
    signed-by=/usr/share/keyrings/docker-archive-keyring.asc] \
    https://download.docker.com/linux/debian \
    $(lsb_release -cs) stable" > /etc/apt/sources.list.d/docker.list
RUN curl -fsSLo /usr/share/keyrings/docker-archive-keyring.asc \
    https://download.docker.com/linux/debian/gpg
USER jenkins
# run dockerfile: docker build -t jenkins-docker .
# run local: docker run -p 8080:8080 -p 50000:50000 --restart=on-failure -d -v jenkins_home:/var/jenkins_home jenkins-docker
# xem container id: docker ps
# log container để có pass: docker logs [id contianer]ex: docker logs 5b530816beaf
# ref https://sweetcode.io/how-to-build-and-push-docker-images-to-docker-hub-using-jenkins-pipeline/