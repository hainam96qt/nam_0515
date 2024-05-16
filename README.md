# Blockchain Implementation with Go and Ethereum

## Overview

This project involves implementing a basic blockchain in Go and using Ethereum smart contracts written in Solidity to store valid hash sequences. The goal is to combine the simplicity and efficiency of Go for blockchain development with the robust, decentralized storage capabilities of Ethereum.

## Project Structure

1. **Go Blockchain Implementation**
    - **Block Structure**: Defines the structure of a block containing fields such as index, timestamp, data, previous hash and hash.
    - **Hash Calculation**: Implements a function to calculate the hash of a block.
    - **Add Block**: Implements a function to add a new block to the blockchain.
    - **Consensus Algorithm**: Chooses and implements a consensus algorithm (e.g., Proof of Work).

2. **Ethereum Smart Contract (Solidity)**
    - **StringStorage Contract**: A smart contract written in Solidity to store and manage hash strings.
    - **Functions**:
        - `addString`: Adds a new hash string with a timestamp.
        - `getString`: Retrieves the timestamp associated with a hash string.
    - **Events**: Emits an event every time a new hash string is added.
    - **Access Control**: Ensures only the contract owner can add new hash strings.

3. **Interfacing Go with Ethereum**
    - **Contract Deployment**: Deploys the `StringStorage` smart contract within file `validhash.sol` to an Ethereum network.
    - **Contract Interaction**: Interacts with the deployed smart contract to add and retrieve hash strings.
    - **Event Listening**: Listens for specific events emitted by the smart contract.

## Getting Started

### Prerequisites

- Go 1.21.4+
- Node.js and npm (for running an Ethereum node like Ganache)
- Solidity Compiler (`solc`)
- abigen (for generating Go bindings from Solidity contracts)
- An Ethereum client (e.g., Ganache)

### Deploying the Smart Contract

1. Compile the Solidity contract:
    ```bash
    cd ./internal/repo/smartcontract
    solc --bin validhash.sol > validhash.bin
    solc --abi validhash.sol > validhash.abi
    abigen --abi=build/validhash.abi --bin=build/validhash.bin --pkg=contracts --out=validhash.go
    ```

2. Run an Ethereum node (Ganache):
    ```bash
    ganache
    ```

3. Run datatabase :
    ```bash
    docker-compose up -d
    ```
### Running Tests
- Run server at local
   ```bash
    go run main.go
   ```
### Running Tests

- Run the tests using:
    ```bash
    go test ./...
    ```