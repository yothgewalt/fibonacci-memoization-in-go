package main

func Fibonacci(n uint) uint {
	memoized := make([]uint, n+1, n+2)
	if n < 2 {
		memoized = memoized[0:2]
	}

	memoized[0] = 0
	memoized[1] = 1

	var initial uint = 0
	for initial = 2; initial <= n; initial++ {
		memoized[initial] = memoized[initial-1] + memoized[initial-2]
	}

	return memoized[n]
}
