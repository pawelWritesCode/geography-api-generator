echo "Building generator,"
go build main.go

echo "Moving binary to binary folder,"
mv main $GOPATH/bin/aiqa-generator2

echo "Changing privilages to binary."
sudo chmod 741 $GOPATH/bin/aiqa-generator2