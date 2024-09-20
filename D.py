import math
 
def arrpp(lst):
    pp = lst[0]
    for i in range(1, len(lst)):
        pp = math.gcd(pp,lst[i])
    return pp
 
t = int(input())
for i in range(t):
    n = int(input())
    one = list(map(int,input().split()))
    two = list(map(int,input().split()))
    a1 = arrpp(one)
    b1 = arrpp(two)
    summ = a1 + b1
    count = 0
    for l in range(n):
        for r in range(l,n):
            ai = one[:l] + two[l:r+1] + one[r+1:]
            bi = two[:l] + one[l:r+1] + two[r+1:]
            A1 = arrpp(ai)
            B1 = arrpp(bi)
            sumkuy = A1 + B1
            if sumkuy > summ:
                summ = sumkuy
                count = 1
            elif sumkuy == summ:
                count += 1
    print(summ,count)
 