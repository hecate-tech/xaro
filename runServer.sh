echo "Using GOPATH=$GOPATH"

build () {
    buildMsg="Building Server"
    
    if [[ $2 == -verbose || $2 == -v  || $1 == -verbose || $1 == -v ]]; then
        buildMsg="$buildMsg Verbosely..."
        echo $buildMsg

        go build -o bin/Server.exe src/communication/server/main.go
    else
        buildMsg="$buildMsg Quietly..."
        echo $buildMsg
        
        go build -o bin/Server.exe src/communication/server/main.go
    fi
}

if [[ $1 == -build || $1 == -b ]]; then
    build $1 $2
else
    if [[ $1 == -clean || $1 == -c ]]; then
        rm ./bin/Server.exe
        exit 1
    fi
fi

if [ ! -f ./bin/Server.exe ]; then
    echo "./bin/Server.exe was not found!"
    build $1 $2
fi

echo "Running Server..."

./bin/Server.exe