package main

import (
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/joho/godotenv/autoload"
	"github.com/memochou1993/eth-go-binding-example/contract"
)

func main() {
	conn, err := ethclient.Dial(os.Getenv("PROVIDER_URL"))
	if err != nil {
		log.Fatal(err)
	}

	todoList, err := contract.NewTodoList(common.HexToAddress(os.Getenv("CONTRACT_ADDRESS")), conn)
	if err != nil {
		log.Fatal(err)
	}

	task, err := todoList.Tasks(&bind.CallOpts{}, new(big.Int))
	if err != nil {
		log.Fatal(err)
	}

	log.Print(task)
}
