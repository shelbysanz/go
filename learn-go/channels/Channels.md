### Quick Channel Notes

1. Sending to a nil channel blocks the sender forever
2. Reading from a nil channel block the reader forever
3. A send to a closed channel "Panics" the program
4. A recieve from a closed channel returns a 0 value instantly
