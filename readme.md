# An Intro to Concurrency
### And Synchronization and all that jazz

TLDR:
When reasoning about concurrency, a programmer must often reason about the consistency
model provided by the platform being developed on, as well as have some general knowledge about
how to enforce a certain ordering of operations given a consistency model. I attempt to
formalize a few example environments which are common in industry, and also introduce some
primitives that can be used to enforce ordering of operations in these systems.

Concurrency is a pretty hot topic in the computer science world, because it allows the
programmer to take advantage of the increasingly fast multi-processor architectures in
computers, and for larger organizations to handle the same operations at much higher scale.

## What is Concurrency?

Before actually discussing concurrency, it is useful to define concurrency:

Concurrency:
> Two events in a program can occur at the same logical time



