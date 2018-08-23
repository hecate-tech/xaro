echo "Using GOPATH=$GOPATH"

build () {
    buildMsg="Building Xaro"
    
    if [[ $2 == -verbose || $2 == -v  || $1 == -verbose || $1 == -v ]]; then
        buildMsg="$buildMsg Verbosely..."
        echo $buildMsg

        mage -v
    else
        buildMsg="$buildMsg Quietly..."
        echo $buildMsg
        
        mage
    fi
}

if [[ $1 == -build || $1 == -b ]]; then
    build $1 $2
fi

if [ ! -f ./bin/Xaro.exe ]; then
    echo "./bin/Xaro.exe was not found!"
    build $1 $2
fi

echo "Running Xaro..."

./bin/Xaro.exe