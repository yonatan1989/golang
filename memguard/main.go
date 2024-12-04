package main

import (
    "fmt"
    "log"

    "github.com/awnumar/memguard"
)

func main() {
    // Ensure that all operations involving sensitive data are done securely
    //defer memguard.DestroyAll()

    // Create a new Enclave with some sensitive data
    sensitiveData := []byte("my super secret data")
    enclave := memguard.NewEnclave(sensitiveData)

    // Safe release of the sensitive data from memory
    //defer enclave.Destroy()

    // Use the Enclave to access the sensitive data securely
    lockedBuffer, err := enclave.Open()
    if err != nil {
        log.Fatal(err)
    }

    // Safe release of the locked buffer from memory
    defer lockedBuffer.Destroy()

    // Access the sensitive data
    //fmt.Printf("Sensitive Data: %s\n", lockedBuffer.String())

    // Securely clear the original sensitive data slice
    memguard.WipeBytes(sensitiveData)
}
