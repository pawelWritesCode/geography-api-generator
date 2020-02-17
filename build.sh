echo "Building generator,"
go build main.go

echo "Moving binary to binary folder,"
mv main $GOPATH/bin/geo-generator

echo "Changing privilages to binary file."
sudo chmod 741 $GOPATH/bin/geo-generator