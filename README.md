# Find Median from Data Stream (Two Heaps)

> Striver SDE Sheet 072 · Day 12 Heaps — a max-heap guards the lower half, a min-heap the upper; the median always sits on a heap top. Go, stdlib only, tested.

## Problem

Numbers arrive one at a time; after each arrival the current median must be available. Re-sorting per insert is O(n log n) per step — the stream formulation demands better.

## Approach

Split the numbers into halves: a **max**-heap for the lower half and a **min**-heap for the upper, so the two inner boundary values — the only ones a median can involve — are both heap tops. Insert into the correct side (compare against the lower half's max), then rebalance so the sizes differ by at most one; rebalancing just moves one top across. Odd count: the median is the bigger heap's top. Even: the mean of both tops. Each operation touches O(log n) elements, and the median read is O(1). (The sheet repeats this problem on Day 22 — solved there with a different structure, a sorted insert, to contrast the trade-offs.)

**Complexity:** Add O(log n), Median O(1), O(n) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
