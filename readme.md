# An Intro to Concurrency
### And Synchronization and all that jazz

TLDR:
When reasoning about concurrency, a programmer must often reason about the consistency
model provided by the platform being developed on, as well as have some general knowledge about
how to enforce a certain ordering of operations with that model. Here, I attempt to
formalize a few example environments which are common in industry, and also introduce some
primitives that can be used to enforce ordering of operations in these systems.

Concurrency is a hot topic in CS, because it not only allows the
programmer to take advantage of the increasingly fast multi-processor architectures in
computers, but also allows for larger organizations to handle the same operations
at much higher scale.

## What is Concurrency?

Before actually discussing concurrency, it is useful to define concurrency:

> Concurrency:
> Two "parts of the program" can occur at the same logical time.

What is "logical time"? Logical time as I define it here is refers to the ordering of events in
in the program. In this case, what this means is that if two parts of a program A and B are
concurrent, then either A could happen first or B could happen.

In this case, a part of program could be anything, some series of lines of codes, a function, or
more.

Concurrency can be seen just by observing that our computers can run multiple applications at
the same time. There likely are not enough cores to run each applicaton at the same time(there
are probably only 1-4 cores on your computer). Thus, in order to allow a user to run more
programs at the same time, the operating system provides concurrency by letting applications
time-share the CPU.

---

Let's look at an example of concurrency in programming, JavaScript,
and more specifically [NodeJS](https://nodejs.org/en/).

JavaScript is a single threaded language, whether it's running in the browser or as a server.
NodeJS self-describes itself as "an asynchronous event-driven JavaScript runtime". What does
that actually mean?

- Events in JavaScript refer to the various hooks, that can be waited on. It's best illustrated
  through a couple examples:
```javascript
// waiting for timer
setTimeout(() => {
  console.log("hello 1 second later")
}, 1000)

// waiting for file events
fs.open("my_file.txt", "r", (err, file) => {
  console.log("My file is open!")
  fs.close(file, (err) => {
    console.log("My file is now closed.")
  })
})
```
- Asynchronous in this case simply refers to the fact that the runtime allows handling of
  multiple events concurrently. Control flow of a program is done by waiting for various events,
  and this allows for some potential benefits, because network calls can now be handled in the
  background, without blocking.

Node is primarily used as a server side language, and the main reason for that is it's
concurrency model, as it allows many requests to be processed, while not waiting for network
round trips.

Let's look at a basic examples of concurrency in JS:

[Byte Count](wc.js)

How do concurrent functions in JavaScript communicate? Well, since there is only one thread to
run on, only one thing is ever running at once. Because only one thing is running at once, we
know that any object changed must've been the result Thus, we can communicate through shared data
structures. This is because any data structure that is shared between different functions can
know that no one else will modify it while it's running, and thus it can be sure that what it's
observed must've been the result of previous code.

Why do people like Node? JavaScript is a [wat](https://www.destroyallsoftware.com/talks/wat)
language, and inherently is slow due to dynamic typing. Yet, the concurrency offered makes it
easy to build efficient sequentially consistent programs. Thus, JavaScript sacrificies
efficiency for understandibility, and that's what it makes attractive as a language.

## What about Parallelism?

Parallelism and concurrency are closely related, and I define parallelism as nearly identical to
concurrency, with one modification:

> Parallelism:
> Two "parts of the program" can occur at the same real time.

For a more in-depth discussion of the differences between parallelism and concurrency, check out this
[video](https://www.youtube.com/watch?v=cN_DpYBzKso) from the creator of Golang.

Real time in this case refers to the wall clock time(i.e. 10:30 A.M EST). Thus, what this means
is that multiple parts of the programming can actually be running while another is.

This is in stark contrast to the model JavaScript provides, where it is only possible for one
line of code to be running at any given time(ignoring things happening in the background).

Let's instead switched to a runtime which allows for parallelism(but doesn't require it).

Golang provides a nice abstraction around the messiness that are threads, allowing the user
instead to focus on the concurrent aspects of the code, and letting the runtime handle execution
of concurrency under the hood.

For starter's, let's look at a simple multiple producers, single consumer example.

A producer, in this case, is just something that outputs a value, and a consumer is something
that uses that value.

[Producer/Consumer](prod_con.go)

Of course, this simple example doesn't capture all the complexity of reasoning about
concurrency.

Consider a bit more involved of an example:
Imagine you have a link of files you need to fetch, and they need to be stored locally.
This is the same problem a client for a server might have, where it needs to send multiple
requests to different servers, and store all of the responses before forwarding them somewhere
else, so this example is a little more practical at the one above.

Let's look at two examples of how to implement that.

[Serially Fetching](naive_fetch.go)

[Fetching links](fetch_stuff.go)

[Fetching links alternative](fetch_stuff2.go)

## Going even further

<!---
TODO add bit about compilers, reorderings, atomics and fences in golang
-->


