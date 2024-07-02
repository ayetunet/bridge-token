#!/bin/bash

# Ensure npm is initialized
if [ ! -f package.json ]; then
    echo "Initializing npm..."
    npm init -y
fi

# Install OpenZeppelin contracts
echo "Installing OpenZeppelin contracts..."
npm install @openzeppelin/contracts @openzeppelin/contracts-upgradeable

# Compile the Solidity contract
echo "Compiling Solidity contract..."
solc --abi --bin --overwrite -o build --include-path node_modules/ --base-path . contracts/AyetuToken.sol

if [ $? -eq 0 ]; then
    echo "Compilation successful."
else
    echo "Compilation failed."
fi
