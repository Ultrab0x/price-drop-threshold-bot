# price-drop-threshold-bot

This script adds a strategy to remove poverty when the price is falling from a certain parameter. It sets the threshold drop price, and then, in a loop, it checks the current token price. If the current price is less than the price drop threshold, it calls the RemoveLiquidity function to remove liquidity from the exchange pool. Otherwise, it simply prints a message that the price is still above the threshold and no action will be taken. This loop runs every minute.
