first = 0
next = 1
sum = 0

while True:
    fib = first + next

    if fib % 2 == 0:
        sum += fib

    if fib > 4000000:
        print sum
        break

    first = next
    next = fib