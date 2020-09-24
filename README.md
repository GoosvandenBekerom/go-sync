# Synchronizing in Go

This repository contains experiments/examples of syncing multiple goroutines.

## Examples

### Sync n goroutines with a single channel
In this example, we read a large file in chunks by using multiple goroutines.
The goroutines send the read bytes to the main goroutine over a single channel.
The main goroutine closes the channel while all bytes of the file have been read.

#### flow
 - Get length of file
 - calculate n by dividing the length of the file with the configured load per goroutine
 - create channel
 - trigger n goroutines and inform them about which chunk to process
 - keep receiving from the channel until the amount of received bytes is bigger than or equal to the length of the file.
 - close the channel

To run this example: `$ go run cmd/single_channel/main.go`
