ps -ef|grep HttpToolServer | grep -v grep | awk '{print $2}' | xargs kill -2 2>/dev/null
sleep 1
bash run.sh
ps -ef|grep HttpToolServer
