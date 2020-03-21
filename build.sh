echo "Building generator,"
go build main.go

echo "Moving binary to bin folders,"
cp main ./bin/geo-generator
mv main $GOPATH/bin/geo-generator

echo "Changing privilages to binary file."
sudo chmod 755 ./bin/geo-generator
sudo chmod 755 $GOPATH/bin/geo-generator