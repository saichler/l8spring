set -e
cd ./spring/vnet/
./build.sh
cd ../main
./build.sh
cd ../ui
./build.sh
