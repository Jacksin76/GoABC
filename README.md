# Concurrent Communication with Goroutines and Channels

This Go program demonstrates concurrent execution and communication between goroutines using channels. It spawns three goroutines that print sequences of letters ("A", "B", "C") along with numbers in a synchronized manner. The synchronization is achieved using channels.

## Code Explanation

The code consists of three goroutines (`A`, `B`, `C`) that communicate with each other through channels. Here's a breakdown of the code:

1. The necessary packages are imported: `fmt`, `sync`, and `time`.

2. The `main` function is defined as the entry point of the program.

3. Channels `chA`, `chB`, and `chC` are created using the `make` function. These channels are used for communication between the goroutines.

4. Three goroutines are launched to execute concurrently:

   - Goroutine `A` prints the letter "A" along with numbers received from `chA`. It then sends data to `chB` for goroutine `B` to receive.

   - Goroutine `B` receives data from `chB` and prints the letter "B" along with the received numbers. It then sends data to `chC` for goroutine `C` to receive.

   - Goroutine `C` receives data from `chC` and prints the letter "C" along with the received numbers. It then sends data to `chA` for goroutine `A` to receive.

5. The `main` function waits for the completion of all goroutines using the `WaitGroup` (`wg`). This ensures that all goroutines finish execution before the program exits.

6. The channels `chA`, `chB`, and `chC` are closed using the `close` function after they are no longer needed.
