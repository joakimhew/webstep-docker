# Maven build container 

FROM maven:3.5.2-jdk-8-alpine AS maven_build
COPY pom.xml /tmp/
WORKDIR /tmp/
RUN mvn package
COPY src /tmp/src/

FROM openjdk:8-jdk-alpine
EXPOSE 8080
COPY --from=maven_build /tmp/target/hello-world-0.1.0.jar /data/hello-world-0.1.0.jar

CMD java -jar /data/hello-world-0.1.0.webstep run -p 8080:8080 -it --rm hello-world-java