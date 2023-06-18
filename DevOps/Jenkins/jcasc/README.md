<a href="https://jenkins.io">
  <picture>
    <source width="400" media="(prefers-color-scheme: dark)" srcset="https://www.jenkins.io/images/jenkins-logo-title-dark.svg">
    <img width="400" src="https://www.jenkins.io/images/jenkins-logo-title.svg">
  </picture>
</a>

# Set up Jenkins with docker compose

## Create a dockerfile with name "Dockerfile"

In this file:

```
FROM jenkins/jenkins:latest // install jenkins latest

// set up config, casc.yaml is jenkins's config file
ENV JAVA_OPTS -Djenkins.install.runSetupWizard=false
ENV CASC_JENKINS_CONFIG /var/jenkins_home/casc.yaml

// install plugins
COPY --chown=jenkins:jenkins plugins.txt /usr/share/jenkins/ref/plugins.txt
RUN jenkins-plugin-cli -f /usr/share/jenkins/ref/plugins.txt

// copy your config in casc.yaml to jenkins's casc.yaml
COPY casc.yaml /var/jenkins_home/casc.yaml

```

## Create a docker compose with name "docker-compose.yml"

In this file:

```
version: "3.8"
services:
  jenkins:
    privileged: true
    user: root
    // when your server restart your jenkins will be ran
    restart: always
    hostname: "localhost"
    container_name: jenkins_config
    // access to loacalhost:9000 => your jenkins
    ports:
      - "9000:80"
    // path to dockerfile
    build: .
    // create admin's account
    environment:
      - "JENKINS_ADMIN_ID=admin"
      - "JENKINS_ADMIN_PASSWORD=minhthu"
    // the place, your data (jenkins) will be stored
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock

```

## How to run

- Step 1: Access your folder you have stored your docker compose file
- Step 2: Run this line:

```
docker-compose up -d // In some cases, you must use "sudo docker-compose up -d"
```

- Step 3: Access your jenkins's website with url http://localhost:9000 with account

```
username: admin
password: minhthu
```
