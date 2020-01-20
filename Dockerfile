FROM adoptopenjdk/openjdk11
VOLUME /tmp
COPY build/libs/*.jar app.jar
EXPOSE 8080 443
ENTRYPOINT exec java $JAVA_AGENT $JAVA_OPTS $JMX_OPTS -jar app.jar