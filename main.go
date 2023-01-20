package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/uniswap/uniswap-go/types"
)

func main() {
	// Connect to a local Ethereum node
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	// Define the Uniswap exchange address and token addresses
	exchangeAddress := "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"
	tokenAddress1 := "0x6B175474E89094C44Da98b954EedeAC495271d0F"
	tokenAddress2 := "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"

	// Create a new Uniswap exchange instance
	exchange, err := types.NewExchange(exchangeAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// Add liquidity to the exchange
	tx, err := exchange.AddLiquidity(
		context.Background(),
		tokenAddress1,
		tokenAddress2,
		1000000000000000000, // 1 token1
		1000000000000000000, // 1 token2
		0,                    // min liquidity
		0,                    // deadline
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Transaction hash:", tx.Hash().Hex())

	// Define the price drop threshold
	priceDropThreshold := big.NewFloat(1.1)

	for {
		// Get the current token price
		tokenPrice, err := exchange.TokenPrice(context.Background(), tokenAddress1)
		if err != nil {
			log.Fatal(err)
		}

		// Check if the price has dropped below the threshold
		if tokenPrice.Cmp(priceDropThreshold) < 0 {
			// Remove liquidity from the exchange
			tx, err := exchange.RemoveLiquidity(
				context.Background(),
				1000000000000000000, // 1 token1
				1000000000000000000, // 1 token2
				0,                    // min liquidity
				0,                    // deadline
			)
			if err != nil {
				log.Fatal(err)
			}
	
			fmt.Println("Transaction hash:", tx.Hash().Hex())
			fmt.Println("Price dropped below threshold, removing liquidity")
			break

		} else {
		fmt.Println("Price is still above threshold, no action taken")
		}
				// Wait for a certain period before checking the price again
		time.Sleep(60 * time.Second) // in this case it will check the price every minute
	}
}
