!#/bin/bash

git add .
git commit -m "New checkpoint! $(date '+%a %M:%H %h %d %Y')"
git push -u origin main
