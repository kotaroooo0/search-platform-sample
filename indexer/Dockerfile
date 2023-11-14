FROM openjdk:8

RUN wget -q https://dl.embulk.org/embulk-latest.jar -O /usr/local/bin/embulk \
    && chmod +x /usr/local/bin/embulk \
    && embulk gem install embulk-input-s3 \
    && embulk gem install embulk-output-elasticsearch \
    && apt-get update \
    && apt-get install gettext-base -y

WORKDIR /workspace

COPY s3-to-elasticsearch.yaml .
COPY entrypoint.sh .

ENTRYPOINT [ "/workspace/entrypoint.sh" ]
