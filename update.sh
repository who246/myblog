git pull
ps auxwww | grep myblog | grep -v grep | awk '{print $2}' | xargs kill -9
nohup ./myblog >/dev/null 2>&1&