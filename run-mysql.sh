echo "Container delete"

docker rm -f starbucks-mysql

echo "Run mysql..."

docker run --name starbucks-mysql \
        -p 3306:3306 \
        -e MYSQL_ROOT_PASSWORD=root123 \
        -e MYSQL_DATABASE=starbucks \
        -e MYSQL_USER=user \
        -e MYSQL_PASSWORD=pass \
        -e TZ=Asia/Seoul \
        -v /c/Users/user/Documents/Workspace/MYWEB/mysql:/var/lib/mysql \
        -d mysql:5.7

echo "Success!!"
