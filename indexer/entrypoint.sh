#!/bin/bash

envsubst < s3-to-solr.yaml > s3-to-solr-envsubst.yaml
embulk guess ./s3-to-solr-envsubst.yaml -o guessed.yaml
embulk run guessed.yaml
