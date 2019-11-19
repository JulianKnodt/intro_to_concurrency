# An Intro to Concurrency
### And Synchronization and all that jazz

TLDR:
When reasoning about concurrency, a programmer must often reason about the consistency
model provided by the platform being developed on, as well as have some general knowledge about
how to enforce a certain ordering of operations given a consistency model. I attempt to
formalize a few example environments which are common in industry, and also introduce some
primitives that can be used to enforce ordering of operations in these systems.

Concurrency is a pretty hot topic in the computer science world, because while it does not offer
direct algorithmic advantages, it offers many

