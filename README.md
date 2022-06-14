# Portal
Portal  is the backend service for bridge.konstellation.tech.

Checkout v1.0

```
git checkout v1.0
```

Makefile

```
make build
```

```
docker-compose up --build -d
```

install solidity
```
https://docs.soliditylang.org/en/develop/installing-solidity.html#binary-packages
```

For DRC & DARC token use 0.6.8 version 
```
git clone https://github.com/ethereum/homebrew-ethereum.git
cd homebrew-ethereum
git checkout 16b4241f874f75fc0210fefc59de989e94476ae8
brew unlink solidity
brew install solidity.rb
```

```
./build/portal start

```
To run in the background
```
screen -dmSL pt ./build/portal start
```

open localhost:7081
and create new database `portal`
