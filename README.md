# go-exercises

A collection of small, independently runnable Go exercises. Each exercise focuses on a core concept by solving a focused problem.

## Run an exercise

From the repository root:

```sh
go run ./concurrency/workerpool
```

## Current exercises

| Exercise | Concept |
| --- | --- |
| `concurrency/abcsequence` | Coordinating a repeating A-B-C sequence with channels |
| `concurrency/contextcancellation` | Cancelling a goroutine with `context.Context` |
| `concurrency/deadlockprevention` | Consistent mutex acquisition order |
| `concurrency/evenodd` | Alternating goroutines with channels |
| `concurrency/fanoutfanin` | Fan-out workers and fan-in result collection |
| `concurrency/mutexsharedstate` | Guarding shared state with `sync.Mutex` |
| `concurrency/producerconsumer` | Producers, consumers, channels, and `sync.WaitGroup` |
| `concurrency/rwmutexcache` | Read-heavy in-memory cache with `sync.RWMutex` |
| `concurrency/singleownergoroutine` | One goroutine owning mutable counter state |
| `concurrency/tokenbucketratelimiter` | Token-bucket rate limiting with a buffered channel, burst capacity, and `context.Context` cancellation |
| `concurrency/workerpool` | Fixed-size worker pool |

## Backend track

Build backend exercises by capability, rather than by a calendar label. The first three exercises build progressively on the same user API; the remaining directories are separate components and a final integration exercise.

```text
backend/
├── users/                    # HTTP API, layers, then safe repository
├── batchprocessor/           # Bounded concurrent job processing
├── retry/                    # Retry, backoff, timeout, and cancellation
├── cache/                    # In-memory TTL cache and cleanup
└── jobservice/               # Final integration exercise
```

Each directory is intentionally empty of implementation until you begin it. Keep the first version simple, and introduce handler/service/repository files only when the relevant exercise calls for them.

## Conventions

- Each runnable exercise is a directory containing a `package main` and `main.go`.
- Directory names use lowercase letters without separators, following Go package naming style.
- Add tests alongside reusable exercise logic as `*_test.go` files.
