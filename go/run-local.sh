set -e

# clean up
rm -rf go.sum
rm -rf go.mod
rm -rf vendor

# fetch dependencies
go mod init
GOPROXY=direct GOPRIVATE=github.com go mod tidy
go mod vendor

docker rm -f unsecure-postgres 2>/dev/null || true
sudo mkdir -p /data/postgres && sudo chmod 777 /data /data/postgres 2>/dev/null || true
docker run -d --name unsecure-postgres -p 5432:5432 -v /data/:/data/ \
  --entrypoint /bin/sh \
  saichler/unsecure-postgres:latest \
  -c "/start-postgres.sh admin admin admin 5432 && tail -f /dev/null"

rm -rf demo
mkdir -p demo
cd tests/mocks/cmd
echo "Building Mocks"
go build -o ../../../demo/mocks_demo
cd ../../../spring/vnet/
echo "Building vnet"
go build -o ../../demo/vnet_demo
cd ../main
echo "Building Spring"
go build -o ../../demo/spring_demo
cd ../ui/main
echo "Building ui"
go build -o ../../../demo/ui_demo
cd ..
cp -r ./web ../../demo/.
cd ../../demo

echo "cd .." > ./kill_demo.sh
echo "rm -rf demo" >> ./kill_demo.sh
echo "rm -rf /data/postgres/admin" >> ./kill_demo.sh
echo "pkill -9 demo" >> ./kill_demo.sh
chmod +x ./kill_demo.sh

./vnet_demo &
sleep 1
./spring_demo local &
sleep 10
./ui_demo &
sleep 8
EXTERNAL_IP=$(ip route get 1.1.1.1 | grep -oP 'src \K[0-9.]+')
read -p "Press Enter to upload mocks"
./mocks_demo --address https://${EXTERNAL_IP}:2773 --user admin --password admin --insecure

read -p "Press Enter to kill the demo"
./kill_demo.sh
