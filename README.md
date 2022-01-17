# Merkle tree basics

This GO project captures how a Merkle tree basically works. It contains:

* a package `merkle` with corresponding tests to deal with tree functionalities
* a `main` to be able to run and play around with the package `merkle`

### Install

You need GO (~1.17) to be [installed locally](https://go.dev/doc/install)

### Test

Ensure the integrity of the project with:

```console
go test ./... -v
```

### Run

You can run the main entry point with:

```console
go run main.go
```

## Additional questions

1. Using the illustration above, let’s assume I know the whole Merkle tree. Someone gives me the L2 
data block, but I don’t trust them. How can I check if L2 data is valid?

-> In that case, we know the whole tree, so simply we can return (using the `level` method) the 
array of hashes corresponding to the last level (i.e. level = height of tree). Then the given hash 
L2 is valid if contained in that array.

2. I know only the L3 data block and the Merkle root. What is the minimum information needed to 
check that the L3 data block and the Merkle root belong to the same Merkle tree?

-> This is what is basically used in Bitcoin. The minimum information needed to check that L3 is 
valid and belong to that tree would be called a `merkle path`. Here for L3, it would consist of
knowing: Hash 1-1, Hash 0 (and of course we have already the root and L3)

3. What are some Merkle tree use cases? 

-> On top of my head, I only remember one use case now. `Light clients` from the Bitcoin network 
being able to verify selected incoming transactions belonging to a block in the blockchain by 
only having block headers locally and asking for a merkle path other services (i.e. full nodes) 
to perform the verification.

Then I will have a look at Wikipedia after that for other real life use cases, which I am sure are 
plenty!
