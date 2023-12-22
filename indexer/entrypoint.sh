#!/bin/bash

envsubst < s3-to-elasticsearch.yaml > s3-to-elasticsearch-envsubst.yaml
embulk guess ./s3-to-elasticsearch-envsubst.yaml -o guessed.yaml
embulk run guessed.yaml
