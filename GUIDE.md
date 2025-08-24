# The best way to cement all of this (structs, embedding, slices, channels, goroutines, packages) is to build something that:

- Has multiple types with methods (struct basics + receivers)
- Uses interfaces (duck typing)
- Uses channels + goroutines (concurrency)
- Has multiple packages (package/import rules)
- Uses embedding (method promotion + extra fields)

# Recommended project: Worker Pool + Task Dispatcher

## Concept:
Build a small CLI tool that downloads multiple URLs concurrently and logs their status codes, using a worker pool pattern.

## Features to include:

1. `Task` **struct** → holds URL + other metadata.
2. `Worker` **struct** → has:
    - `id` (int)
    - a `Logger` embedded **struct** with a `.Log()` method
3. `Dispatcher` struct → manages:
    - slice of workers (`[]*Worker`)
    - channels for tasks and results
4. Concurrency:
    - Workers run as goroutines
    - Communication via channels (`chan Task`, `chan Result`)
5. Interfaces:
    - Define a `Processor` interface with a `Process(Task) Result` method
    - Make `Worker` implement it
6. Multiple packages:
    - `models` → define `Task`, `Result`
    - `worker` → define `Worker`, `Dispatcher`
    - `main` → glue it all together
7. Struct embedding:
    - Embed `Logger` into `Worker` so it inherits `.Log()`
8. Slices:
    - Slice of `Processor` to store workers
9. Buffered vs unbuffered channels:
    - Use unbuffered for signaling
    - Try buffered for task queue

# Example flow
- Create Dispatcher with N workers
- Pass a slice of URLs
- Dispatcher sends tasks to workers via channel
- Workers fetch URL, send result via channel
- Main receives results, prints summary