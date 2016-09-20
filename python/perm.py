# Python program to print all permutations with
# duplicates allowed
 
def toString(List):
    return ''.join(List)
 
# Function to print permutations of string
# This function takes three parameters:
# 1. String
# 2. Starting index of the string
# 3. Ending index of the string.
perm_list = []
def permute(a, l, r):
    global perm_list
    if l==r:
        #print toString(a)
        perm_list.append(toString(a))
    else:
        for i in xrange(l,r+1):
            a[l], a[i] = a[i], a[l]
            permute(a, l+1, r)
            a[l], a[i] = a[i], a[l] # backtrack
  
if __name__ == '__main__':
    # Driver program to test the above function
    string = "RRRRDDDD"
    n = len(string)
    a = list(string)
    permute(a, 0, n-1)
    #print perm_list
    #print "-"*3
    #print list(set(perm_list))
    print(len(list(set(perm_list))))
 
    # This code is contributed by Bhavya Jain
    # This algo is not memory efficient.
    # R=10 and D=10 takes too much memory (it's not even 20,20)
