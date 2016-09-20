import itertools

def kbits(n, k):
    # memory eager approach
    result = []
    for bits in itertools.combinations(range(n), k):
        s = ['0'] * n
        for bit in bits:
            s[bit] = '1'
        result.append(''.join(s))
    return result

def kbits2(n, k):
    result = 0
    for bits in itertools.combinations(range(n), k):
        s = ['0'] * n
        for bit in bits:
            s[bit] = '1'
        result += 1
    return result

if __name__ == '__main__':
    #result = kbits(40, 20)
    #print len(result)
    #print len(list(set(result)))

    result = kbits2(40, 20)
    #result = kbits2(4, 2)
    #print result
    #result = kbits2(20, 10)
    print result
