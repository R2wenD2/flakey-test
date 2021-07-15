#!/usr/bin/env bash
count=`jq --slurp < out.json '.[] | select(.Action != "output") | select(.Action != "run") | select(.Test != null) | {(.Test): .Action}' | jq -s . | jq 'unique'| jq length`
if [ count != 2 ] 
then
  echo "At leaste one test is flakey"
  exit 1
 fi
