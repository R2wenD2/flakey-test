if [jq --slurp < out.json '.[] | select(.Action != "output") | select(.Action != "run") | select(.Test != null) | {(.Test): .Action}' | jq -s . | jq 'unique'| jq length] > 2
then
  echo fail
fi
  
