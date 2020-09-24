# Synchronizing in Go

This repository contains experiments/examples of syncing multiple goroutines.

## Examples

### Sync n goroutines with a single channel
In this example, we read a large file in chunks by using multiple goroutines.
The goroutines send the read bytes to the main goroutine over a single channel.
The main goroutine closes the channel while all bytes of the file have been read.

#### Flow
 - Get length of file
 - calculate n by dividing the length of the file with the configured load per goroutine
 - create channel
 - trigger n goroutines and inform them about which chunk to process
 - keep receiving from the channel until the amount of received bytes is bigger than or equal to the length of the file.
 - close the channel

To run this example: `$ go run cmd/single_channel/main.go`
average runtime for reading `resources/large_file.txt` is between `170ms` and `200ms`

### Sync n goroutines by merging n channels
In this example, we read a large file in chunks by using multiple goroutines.
Each goroutine has its own channel over which it sends the read bytes, which is closed when the goroutine terminates.
The main goroutine merges all n channels into a single one which it consumes and then closes.

#### Flow
 - Get length of file
 - calculate n by dividing the length of the file with the configured load per goroutine
 - create an empty slice of channels
 - trigger n goroutines, inform them about which chunk to process, give them a new channel and add that channel to the slice.
 - merge all channels in the slice into one (merge function handles closing of this channel)
 - keep receiving from the channel until it closes.

To run this example: `$ go run cmd/merge_channels/main.go`
average runtime for reading `resources/large_file.txt` is between `110ms` and `130ms`
