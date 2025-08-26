# Experiments

_A sandbox of learning and experimentation — applying what I learn, one project at a time. When in doubt code it out_

---

## Projects

### Bloom Filter Visualization (React + Tailwind)

- **Implemented a Bloom filter**, a probabilistic data structure used for element membership testing.
- Built a **visual demo UI** using **React** and **Tailwind CSS** to showcase how the filter works in real time.
- This helped deepen my understanding of **probabilistic algorithms** and frontend integration.

### In-Memory Cache (Go)

- Developed a **basic in-memory cache** in **Go** with essential features:
  - **TTL-based key expiration** leveraging **goroutines** for periodic cleanup.
  - **Configurable loader function pattern** to dynamically handle data loading.
- Started exploring how databases communicate using the **RESP protocol** and how snapshots are taken—**implementation planned in next iterations**.

### Go-Based Web Server (Work in Progress)

- Currently building a **web server from scratch** in **Go**.
- Learning core concepts like **networking fundamentals**, **concurrency**, and **HTTP request handling**, paving the way for production-ready backend services.

---

## Tech Stack & Concepts

| Project           | Technologies / Concepts                                                 |
| ----------------- | ----------------------------------------------------------------------- |
| Bloom Filter Demo | React, Tailwind CSS, Probabilistic Data Structures                      |
| In-Memory Cache   | Go, Goroutines, TTL Cache, Config Patterns, RESP Protocol, DB Snapshots |
| Web Server        | Go, Networking, Concurrency, HTTP Fundamentals                          |

---

## Roadmap

- Add **snapshot and persistence** capabilities to the in-memory cache leveraging insights from RESP and DB mechanisms.
- Expand the Go web server to support routing, middleware, and RESTful endpoints.
- Integrate unit tests and benchmarks to measure performance and reliability.

---

## Getting Started

1. **Clone the repository:**
   ```bash
   git clone https://github.com/Soham2402/Experiments.git
   ```
