curl -s https://acad.learn2earn.ng/assets/superhero/all.json | jq -r '.[] | select(.id==70) | .name' | awk '{print "\"" $0 "\""}'
